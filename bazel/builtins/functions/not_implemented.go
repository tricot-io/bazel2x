// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package functions

import (
	"go.starlark.net/starlark"
)

// NotImplemented is used for Bazel functions (including rules) that we haven't implemented (yet).
//
// Note: While some functions are executed at the top-level (for their side effects), some are used
// for their return values, in which case the None return will cause execution to fail.
func NotImplemented(fnName string) *starlark.Builtin {
	return starlark.NewBuiltin(fnName, func(thread *starlark.Thread, _ *starlark.Builtin,
		args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

		return starlark.None, nil
	})
}

// NotImplemented2 is used for Bazel functions that we haven't implemented (yet), which are required
// to return a function.
func NotImplemented2(fnName string) *starlark.Builtin {
	return starlark.NewBuiltin(fnName, func(thread *starlark.Thread, _ *starlark.Builtin,
		args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

		return NotImplemented(fnName + "_result"), nil
	})
}
