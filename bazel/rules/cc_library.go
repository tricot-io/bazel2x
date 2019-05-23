// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"fmt"

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

func (self *CcLibraryTarget) Process(ctx core.Context) error {
	if err := self.TargetCommon.Process(ctx); err != nil {
		return nil
	}
	// TODO(vtl)
	return nil
}

// CcLibrary implements the Bazel cc_library rule.
func CcLibrary(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	// TODO(vtl): Factor out common attributes.
	// cc_library
	// Required:
	var name string
	// Optional:
	var deps *starlark.List
	var srcs *starlark.List
	var data *starlark.List
	var hdrs *starlark.List
	var alwayslink bool
	var compatible_with *starlark.List
	var copts *starlark.List
	var defines *starlark.List
	var deprecation string
	var distribs *starlark.List
	var exec_compatible_with *starlark.List
	var features *starlark.List
	var include_prefix string
	var includes *starlark.List
	var licenses *starlark.List
	var linkopts *starlark.List
	var linkstatic bool
	var nocopts string
	var restricted_to *starlark.List
	var strip_include_prefix string
	var tags *starlark.List
	var testonly bool
	var textual_hdrs *starlark.List
	var toolchains *starlark.List
	var visibility *starlark.List
	var win_def_file string

	err := starlark.UnpackArgs(
		"cc_library", args, kwargs,
		"name", &name,
		"deps?", &deps,
		"srcs?", &srcs,
		"data?", &data,
		"hdrs?", &hdrs,
		"alwayslink?", &alwayslink,
		"compatible_with?", &compatible_with,
		"copts?", &copts,
		"defines?", &defines,
		"deprecation?", &deprecation,
		"distribs?", &distribs,
		"exec_compatible_with?", &exec_compatible_with,
		"features?", &features,
		"include_prefix?", &include_prefix,
		"includes?", &includes,
		"licenses?", &licenses,
		"linkopts?", &linkopts,
		"linkstatic?", &linkstatic,
		"nocopts?", &nocopts,
		"restricted_to?", &restricted_to,
		"strip_include_prefix?", &strip_include_prefix,
		"tags?", &tags,
		"testonly?", &testonly,
		"textual_hdrs?", &textual_hdrs,
		"toolchains?", &toolchains,
		"visibility?", &visibility,
		"win_def_file?", &win_def_file)
	if err != nil {
		return starlark.None, err
	}

	ctx := core.GetContext(thread)

//FIXME
	target := CcLibraryTarget{}
	err = ProcessRuleArgs(args, kwargs, ctx, &target)
	if err != nil {
		return starlark.None, err
	}
fmt.Printf("--> %v\n", TargetToString("cc_library", &target))

	nameLabel := core.Label{ctx.Label().Workspace, ctx.Label().Package, core.TargetName(name)}
	if !nameLabel.IsValid() {
		return starlark.None, fmt.Errorf("invalid target name %v", name)
	}

	// TODO(vtl)
	// fmt.Printf("--> %v\n", nameLabel)

	return starlark.None, nil
}
