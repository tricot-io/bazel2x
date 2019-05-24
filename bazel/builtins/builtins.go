// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package builtins provides the initial globals (builtins) for executing a Bazel file.
package builtins

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/builtins/functions"
	"bazel2cmake/bazel/builtins/rules"
	"bazel2cmake/bazel/builtins/workspace_rules"
	"bazel2cmake/bazel/core"
)

// InitialGlobals returns the initial globals (builtins) for executing a Bazel file.
func InitialGlobals(ctx core.Context) starlark.StringDict {
	return starlark.StringDict{
		// Globals
		// https://docs.bazel.build/versions/master/skylark/lib/globals.html
		"analysis_test_transition": functions.NotImplemented("analysis_test_transition"),
		"aspect":                   functions.NotImplemented("aspect"),
		// Note: bind is under workspace rules (below).
		"configuration_field": functions.NotImplemented("configuration_field"),
		"depset":              functions.NotImplemented("depset"),
		"existing_rules":      functions.NotImplemented("existing_rules"),
		"fail":                functions.NotImplemented("fail"),
		// TODO(vtl): Maybe this should be delegated somehow.
		"PACKAGE_NAME": starlark.String(string(ctx.Label().Package)),
		"provider":     functions.NotImplemented("provider"),
		"register_execution_platforms": functions.NotImplemented(
			"register_execution_platforms"),
		"register_toolchains": functions.NotImplemented("register_toolchains"),
		// TODO(vtl): Maybe this should be delegated somehow.
		"REPOSITORY_NAME": starlark.String("@" + string(ctx.Label().Workspace)),
		"repository_rule": functions.NotImplemented("repository_rule"),
		"rule":            functions.NotImplemented("rule"),
		"select":          functions.NotImplemented("select"),
		"workspace":       functions.NotImplemented("workspace"),

		// Functions
		// https://docs.bazel.build/versions/master/be/functions.html
		"package":       functions.NotImplemented("package"),
		"package_group": functions.NotImplemented("package_group"),
		"exports_files": functions.NotImplemented("exports_files"),
		"glob":          functions.NotImplemented("glob"),
		// Note: Select and Workspace are under "globals".

		// Android Rules
		// https://docs.bazel.build/versions/master/be/android.html
		"android_binary":  rules.NotImplemented("android_binary"),
		"aar_import":      rules.NotImplemented("aar_import"),
		"android_library": rules.NotImplemented("android_library"),
		"android_instrumentation_test": rules.NotImplemented(
			"android_instrumentation_test"),
		"android_local_test":     rules.NotImplemented("android_local_test"),
		"android_device":         rules.NotImplemented("android_device"),
		"android_ndk_repository": rules.NotImplemented("android_ndk_repository"),
		"android_sdk_repository": rules.NotImplemented("android_sdk_repository"),

		// C/C++ Rules
		// https://docs.bazel.build/versions/master/be/c-cpp.html
		"cc_binary":          rules.CcBinary,
		"cc_import":          rules.NotImplemented("cc_import"),
		"cc_library":         rules.CcLibrary,
		"cc_proto_library":   rules.NotImplemented("cc_proto_library"),
		"fdo_prefetch_hints": rules.NotImplemented("fdo_prefetch_hints"),
		"fdo_profile":        rules.NotImplemented("fdo_profile"),
		"cc_test":            rules.CcTest,
		"cc_toolchain":       rules.NotImplemented("cc_toolchain"),
		"cc_toolchain_suite": rules.NotImplemented("cc_toolchain_suite"),

		// Java Rules
		// https://docs.bazel.build/versions/master/be/java.html
		"java_binary":                rules.NotImplemented("java_binary"),
		"java_import":                rules.NotImplemented("java_import"),
		"java_library":               rules.NotImplemented("java_library"),
		"java_lite_proto_library":    rules.NotImplemented("java_lite_proto_library"),
		"java_proto_library":         rules.NotImplemented("java_proto_library"),
		"java_test":                  rules.NotImplemented("java_test"),
		"java_package_configuration": rules.NotImplemented("java_package_configuration"),
		"java_plugin":                rules.NotImplemented("java_plugin"),
		"java_runtime":               rules.NotImplemented("java_runtime"),
		"java_toolchain":             rules.NotImplemented("java_toolchain"),

		// Objective-C Rules
		// https://docs.bazel.build/versions/master/be/objective-c.html
		"apple_binary":         rules.NotImplemented("apple_binary"),
		"apple_static_library": rules.NotImplemented("apple_static_library"),
		"j2objc_library":       rules.NotImplemented("j2objc_library"),
		"objc_import":          rules.NotImplemented("objc_import"),
		"objc_library":         rules.NotImplemented("objc_library"),
		"objc_proto_library":   rules.NotImplemented("objc_proto_library"),

		// Protocol Buffer Rules
		// https://docs.bazel.build/versions/master/be/protocol-buffer.html
		"proto_lang_toolchain": rules.NotImplemented("proto_lang_toolchain"),
		"proto_library":        rules.NotImplemented("proto_library"),

		// Python Rules
		// https://docs.bazel.build/versions/master/be/python.html
		"py_binary":  rules.NotImplemented("py_binary"),
		"py_library": rules.NotImplemented("py_library"),
		"py_test":    rules.NotImplemented("py_test"),
		"py_runtime": rules.NotImplemented("py_runtime"),

		// Shell Rules
		// https://docs.bazel.build/versions/master/be/shell.html
		"sh_binary":  rules.NotImplemented("sh_binary"),
		"sh_library": rules.NotImplemented("sh_library"),
		"sh_test":    rules.NotImplemented("sh_test"),

		// Extra Actions Rules
		// https://docs.bazel.build/versions/master/be/extra-actions.html
		"action_listener": rules.NotImplemented("action_listener"),
		"extra_action":    rules.NotImplemented("extra_action"),

		// General Rules
		// https://docs.bazel.build/versions/master/be/general.html
		"filegroup":      rules.NotImplemented("filegroup"),
		"genquery":       rules.NotImplemented("genquery"),
		"test_suite":     rules.NotImplemented("test_suite"),
		"alias":          rules.NotImplemented("alias"),
		"config_setting": rules.NotImplemented("config_setting"),
		"genrule":        rules.NotImplemented("genrule"),

		// Platform Rules
		// https://docs.bazel.build/versions/master/be/platform.html
		"constraint_setting": rules.NotImplemented("constraint_setting"),
		"constraint_value":   rules.NotImplemented("constraint_value"),
		"platform":           rules.NotImplemented("platform"),
		"toolchain":          rules.NotImplemented("toolchain"),

		// Workspace Rules
		// https://docs.bazel.build/versions/master/be/workspace.html
		"bind":                 workspace_rules.NotImplemented("bind"),
		"local_repository":     workspace_rules.NotImplemented("local_repository"),
		"maven_jar":            workspace_rules.NotImplemented("maven_jar"),
		"maven_server":         workspace_rules.NotImplemented("maven_server"),
		"new_local_repository": workspace_rules.NotImplemented("new_local_repository"),
		"xcode_config":         workspace_rules.NotImplemented("xcode_config"),
		"xcode_version":        workspace_rules.NotImplemented("xcode_version"),
	}
}
