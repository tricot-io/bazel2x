// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

// NotImplemented is used for Bazel rules that we haven't implemented (yet).
func NotImplemented(ruleName string) *starlark.Builtin {
	return newRule(ruleName,
		func(ctx core.Context, args starlark.Tuple, kwargs []starlark.Tuple) error {

		return nil
	})
}
