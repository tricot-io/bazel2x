// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

// Globals
// https://docs.bazel.build/versions/master/skylark/lib/globals.html
type BuiltinsGlobalsIface interface {
	// TODO(vtl)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#analysis_test_transition
	// transition analysis_test_transition(settings)
	AnalysisTestTransition(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#aspect
	// Aspect aspect(implementation, attr_aspects=[], attrs=None, required_aspect_providers=[],
	//	provides=[], fragments=[], host_fragments=[], toolchains=[], doc='')
	Aspect(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#bind
	// None bind(name, actual=None)
	Bind(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#configuration_field
	// LateBoundDefault configuration_field(fragment, name)
	ConfigurationField(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#depset
	// depset depset(items=[], order="default", *, direct=None, transitive=None)
	Depset(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#existing_rules
	// unknown existing_rules()
	ExistingRules(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#existing_rules
	// None fail(msg=None, attr=None)
	Fail(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#PACKAGE_NAME
	// PACKAGE_NAME [string value]
	// TODO(vtl)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#provider
	// Provider provider(doc='', *, fields=None)
	Provider(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#register_execution_platforms
	// None register_execution_platforms(*platform_labels)
	RegisterExecutionPlatforms(args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#register_toolchains
	// None register_toolchains(*toolchain_labels)
	RegisterToolchains(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#REPOSITORY_NAME
	// REPOSITORY_NAME [string]
	// TODO(vtl)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#repository_rule
	// function repository_rule(implementation, *, attrs=None, local=False, environ=[], doc='')
	RepositoryRule(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#rule
	// function rule(implementation, test=False, attrs=None, outputs=None, executable=False,
	//     output_to_genfiles=False, fragments=[], host_fragments=[], _skylark_testable=False,
	//     toolchains=[], doc='', *, provides=[], execution_platform_constraints_allowed=False,
	//     exec_compatible_with=[], analysis_test=unbound, build_setting=None, cfg=None)
	Rule(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#select
	// unknown select(x, no_match_error='')
	//
	// https://docs.bazel.build/versions/master/be/functions.html#select
	// select(
	//     {conditionA: valuesA, conditionB: valuesB, ...},
	//     no_match_error = "custom message"
	// )
	Select(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/skylark/lib/globals.html#workspace
	// None workspace(name, managed_directories={})
	//
	// https://docs.bazel.build/versions/master/be/functions.html#workspace
	// workspace(name = "com_example_project")
	Workspace(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
}

// Build functions
// https://docs.bazel.build/versions/master/be/functions.html
type BuiltinsBuildFunctionsIface interface {
	// https://docs.bazel.build/versions/master/be/functions.html#package
	// package(default_deprecation, default_testonly, default_visibility, features)
	Package(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#package_group
	// package_group(name, packages, includes)
	PackageGroup(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#exports_files
	// exports_files([label, ...], visibility, licenses)
	ExportsFiles(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// https://docs.bazel.build/versions/master/be/functions.html#glob
	// glob(include, exclude=[], exclude_directories=1)
	Glob(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

	// Note: Select and Workspace are under "globals".
}

type BuiltinsIface interface {
	BuiltinsGlobalsIface
	BuiltinsBuildFunctionsIface

	// TODO(vtl): More (e.g., rules).
}

func getBuiltinsImpl(thread *starlark.Thread) BuiltinsIface {
	return GetContext(thread).BuiltinsImpl
}

func MakeInitialGlobals(ctx *Context) starlark.StringDict {
	// TODO(vtl): Customize this.
	return starlark.StringDict{
		"analysis_test_transition": starlark.NewBuiltin("analysis_test_transition",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AnalysisTestTransition(args, kwargs)
			}),
		"aspect": starlark.NewBuiltin("aspect",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Aspect(args, kwargs)
			}),
		"bind": starlark.NewBuiltin("bind",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Bind(args, kwargs)
			}),
		"configuration_field": starlark.NewBuiltin("configuration_field",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).ConfigurationField(args, kwargs)
			}),
		"depset": starlark.NewBuiltin("depset",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Provider(args, kwargs)
			}),
		"existing_rules": starlark.NewBuiltin("existing_rules",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).ExistingRules(args, kwargs)
			}),
		"fail": starlark.NewBuiltin("fail",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Fail(args, kwargs)
			}),
		// TODO(vtl): PACKAGE_NAME
		"provider": starlark.NewBuiltin("provider",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Provider(args, kwargs)
			}),
		"register_execution_platforms": starlark.NewBuiltin("register_execution_platforms",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).RegisterExecutionPlatforms(args,
					kwargs)
			}),
		"register_toolchains": starlark.NewBuiltin("register_toolchains",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).RegisterToolchains(args, kwargs)
			}),
		// TODO(vtl): REPOSITORY_NAME
		"repository_rule": starlark.NewBuiltin("repository_rule",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).RepositoryRule(args, kwargs)
			}),
		"rule": starlark.NewBuiltin("rule",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Rule(args, kwargs)
			}),
		"select": starlark.NewBuiltin("select",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Select(args, kwargs)
			}),
		"workspace": starlark.NewBuiltin("workspace",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Workspace(args, kwargs)
			}),
		// TODO(vtl): The rest of them.
		// Build functions
		// https://docs.bazel.build/versions/master/be/functions.html
		"package": starlark.NewBuiltin("package",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Package(args, kwargs)
			}),
		"package_group": starlark.NewBuiltin("package_group",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).PackageGroup(args, kwargs)
			}),
		"exports_files": starlark.NewBuiltin("exports_files",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).ExportsFiles(args, kwargs)
			}),
		"glob": starlark.NewBuiltin("glob",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).Glob(args, kwargs)
			}),
		// TODO(vtl): More (e.g., globals, rules).
	}
}

/*

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
