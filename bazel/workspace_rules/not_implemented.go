// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package workspace_rules

import (
	"go.starlark.net/starlark"
)

// NotImplemented is used for Bazel workspace rules that we haven't implemented (yet).
func NotImplemented(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	// TODO(vtl): Check that we're in a WORKSPACE file.

	return starlark.None, nil
}
