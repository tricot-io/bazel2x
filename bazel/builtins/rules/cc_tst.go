// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// NOTE: This file is named "cc_tst.go", since those ending in "_test.go" are test files.

package rules

import (
	"go.starlark.net/starlark"

	builtins_args "bazel2cmake/bazel/builtins/args"
	"bazel2cmake/bazel/core"
)

// TODO(vtl): The attributes are identical to CcBinaryTarget, except the latter also has linkshared.
type CcTestTarget struct {
	TargetCommon
	TargetCommonTest
	Srcs                   *[]core.Label `bazel:"srcs"`
	AdditionalLinkerInputs *[]core.Label `bazel:"additional_linker_inputs"`
	Copts                  *[]string     `bazel:"copts"`
	Defines                *[]string     `bazel:"defines"`
	Includes               *[]string     `bazel:"includes"`
	Linkopts               *[]string     `bazel:"linkopts"`
	Linkstatic             *bool         `bazel:"linkstatic"`
	Malloc                 *core.Label   `bazel:"malloc"`
	Nocopts                *[]string     `bazel:"nocopts"`
	Stamp                  *int64        `bazel:"stamp"`
	WinDefFile             *core.Label   `bazel:"win_def_file"`
}

var _ builtins_args.ProcessArgsTarget = (*CcTestTarget)(nil)
var _ core.Target = (*CcTestTarget)(nil)

func (self *CcTestTarget) DidProcessArgs(ctx core.Context) error {
	// TODO(vtl): Check fields.
	return nil
}

func (self *CcTestTarget) String() string {
	return targetToString("cc_test", self)
}

// CcTest implements the Bazel cc_test rule.
var CcTest = newRule("cc_test",
	func(ctx core.Context, args starlark.Tuple, kwargs []starlark.Tuple) error {

		target := &CcTestTarget{}
		if err := builtins_args.ProcessArgs(args, kwargs, ctx, target); err != nil {
			return err
		}
		ctx.BuildTargets().Add(target)
		return nil
	})
