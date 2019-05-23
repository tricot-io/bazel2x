// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"go.starlark.net/starlark"

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

var _ ProcessRuleArgsTargetStruct = (*CcBinaryTarget)(nil)
var _ core.Target = (*CcBinaryTarget)(nil)

func (self *CcBinaryTarget) Process(ctx core.Context) error {
	if err := self.TargetCommon.Process(ctx); err != nil {
		return nil
	}
	if err := self.TargetCommonBinary.Process(ctx); err != nil {
		return nil
	}
	// TODO(vtl): Check other fields.
	return nil
}

func (self *CcBinaryTarget) String() string {
	return targetToString("cc_binary", self)
}

// CcBinary implements the Bazel cc_binary rule.
func CcBinary(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	ctx := core.GetContext(thread)
	target := &CcBinaryTarget{}
	err := ProcessRuleArgs(args, kwargs, ctx, target)
	if err != nil {
		return starlark.None, err
	}
	ctx.BuildTargets().Add(target)
	return starlark.None, nil
}
