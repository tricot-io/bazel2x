// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package workspace_rules contains implementations of Bazel workspace rules (only callable from
// WORKSPACE files).
package workspace_rules

import (
	"fmt"

	"go.starlark.net/starlark"

	"src.tricot.io/public/bazel2x/bazel/core"
)

// TODO(vtl): Mostly copy-pasta of rules.newRule.
func newWorkspaceRule(ruleName string, impl func(ctx core.Context, args starlark.Tuple,
	kwargs []starlark.Tuple) error) *starlark.Builtin {

	return starlark.NewBuiltin(ruleName, func(thread *starlark.Thread, _ *starlark.Builtin,
		args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

		ctx := core.GetContext(thread)

		if ctx.FileType() != core.FileTypeWorkspace {
			return starlark.None, fmt.Errorf(
				"%v: %v: workspace rule can only be called from a WORKSPACE file",
				ctx.Label(), ruleName)
		}

		err := impl(ctx, args, kwargs)
		if err != nil {
			return starlark.None, fmt.Errorf("%v: %v: %v", ctx.Label(), ruleName, err)
		}

		return starlark.None, nil
	})
}
