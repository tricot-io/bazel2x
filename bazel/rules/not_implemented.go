// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type NotImplementedTarget struct {}

var _ core.Target = (*NotImplementedTarget)(nil)

func (self *NotImplementedTarget) String() string {
	// You probably should never see this.
	return "NOT IMPLEMENTED"
}

func (self *NotImplementedTarget) Label() core.Label {
	return core.Label{}
}

// NotImplemented is used for Bazel rules that we haven't implemented (yet).
func NotImplemented(thread *starlark.Thread, _ *starlark.Builtin, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}
