// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"fmt"

	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

// BuiltinsImpl implements Builtins.
type BuiltinsImpl struct {
	NoOpBuiltinsGlobalsImpl
	NoOpBuiltinsFunctionsImpl
	NoOpBuiltinsAndroidRulesImpl
	//NoOpBuiltinsCcRulesImpl
	NoOpBuiltinsJavaRulesImpl
	NoOpBuiltinsObjCRulesImpl
	NoOpBuiltinsProtoBufRulesImpl
	NoOpBuiltinsPythonRulesImpl
	NoOpBuiltinsShellRulesImpl
	NoOpBuiltinsExtraActionsRulesImpl
	NoOpBuiltinsGeneralRulesImpl
	NoOpBuiltinsPlatformRulesImpl
	NoOpBuiltinsWorkspaceRulesImpl

	ctx Context
}

func (self *BuiltinsImpl) CcBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) CcImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) CcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
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
		"srcs", &srcs,
		"data", &data,
		"hdrs", &hdrs,
		"alwayslink", &alwayslink,
		"compatible_with", &compatible_with,
		"copts", &copts,
		"defines", &defines,
		"deprecation", &deprecation,
		"distribs", &distribs,
		"exec_compatible_with", &exec_compatible_with,
		"features", &features,
		"include_prefix", &include_prefix,
		"includes", &includes,
		"licenses", &licenses,
		"linkopts", &linkopts,
		"linkstatic", &linkstatic,
		"nocopts", &nocopts,
		"restricted_to", &restricted_to,
		"strip_include_prefix", &strip_include_prefix,
		"tags", &tags,
		"testonly", &testonly,
		"textual_hdrs", &textual_hdrs,
		"toolchains", &toolchains,
		"visibility", &visibility,
		"win_def_file", &win_def_file)
	if err != nil {
		return starlark.None, err
	}

	nameLabel := core.Label{
		Workspace: self.ctx.Label().Workspace,
		Package:   self.ctx.Label().Package,
		Target:    core.TargetName(name),
	}
	if !nameLabel.IsValid() {
		return starlark.None, fmt.Errorf("invalid target name %v", name)
	}

	// TODO(vtl)
	// fmt.Printf("--> %v\n", nameLabel)

	return starlark.None, nil
}

func (self *BuiltinsImpl) CcProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) FdoPrefetchHints(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) FdoProfile(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) CcTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) CcToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *BuiltinsImpl) CcToolchainSuite(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ Builtins = (*BuiltinsImpl)(nil)

func NewBuiltinsImpl(ctx Context) *BuiltinsImpl {
	return &BuiltinsImpl{ctx: ctx}
}
