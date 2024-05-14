// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package genwinsyscallproto

import (
	"strings"
	"testing"

	"github.com/microsoft/go-winmd"
	"github.com/microsoft/go-winmd/flags"
)

func TestContext_writeType_cycle(t *testing.T) {
	t.Skip("cycles can't be built with SigType rather than *SigType, and this code only supports SigType")

	p1 := winmd.SigType{Kind: flags.ElementType_PTR}
	p2 := winmd.SigType{Kind: flags.ElementType_PTR}
	// These are copies, not pointers...
	var p1a any = p1
	var p2a any = p2
	// ...So this doesn't cause a cycle.
	p2.Value = p1a
	p1.Value = p2a

	// If we can create a cycle (e.g. if we change to *SigType e.g. for performance reasons) then
	// this code would test it. The strings.Contains check should likely be changed to errors.Is if
	// we do that.
	var b strings.Builder
	var c Context
	if err := c.writeType(&b, &p1, ArchAll); err == nil {
		t.Fatalf("expected error due to detected cycle, but no error was returned")
	} else if !strings.Contains(err.Error(), "cycle") {
		t.Fatalf("got an error, but not a cycle detection error: %v", err)
	}
}
