// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules // import "src.tricot.io/public/bazel2x/bazel/builtins/rules"

import (
	"go.starlark.net/starlark"

	builtins_args "src.tricot.io/public/bazel2x/bazel/builtins/args"
	"src.tricot.io/public/bazel2x/bazel/core"
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

var _ builtins_args.ProcessArgsTarget = (*CcLibraryTarget)(nil)
var _ core.Target = (*CcLibraryTarget)(nil)

func (self *CcLibraryTarget) DidProcessArgs(ctx core.Context) error {
	// TODO(vtl): Check fields.
	return nil
}

func (self *CcLibraryTarget) String() string {
	return targetToString("cc_library", self)
}

// CcLibrary implements the Bazel cc_library rule.
var CcLibrary = newRule("cc_library",
	func(ctx core.Context, args starlark.Tuple, kwargs []starlark.Tuple) error {

		target := &CcLibraryTarget{}
		if err := builtins_args.ProcessArgs(args, kwargs, ctx, target); err != nil {
			return err
		}
		ctx.BuildTargets().Add(target)
		return nil
	})
