// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"go.starlark.net/starlark"

	builtins_args "bazel2cmake/bazel/builtins/args"
	"bazel2cmake/bazel/core"
)

type CcBinaryTarget struct {
	TargetCommon
	TargetCommonBinary
	Srcs                   *[]core.Label `bazel:"srcs"`
	AdditionalLinkerInputs *[]core.Label `bazel:"additional_linker_inputs"`
	Copts                  *[]string     `bazel:"copts"`
	Defines                *[]string     `bazel:"defines"`
	Includes               *[]string     `bazel:"includes"`
	Linkopts               *[]string     `bazel:"linkopts"`
	Linkshared             *bool         `bazel:"linkshared"`
	Linkstatic             *bool         `bazel:"linkstatic"`
	Malloc                 *core.Label   `bazel:"malloc"`
	Nocopts                *[]string     `bazel:"nocopts"`
	Stamp                  *int64        `bazel:"stamp"`
	WinDefFile             *core.Label   `bazel:"win_def_file"`
}

var _ builtins_args.ProcessArgsTarget = (*CcBinaryTarget)(nil)
var _ core.Target = (*CcBinaryTarget)(nil)

func (self *CcBinaryTarget) DidProcessArgs(ctx core.Context) error {
	// TODO(vtl): Check fields.
	return nil
}

func (self *CcBinaryTarget) String() string {
	return targetToString("cc_binary", self)
}

// CcBinary implements the Bazel cc_binary rule.
var CcBinary = newRule("cc_binary",
	func(ctx core.Context, args starlark.Tuple, kwargs []starlark.Tuple) error {

		target := &CcBinaryTarget{}
		if err := builtins_args.ProcessArgs(args, kwargs, ctx, target); err != nil {
			return err
		}
		ctx.BuildTargets().Add(target)
		return nil
	})
