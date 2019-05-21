// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

type BuiltinsIface interface {
	// Globals
	// https://docs.bazel.build/versions/master/skylark/lib/globals.html

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#select
	// unknown select(x, no_match_error='')
	//
	// https://docs.bazel.build/versions/master/be/functions.html#select
	// select(
	//     {conditionA: valuesA, conditionB: valuesB, ...},
	//     no_match_error = "custom message"
	// )
	Select(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#workspace
	// None workspace(name, managed_directories={})
	//
	// https://docs.bazel.build/versions/master/be/functions.html#workspace
	// workspace(name = "com_example_project")
	Workspace(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// TODO(vtl): The rest of them.

	// Build functions
	// https://docs.bazel.build/versions/master/be/functions.html

	// https://docs.bazel.build/versions/master/be/functions.html#package
	// package(default_deprecation, default_testonly, default_visibility, features)
	Package(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#package_group
	// package_group(name, packages, includes)
	PackageGroup(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#exports_files
	// exports_files([label, ...], visibility, licenses)
	ExportsFiles(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#glob
	// glob(include, exclude=[], exclude_directories=1)
	Glob(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value,
		error)

	// TODO(vtl): More (e.g., rules).
}

const builtinsImplKey = "bazel2make-bazel-builtins-impl"

func SetBuiltinsImpl(thread *starlark.Thread, builtinsImpl BuiltinsIface) {
	thread.SetLocal(builtinsImplKey, builtinsImpl)
}

func GetBuiltinsImpl(thread *starlark.Thread) BuiltinsIface {
	return thread.Local(builtinsImplKey).(BuiltinsIface)
}

func MakeBuiltins(ctx *Context) starlark.StringDict {
	// TODO(vtl): Customize this.
	return starlark.StringDict{
		// Globals
		// https://docs.bazel.build/versions/master/skylark/lib/globals.html
		"select": starlark.NewBuiltin("select",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return GetBuiltinsImpl(thread).Select(thread, args, kwargs)
			}),
		"workspace": starlark.NewBuiltin("workspace",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return GetBuiltinsImpl(thread).Workspace(thread, args, kwargs)
			}),
		// TODO(vtl): The rest of them.
		// Build functions
		// https://docs.bazel.build/versions/master/be/functions.html
		"package": starlark.NewBuiltin("package",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return GetBuiltinsImpl(thread).Package(thread, args, kwargs)
			}),
		"package_group": starlark.NewBuiltin("package_group",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return GetBuiltinsImpl(thread).PackageGroup(thread, args, kwargs)
			}),
		"exports_files": starlark.NewBuiltin("exports_files",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return GetBuiltinsImpl(thread).ExportsFiles(thread, args, kwargs)
			}),
		"glob": starlark.NewBuiltin("glob",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return GetBuiltinsImpl(thread).Glob(thread, args, kwargs)
			}),
		// TODO(vtl): More (e.g., globals, rules).
	}
}

/*
# Globals

https://docs.bazel.build/versions/master/skylark/lib/globals.html#analysis_test_transition
analysis_test_transition [function; experimental]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#aspect
aspect [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#bind
bind [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#configuration_field
configuration_field [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#depset
depset [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#existing_rules
existing_rules [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#existing_rules
fail [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#PACKAGE_NAME
PACKAGE_NAME [string; deprecated]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#provider
provider [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#register_execution_platforms
register_execution_platforms [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#register_toolchains
register_toolchains [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#REPOSITORY_NAME
REPOSITORY_NAME [string; deprecated]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#repository_rule
repository_rule [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#rule
rule [function]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#select
select [function; duplicated above]

https://docs.bazel.build/versions/master/skylark/lib/globals.html#workspace
workspace [function; duplicated above]

# Rules

# Android
android_binary
aar_import
android_library
android_instrumentation_test
android_local_test
android_device
android_ndk_repository
android_sdk_repository

# C / C++
cc_binary
cc_import
cc_library
cc_proto_library
fdo_prefetch_hints
fdo_profile
cc_test
cc_toolchain
cc_toolchain_suite

# Java
java_binary
java_import
java_library
java_lite_proto_library
java_proto_library
java_test
java_package_configuration
java_plugin
java_runtime
java_toolchain

# Objective-C
apple_binary
apple_static_library
j2objc_library
objc_import
objc_library
objc_proto_library

# Protocol Buffer
proto_lang_toolchain
proto_library

# Python
py_binary
py_library
py_test
py_runtime

# Shell
sh_binary
sh_library
sh_test

# Extra Actions

action_listener
extra_action

# General
filegroup
genquery
test_suite
alias
config_setting
genrule

# Platform
constraint_setting
constraint_value
platform
toolchain

# Workspace
bind
local_repository
maven_jar
maven_server
new_local_repository
xcode_config
xcode_version
*/
