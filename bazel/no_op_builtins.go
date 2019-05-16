// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

type NoOpBuiltinsImpl struct{}

func (self *NoOpBuiltinsImpl) Package(thread *starlark.Thread, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) PackageGroup(thread *starlark.Thread, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ExportsFiles(thread *starlark.Thread, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Glob(thread *starlark.Thread, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Select(thread *starlark.Thread, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Workspace(thread *starlark.Thread, args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {

	return starlark.None, nil
}

var _ BuiltinsIface = (*NoOpBuiltinsImpl)(nil)
