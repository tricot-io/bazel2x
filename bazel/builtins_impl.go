// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"bazel2cmake/bazel/core"
)

// BuiltinsImpl implements Builtins.
type BuiltinsImpl struct {
	NoOpBuiltinsGlobalsImpl
	NoOpBuiltinsFunctionsImpl
	NoOpBuiltinsAndroidRulesImpl
	NoOpBuiltinsCcRulesImpl
	NoOpBuiltinsJavaRulesImpl
	NoOpBuiltinsObjCRulesImpl
	NoOpBuiltinsProtoBufRulesImpl
	NoOpBuiltinsPythonRulesImpl
	NoOpBuiltinsShellRulesImpl
	NoOpBuiltinsExtraActionsRulesImpl
	NoOpBuiltinsGeneralRulesImpl
	NoOpBuiltinsPlatformRulesImpl
	NoOpBuiltinsWorkspaceRulesImpl

	ctx core.Context
}

var _ Builtins = (*BuiltinsImpl)(nil)

func NewBuiltinsImpl(ctx core.Context) *BuiltinsImpl {
	return &BuiltinsImpl{ctx: ctx}
}
