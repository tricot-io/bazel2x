// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

type NoOpBuiltinsImpl struct{}

func (self *NoOpBuiltinsImpl) Select(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Workspace(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Package(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) PackageGroup(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ExportsFiles(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Glob(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsIface = (*NoOpBuiltinsImpl)(nil)
