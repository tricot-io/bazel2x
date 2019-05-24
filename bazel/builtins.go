// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
	"bazel2cmake/bazel/rules"
	"bazel2cmake/bazel/workspace_rules"
)

// Globals
// https://docs.bazel.build/versions/master/skylark/lib/globals.html
type BuiltinsGlobals interface {
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

// Functions
// https://docs.bazel.build/versions/master/be/functions.html
type BuiltinsFunctions interface {
	Package(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	PackageGroup(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	ExportsFiles(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	Glob(args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)
	// Note: Select and Workspace are under "globals".
}

// TODO(vtl): Probably should split this into rules and non-rules.
type Builtins interface {
	BuiltinsGlobals
	BuiltinsFunctions
}

func getBuiltinsImpl(thread *starlark.Thread) Builtins {
	return GetContextImpl(thread).builtinsImpl
}

func MakeInitialGlobals(ctx core.Context) starlark.StringDict {
	return starlark.StringDict{
		// Globals
		// https://docs.bazel.build/versions/master/skylark/lib/globals.html
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
		"PACKAGE_NAME": starlark.String(string(ctx.Label().Package)),
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
		"REPOSITORY_NAME": starlark.String("@" + string(ctx.Label().Workspace)),
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

		// Functions
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
		// Note: Select and Workspace are under "globals".

		// Android Rules
		// https://docs.bazel.build/versions/master/be/android.html
		"android_binary": starlark.NewBuiltin("android_binary", rules.NotImplemented),
		"aar_import": starlark.NewBuiltin("aar_import", rules.NotImplemented),
		"android_library": starlark.NewBuiltin("android_library", rules.NotImplemented),
		"android_instrumentation_test": starlark.NewBuiltin("android_instrumentation_test",
			rules.NotImplemented),
		"android_local_test": starlark.NewBuiltin("android_local_test",
			rules.NotImplemented),
		"android_device": starlark.NewBuiltin("android_device", rules.NotImplemented),
		"android_ndk_repository": starlark.NewBuiltin("android_ndk_repository",
			rules.NotImplemented),
		"android_sdk_repository": starlark.NewBuiltin("android_sdk_repository",
			rules.NotImplemented),

		// C/C++ Rules
		// https://docs.bazel.build/versions/master/be/c-cpp.html
		"cc_binary":        starlark.NewBuiltin("cc_binary", rules.CcBinary),
		"cc_import":        starlark.NewBuiltin("cc_import", rules.NotImplemented),
		"cc_library":       starlark.NewBuiltin("cc_library", rules.CcLibrary),
		"cc_proto_library": starlark.NewBuiltin("cc_proto_library", rules.NotImplemented),
		"fdo_prefetch_hints": starlark.NewBuiltin("fdo_prefetch_hints",
			rules.NotImplemented),
		"fdo_profile":  starlark.NewBuiltin("fdo_profile", rules.NotImplemented),
		"cc_test":      starlark.NewBuiltin("cc_test", rules.NotImplemented),
		"cc_toolchain": starlark.NewBuiltin("cc_toolchain", rules.NotImplemented),
		"cc_toolchain_suite": starlark.NewBuiltin("cc_toolchain_suite",
			rules.NotImplemented),

		// Java Rules
		// https://docs.bazel.build/versions/master/be/java.html
		"java_binary": starlark.NewBuiltin("java_binary", rules.NotImplemented),
		"java_import": starlark.NewBuiltin("java_import", rules.NotImplemented),
		"java_library": starlark.NewBuiltin("java_library", rules.NotImplemented),
		"java_lite_proto_library": starlark.NewBuiltin("java_lite_proto_library",
			rules.NotImplemented),
		"java_proto_library": starlark.NewBuiltin("java_proto_library",
			rules.NotImplemented),
		"java_test": starlark.NewBuiltin("java_test", rules.NotImplemented),
		"java_package_configuration": starlark.NewBuiltin("java_package_configuration",
			rules.NotImplemented),
		"java_plugin": starlark.NewBuiltin("java_plugin", rules.NotImplemented),
		"java_runtime": starlark.NewBuiltin("java_runtime", rules.NotImplemented),
		"java_toolchain": starlark.NewBuiltin("java_toolchain", rules.NotImplemented),

		// Objective-C Rules
		// https://docs.bazel.build/versions/master/be/objective-c.html
		"apple_binary": starlark.NewBuiltin("apple_binary", rules.NotImplemented),
		"apple_static_library": starlark.NewBuiltin("apple_static_library",
			rules.NotImplemented),
		"j2objc_library": starlark.NewBuiltin("j2objc_library", rules.NotImplemented),
		"objc_import": starlark.NewBuiltin("objc_import", rules.NotImplemented),
		"objc_library": starlark.NewBuiltin("objc_library", rules.NotImplemented),
		"objc_proto_library": starlark.NewBuiltin("objc_proto_library",
			rules.NotImplemented),

		// Protocol Buffer Rules
		// https://docs.bazel.build/versions/master/be/protocol-buffer.html
		"proto_lang_toolchain": starlark.NewBuiltin("proto_lang_toolchain",
			rules.NotImplemented),
		"proto_library": starlark.NewBuiltin("proto_library", rules.NotImplemented),

		// Python Rules
		// https://docs.bazel.build/versions/master/be/python.html
		"py_binary": starlark.NewBuiltin("py_binary", rules.NotImplemented),
		"py_library": starlark.NewBuiltin("py_library", rules.NotImplemented),
		"py_test": starlark.NewBuiltin("py_test", rules.NotImplemented),
		"py_runtime": starlark.NewBuiltin("py_runtime", rules.NotImplemented),

		// Shell Rules
		// https://docs.bazel.build/versions/master/be/shell.html
		"sh_binary": starlark.NewBuiltin("sh_binary", rules.NotImplemented),
		"sh_library": starlark.NewBuiltin("sh_library", rules.NotImplemented),
		"sh_test": starlark.NewBuiltin("sh_test", rules.NotImplemented),

		// Extra Actions Rules
		// https://docs.bazel.build/versions/master/be/extra-actions.html
		"action_listener": starlark.NewBuiltin("action_listener", rules.NotImplemented),
		"extra_action": starlark.NewBuiltin("extra_action", rules.NotImplemented),

		// General Rules
		// https://docs.bazel.build/versions/master/be/general.html
		"filegroup": starlark.NewBuiltin("filegroup", rules.NotImplemented),
		"genquery": starlark.NewBuiltin("genquery", rules.NotImplemented),
		"test_suite": starlark.NewBuiltin("test_suite", rules.NotImplemented),
		"alias": starlark.NewBuiltin("alias", rules.NotImplemented),
		"config_setting": starlark.NewBuiltin("config_setting", rules.NotImplemented),
		"genrule": starlark.NewBuiltin("genrule", rules.NotImplemented),

		// Platform Rules
		// https://docs.bazel.build/versions/master/be/platform.html
		"constraint_setting": starlark.NewBuiltin("constraint_setting",
			rules.NotImplemented),
		"constraint_value": starlark.NewBuiltin("constraint_value", rules.NotImplemented),
		"platform": starlark.NewBuiltin("platform", rules.NotImplemented),
		"toolchain": starlark.NewBuiltin("toolchain", rules.NotImplemented),

		// Workspace Rules
		// https://docs.bazel.build/versions/master/be/workspace.html
		// Note: Bind is under "Globals" (above).
		"local_repository": starlark.NewBuiltin("local_repository",
			workspace_rules.NotImplemented),
		"maven_jar": starlark.NewBuiltin("maven_jar", workspace_rules.NotImplemented),
		"maven_server": starlark.NewBuiltin("maven_server", workspace_rules.NotImplemented),
		"new_local_repository": starlark.NewBuiltin("new_local_repository",
			workspace_rules.NotImplemented),
		"xcode_config": starlark.NewBuiltin("xcode_config", workspace_rules.NotImplemented),
		"xcode_version": starlark.NewBuiltin("xcode_version",
			workspace_rules.NotImplemented),
	}
}
