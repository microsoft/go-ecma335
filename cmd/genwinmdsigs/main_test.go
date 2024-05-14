// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"debug/pe"
	"go/format"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/microsoft/go-winmd"
	"github.com/microsoft/go-winmd/genwinsyscallproto"
)

func TestWriteMethod(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := map[genwinsyscallproto.Arch]*strings.Builder{
		genwinsyscallproto.Arch386:   {},
		genwinsyscallproto.ArchAMD64: {},
		genwinsyscallproto.ArchARM64: {},
		genwinsyscallproto.ArchAll:   {},
		genwinsyscallproto.ArchNone:  {},
	}
	r := regexp.MustCompile(`^(Windows\.Win32\.Storage\.FileSystem|Windows\.Win32\.Security\.Cryptography|Windows\.Win32\.System\.Diagnostics\.Debug)::`)
	if err := writePrototypes(b, f, r); err != nil {
		t.Fatal(err)
	}
	for arch, w := range b {
		formattedContent, err := format.Source([]byte(w.String()))
		if err != nil {
			t.Fatal(err)
		}
		target := "prototypes.golden"
		if arch != genwinsyscallproto.ArchAll {
			target += "_" + arch.String()
		}
		target += ".go"
		Check(t, "go test ./cmd/genwinmdsigs", filepath.Join("testdata", target), string(formattedContent))
	}
}

func TestFullFile(t *testing.T) {
	f, err := openTestWinmd()
	if err != nil {
		t.Fatal(err)
	}
	b := map[genwinsyscallproto.Arch]*strings.Builder{
		genwinsyscallproto.Arch386:   {},
		genwinsyscallproto.ArchAMD64: {},
		genwinsyscallproto.ArchARM64: {},
		genwinsyscallproto.ArchAll:   {},
		genwinsyscallproto.ArchNone:  {},
	}
	if err := writePrototypes(b, f, nil); err != nil {
		t.Fatal(err)
	}
	for _, w := range b {
		_, err = format.Source([]byte(w.String()))
		if err != nil {
			t.Fatal(err)
		}
	}

	// The generated source code is ~4 MB, so don't write it to source control as a golden file.
	// This test only checks that the generation process doesn't fail and doesn't take an
	// exceptionally long time.

	// To see the output, use:
	//   go run ./cmd/genwinmdsigs -o all.go.temp -source .\testdata\Windows.Win32.winmd
}

func openTestWinmd() (*winmd.Metadata, error) {
	pefile, err := pe.Open("../../testdata/Windows.Win32.winmd")
	if err != nil {
		return nil, err
	}
	defer pefile.Close()
	return winmd.New(pefile)
}
