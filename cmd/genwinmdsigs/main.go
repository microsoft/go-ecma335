// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"debug/pe"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/microsoft/go-winmd"
	"github.com/microsoft/go-winmd/genwinsyscallproto"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	source := flag.String("source", "", "The win32metadata file to parse and generate signatures for.")
	output := flag.String("o", "", "Output file name (prints to stdout if omitted).")
	filter := flag.String("filter", "", "If specified, only generate signatures for methods 'namespace::method' that match this regex.")
	templatePath := flag.String("template", "", "If specified, a text/template style template file to use. Use '{{.SysCalls}}' for the '//sys' content.")

	flag.Parse()

	if *source == "" {
		return errors.New("no source passed: nothing to do")
	}
	var filterRegexp *regexp.Regexp
	if *filter != "" {
		var err error
		filterRegexp, err = regexp.Compile(*filter)
		if err != nil {
			return err
		}
	}

	var contentTemplate *template.Template
	if *templatePath != "" {
		var err error
		contentTemplate, err = template.ParseFiles(*templatePath)
		if err != nil {
			return err
		}
	}

	start := time.Now()

	pefile, err := pe.Open(*source)
	if err != nil {
		return err
	}
	defer pefile.Close()
	f, err := winmd.New(pefile)
	if err != nil {
		return err
	}

	b := map[genwinsyscallproto.Arch]*strings.Builder{
		genwinsyscallproto.Arch386:   {},
		genwinsyscallproto.ArchAMD64: {},
		genwinsyscallproto.ArchARM64: {},
		genwinsyscallproto.ArchAll:   {},
		genwinsyscallproto.ArchNone:  {},
	}

	if err := writePrototypes(b, f, filterRegexp); err != nil {
		return err
	}

	for arch, w := range b {
		if w.Len() == 0 {
			continue
		}
		content := w.String()
		var finalContent string

		if contentTemplate != nil {
			var templateBuilder strings.Builder
			if err := contentTemplate.Execute(&templateBuilder, struct {
				SysCalls string
				Arch     string
			}{content, arch.String()}); err != nil {
				return err
			}
			finalContent = templateBuilder.String()
		} else {
			finalContent = content
		}

		formattedContent, err := format.Source([]byte(finalContent))
		if err != nil {
			log.Printf("Unable to format generated code, writing unformatted code instead. Error: %v", err)
			formattedContent = []byte(finalContent)
		}

		end := time.Now()
		log.Printf("Time elapsed to produce sys signatures: %v\n", end.Sub(start))

		if *output != "" {
			target := *output
			if arch != genwinsyscallproto.ArchAll {
				target = strings.TrimSuffix(target, ".go") + "_" + arch.String() + ".go"
			}
			os.WriteFile(target, formattedContent, 0666)
		} else {
			log.Printf("Printing signature results for %s because no output path was specified:\n", arch)
			log.Println("---")
			log.Println(finalContent)
		}
	}
	return nil
}

func writePrototypes(b map[genwinsyscallproto.Arch]*strings.Builder, f *winmd.Metadata, filterRegexp *regexp.Regexp) error {
	context, err := genwinsyscallproto.NewContext(f)
	if err != nil {
		return err
	}

	for i := uint32(0); i < f.Tables.TypeDef.Len; i++ {
		r, err := f.Tables.TypeDef.Record(winmd.Index(i))
		if err != nil {
			return err
		}

		if !strings.Contains(r.Name.String(), "Apis") {
			continue
		}

		archSeen := make(map[genwinsyscallproto.Arch]bool)
		for j := r.MethodList.Start; j < r.MethodList.End; j++ {
			md, err := f.Tables.MethodDef.Record(j)
			if err != nil {
				return err
			}
			spec := fmt.Sprintf("%v::%v.%v", r.Namespace, r.Name, md.Name)
			if filterRegexp != nil && !filterRegexp.MatchString(spec) {
				continue
			}

			supportedArches := context.MethodDefSupportedArch(j)
			for _, arch := range supportedArches.Unique() {
				w := b[arch]

				// Write a comment describing this chunk of methods.
				if !archSeen[arch] {
					archSeen[arch] = true
					w.WriteString("\n\n// APIs for ")
					w.WriteString(r.Namespace.String())
				}
				w.WriteString("\n")

				if err := context.WriteMethod(w, j, md, arch); err != nil {
					// Include context in the error for diag purposes.
					// writeSys may have partially written into b. This is actually convenient for diag.
					lines := strings.Split(w.String(), "\n")
					if len(lines) > 5 {
						lines = lines[len(lines)-5:]
					}

					return fmt.Errorf(
						"error context: \n---\n%v\n---\nfailed to write sys line for %v.Apis method %v: %v",
						strings.Join(lines, "\n"), r.Namespace, md.Name, err)
				}
			}
		}
	}
	return context.WriteUsedTypeDefs(b)
}
