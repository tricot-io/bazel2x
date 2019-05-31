// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package builtins provides the initial globals (builtins) for executing a Bazel file.
package builtins // import "src.tricot.io/public/bazel2x/bazel/builtins"

import (
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"

	"src.tricot.io/public/bazel2x/bazel/builtins/functions"
	"src.tricot.io/public/bazel2x/bazel/builtins/rules"
	"src.tricot.io/public/bazel2x/bazel/builtins/workspace_rules"
	"src.tricot.io/public/bazel2x/bazel/core"
)

func starlarkUnion(dicts ...starlark.StringDict) starlark.StringDict {
	rv := starlark.StringDict{}
	for _, dict := range dicts {
		for k, v := range dict {
			rv[k] = v
		}
	}
	return rv
}

// rules are rules (exposed as globals in BUILD files and via native in .bzl (and BUILD) files).
var rulesGlobals = starlark.StringDict{
	// Android Rules
	// https://docs.bazel.build/versions/master/be/android.html
	"android_binary":               rules.NotImplemented("android_binary"),
	"aar_import":                   rules.NotImplemented("aar_import"),
	"android_library":              rules.NotImplemented("android_library"),
	"android_instrumentation_test": rules.NotImplemented("android_instrumentation_test"),
	"android_local_test":           rules.NotImplemented("android_local_test"),
	"android_device":               rules.NotImplemented("android_device"),
	"android_ndk_repository":       rules.NotImplemented("android_ndk_repository"),
	"android_sdk_repository":       rules.NotImplemented("android_sdk_repository"),

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
}

// commonGlobals are globals that are common to BUILD, .bzl, and WORKSPACE files.
var commonGlobals = starlark.StringDict{
	"depset": functions.NotImplemented("depset"),
	"fail":   functions.NotImplemented("fail"),
	"select": functions.NotImplemented("select"),
}

// buildAndbzlCommonGlobals are globals that are common to BUILD and .bzl files.
var buildAndbzlCommonGlobals = starlark.StringDict{
	"configuration_field": functions.NotImplemented("configuration_field"),
	"native": &starlarkstruct.Module{
		Name: "native",
		// Note: Rules will be added separately.
		Members: starlarkUnion(
			starlark.StringDict{
				// Non-rule Members
				// https://docs.bazel.build/versions/master/skylark/lib/native.html
				"existing_rule": functions.NotImplemented("existing_rule"),
				"existing_rules": functions.NotImplementedRv("existing_rules",
					&starlark.List{}),
				"exports_files":   functions.NotImplemented("exports_files"),
				"glob":            functions.NotImplemented("glob"),
				"package_group":   functions.NotImplemented("package_group"),
				"package_name":    functions.NotImplemented("package_name"),
				"repository_name": functions.NotImplemented("repository_name"),
			},
			rulesGlobals,
		),
	},
}

// buildGlobals are globals for BUILD files.
//
// These (and others) are variously documented in:
//   https://docs.bazel.build/versions/master/skylark/lib/globals.html
//   https://docs.bazel.build/versions/master/be/functions.html
//   https://docs.bazel.build/versions/master/skylark/lib/native.html
//
// Notes:
//   analysis_test_transition is not supported.
//   PACKAGE_NAME has actually been "removed" from Bazel (using it now causes an error).
//   REPOSITORY_NAME has actually been "removed" from Bazel (using it now causes an error).
var buildGlobals = starlarkUnion(
	commonGlobals,
	buildAndbzlCommonGlobals,
	rulesGlobals,
	starlark.StringDict{
		"existing_rules":  functions.NotImplementedRv("existing_rules", &starlark.List{}),
		"exports_files":   functions.NotImplemented("exports_files"),
		"glob":            functions.NotImplemented("glob"),
		"package":         functions.NotImplemented("package"),
		"package_group":   functions.NotImplemented("package_group"),
		"package_name":    functions.NotImplemented("package_name"),
		"repository_name": functions.NotImplemented("repository_name"),
	},
)

var bzlGlobals = starlarkUnion(
	commonGlobals,
	buildAndbzlCommonGlobals,
	starlark.StringDict{
		"aspect":   functions.NotImplemented("aspect"),
		"provider": functions.NotImplemented("provider"),
		"repository_rule": functions.NotImplementedRv("repository_rule",
			functions.NotImplemented("repository_rule_rv")),
		"rule": functions.NotImplementedRv("rule", functions.NotImplemented("rule_rv")),

		// https://docs.bazel.build/versions/master/skylark/lib/attr.html
		"attr": &starlarkstruct.Module{
			Name: "attr",
			Members: starlark.StringDict{
				"bool":     functions.NotImplemented("bool"),
				"int":      functions.NotImplemented("int"),
				"int_list": functions.NotImplemented("int_list"),
				"label":    functions.NotImplemented("label"),
				"label_keyed_string_dict": functions.NotImplemented(
					"label_keyed_string_dict"),
				"label_list":       functions.NotImplemented("label_list"),
				"license":          functions.NotImplemented("license"),
				"output":           functions.NotImplemented("output"),
				"output_list":      functions.NotImplemented("output_list"),
				"string":           functions.NotImplemented("string"),
				"string_dict":      functions.NotImplemented("string_dict"),
				"string_list":      functions.NotImplemented("string_list"),
				"string_list_dict": functions.NotImplemented("string_list_dict"),
			},
		},
	},
)

var workspaceGlobals = starlarkUnion(
	commonGlobals,
	starlark.StringDict{
		"register_execution_platforms": functions.NotImplemented(
			"register_execution_platforms"),
		"register_toolchains": functions.NotImplemented("register_toolchains"),
		"workspace":           workspace_rules.Workspace,

		// Workspace Rules
		// https://docs.bazel.build/versions/master/be/workspace.html
		"bind":                 workspace_rules.NotImplemented("bind"),
		"local_repository":     workspace_rules.NotImplemented("local_repository"),
		"maven_jar":            workspace_rules.NotImplemented("maven_jar"),
		"maven_server":         workspace_rules.NotImplemented("maven_server"),
		"new_local_repository": workspace_rules.NotImplemented("new_local_repository"),
		"xcode_config":         workspace_rules.NotImplemented("xcode_config"),
		"xcode_version":        workspace_rules.NotImplemented("xcode_version"),
	},
)

// InitialGlobals returns the initial globals (builtins) for executing a Bazel file.
func InitialGlobals(fileType core.FileType) starlark.StringDict {
	switch fileType {
	case core.FileTypeBuild:
		return buildGlobals
	case core.FileTypeBzl:
		return bzlGlobals
	case core.FileTypeWorkspace:
		return workspaceGlobals
	default:
		panic(fileType)
	}
}
