// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package rules contains implementations of Bazel build rules (only callable from BUILD[.bazel]
// files).
package rules

import (
	"fmt"

	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

// newRule creates a new rule(-like) starlark.Builtin.
func newRule(ruleName string, impl func(ctx core.Context, args starlark.Tuple,
	kwargs []starlark.Tuple) error) *starlark.Builtin {

	return starlark.NewBuiltin(ruleName, func(thread *starlark.Thread, _ *starlark.Builtin,
		args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

		ctx := core.GetContext(thread)

		if ctx.FileType() != core.FileTypeBuild {
			return starlark.None, fmt.Errorf(
				"%v: %v: rule can only be called from a BUILD[.bazel] file",
				ctx.Label(), ruleName)
		}

		err := impl(ctx, args, kwargs)
		if err != nil {
			return starlark.None, fmt.Errorf("%v: %v: %v", ctx.Label(), ruleName, err)
		}

		return starlark.None, nil
	})
}
