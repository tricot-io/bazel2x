// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

// Globals
// https://docs.bazel.build/versions/master/skylark/lib/globals.html
type BuiltinsGlobalsIface interface {
	AnalysisTestTransition(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Aspect(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Bind(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	ConfigurationField(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Depset(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	ExistingRules(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Fail(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	// TODO(vtl): Helper for PACKAGE_NAME?
	Provider(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	RegisterExecutionPlatforms(args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)
	RegisterToolchains(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	// TODO(vtl): Helper for REPOSITORY_NAME?
	RepositoryRule(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Rule(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Select(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Workspace(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
}

// Build functions
// https://docs.bazel.build/versions/master/be/functions.html
type BuiltinsBuildFunctionsIface interface {
	Package(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	PackageGroup(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	ExportsFiles(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Glob(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	// Note: Select and Workspace are under "globals".
}

// Android rules
// https://docs.bazel.build/versions/master/be/android.html#android-rules
type BuiltinsAndroidRulesIface interface {
	AndroidBinary(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	AarImport(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	AndroidLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	AndroidInstrumentationTest(args starlark.Tuple, kwargs []starlark.Tuple) (
		starlark.Value, error)
	AndroidLocalTest(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	AndroidDevice(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	AndroidNdkRepository(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	AndroidSdkRepository(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
}

// C/C++ rules
// android_sdk_repository(name, api_level, build_tools_version, path)
type BuiltinsCcRulesIface interface {
	CcBinary(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	CcImport(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	CcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	CcProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	FdoPrefetchHints(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	FdoProfile(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	CcTest(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	CcToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	CcToolchainSuite(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
}

type BuiltinsIface interface {
	BuiltinsGlobalsIface
	BuiltinsBuildFunctionsIface
	BuiltinsAndroidRulesIface
	BuiltinsCcRulesIface

	// TODO(vtl): More (e.g., rules).
}

func getBuiltinsImpl(thread *starlark.Thread) BuiltinsIface {
	return GetContext(thread).BuiltinsImpl
}

func MakeInitialGlobals(ctx *Context) starlark.StringDict {
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
		// TODO(vtl): Maybe this should be delegated somehow.
		"PACKAGE_NAME": starlark.String(string(ctx.Label.Package)),
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
		// TODO(vtl): Maybe this should be delegated somehow.
		"REPOSITORY_NAME": starlark.String("@"+string(ctx.Label.Workspace)),
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
		"android_binary": starlark.NewBuiltin("android_binary",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidBinary(args, kwargs)
			}),
		"aar_import": starlark.NewBuiltin("aar_import",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AarImport(args, kwargs)
			}),
		"android_library": starlark.NewBuiltin("android_library",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidLibrary(args, kwargs)
			}),
		"android_instrumentation_test": starlark.NewBuiltin("android_instrumentation_test",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidInstrumentationTest(args,
					kwargs)
			}),
		"android_local_test": starlark.NewBuiltin("android_local_test",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidLocalTest(args, kwargs)
			}),
		"android_device": starlark.NewBuiltin("android_device",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidDevice(args, kwargs)
			}),
		"android_ndk_repository": starlark.NewBuiltin("android_ndk_repository",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidNdkRepository(args, kwargs)
			}),
		"android_sdk_repository": starlark.NewBuiltin("android_sdk_repository",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).AndroidSdkRepository(args, kwargs)
			}),
		"cc_binary": starlark.NewBuiltin("cc_binary",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcBinary(args, kwargs)
			}),
		"cc_import": starlark.NewBuiltin("cc_import",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcImport(args, kwargs)
			}),
		"cc_library": starlark.NewBuiltin("cc_library",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcLibrary(args, kwargs)
			}),
		"cc_proto_library": starlark.NewBuiltin("cc_proto_library",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcProtoLibrary(args, kwargs)
			}),
		"fdo_prefetch_hints": starlark.NewBuiltin("fdo_prefetch_hints",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).FdoPrefetchHints(args, kwargs)
			}),
		"fdo_profile": starlark.NewBuiltin("fdo_profile",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).FdoProfile(args, kwargs)
			}),
		"cc_test": starlark.NewBuiltin("cc_test",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcTest(args, kwargs)
			}),
		"cc_toolchain": starlark.NewBuiltin("cc_toolchain",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcToolchain(args, kwargs)
			}),
		"cc_toolchain_suite": starlark.NewBuiltin("cc_toolchain_suite",
			func(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
				kwargs []starlark.Tuple) (starlark.Value, error) {
				return getBuiltinsImpl(thread).CcToolchainSuite(args, kwargs)
			}),
		// TODO(vtl): More rules.
	}
}

/*

# Rules

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
