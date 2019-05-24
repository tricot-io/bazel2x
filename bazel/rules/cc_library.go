// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type CcLibraryTarget struct {
	TargetCommon
	Srcs               *[]core.Label `bazel:"srcs"`
	Hdrs               *[]core.Label `bazel:"hdrs"`
	Alwayslink         *bool         `bazel:"alwayslink"`
	Copts              *[]string     `bazel:"copts"`
	Defines            *[]string     `bazel:"defines"`
	IncludePrefix      *string       `bazel:"include_prefix"`
	Includes           *[]string     `bazel:"includes"`
	Linkopts           *[]string     `bazel:"linkopts"`
	Linkstatic         *bool         `bazel:"linkstatic"`
	Nocopts            *[]string     `bazel:"nocopts"`
	StripIncludePrefix *string       `bazel:"strip_include_prefix"`
	TextualHdrs        *[]core.Label `bazel:"textual_hdrs"`
	WinDefFile         *core.Label   `bazel:"win_def_file"`
}

var _ ProcessRuleArgsTargetStruct = (*CcLibraryTarget)(nil)
var _ core.Target = (*CcLibraryTarget)(nil)

func (self *CcLibraryTarget) Process(ctx core.Context) error {
	if err := self.TargetCommon.Process(ctx); err != nil {
		return nil
	}
	// TODO(vtl): Check other fields.
	return nil
}

func (self *CcLibraryTarget) String() string {
	return targetToString("cc_library", self)
}

// CcLibrary implements the Bazel cc_library rule.
var CcLibrary = newRule("cc_library", func(ctx core.Context, kwargs []starlark.Tuple) error {
	target := &CcLibraryTarget{}
	if err := processRuleArgs(kwargs, ctx, target); err != nil {
		return err
	}
	ctx.BuildTargets().Add(target)
	return nil
})
