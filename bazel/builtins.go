// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

type BuiltinsIface interface {
	// https://docs.bazel.build/versions/master/be/functions.html#package
	//
	// BUILD file only.
	//
	// package(default_deprecation, default_testonly, default_visibility, features)
	Package(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#package_group
	//
	// BUILD file only.
	//
	// package_group(name, packages, includes)
	PackageGroup(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#exports_files
	//
	// BUILD file only.
	//
	// exports_files([label, ...], visibility, licenses)
	ExportsFiles(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#glob
	//
	// glob(include, exclude=[], exclude_directories=1)
	Glob(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value,
		error)

	// https://docs.bazel.build/versions/master/be/functions.html#select
	//
	// select(
	//     {conditionA: valuesA, conditionB: valuesB, ...},
	//     no_match_error = "custom message"
	// )
	Select(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#workspace
	//
	// WORKSPACE file only.
	//
	// workspace(name = "com_example_project")
	Workspace(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)
}

const builtinsImplKey = "bazel2make-bazel-builtins-impl"

func SetBuiltinsImpl(thread *starlark.Thread, builtinsImpl BuiltinsIface) {
	thread.SetLocal(builtinsImplKey, builtinsImpl)
}

func GetBuiltinsImpl(thread *starlark.Thread) BuiltinsIface {
	return thread.Local(builtinsImplKey).(BuiltinsIface)
}

var Builtins = starlark.StringDict{
	"package":       starlark.NewBuiltin("package", packageThunk),
	"package_group": starlark.NewBuiltin("package_group", packageGroupThunk),
	"exports_files": starlark.NewBuiltin("exports_files", exportsFilesThunk),
	"glob":          starlark.NewBuiltin("glob", globThunk),
	"select":        starlark.NewBuiltin("select", selectThunk),
	"workspace":     starlark.NewBuiltin("workspace", workspaceThunk),
}

func packageThunk(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return GetBuiltinsImpl(thread).Package(thread, args, kwargs)
}

func packageGroupThunk(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return GetBuiltinsImpl(thread).PackageGroup(thread, args, kwargs)
}

func exportsFilesThunk(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return GetBuiltinsImpl(thread).ExportsFiles(thread, args, kwargs)
}

func globThunk(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return GetBuiltinsImpl(thread).Glob(thread, args, kwargs)
}

func selectThunk(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return GetBuiltinsImpl(thread).Select(thread, args, kwargs)
}

func workspaceThunk(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return GetBuiltinsImpl(thread).Workspace(thread, args, kwargs)
}
