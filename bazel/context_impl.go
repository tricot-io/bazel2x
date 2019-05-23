// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type ContextImpl struct {
	build *Build

	label    core.Label
	fileType core.FileType

	// TODO(vtl): Remove this.
	builtinsImpl Builtins
}

var _ core.Context = (*ContextImpl)(nil)

func (self *ContextImpl) Label() core.Label {
	return self.label
}

func (self *ContextImpl) FileType() core.FileType {
	return self.fileType
}

func (self *ContextImpl) BuildTargets() core.BuildTargets {
	return self.build.BuildTargets
}

func (self *ContextImpl) CreateThread(label core.Label, fileType core.FileType) *starlark.Thread {
	return CreateThread(self.build, label, fileType)
}

func (self *ContextImpl) MakeInitialGlobals() starlark.StringDict {
	return MakeInitialGlobals(self)
}

func GetContextImpl(thread *starlark.Thread) *ContextImpl {
	return core.GetContext(thread).(*ContextImpl)
}

func CreateThread(build *Build, label core.Label, fileType core.FileType) *starlark.Thread {
	// Create the thread.
	thread := &starlark.Thread{Name: "exec " + label.String(), Load: Load}

	// Create a new context (with the same loader).
	ctx := &ContextImpl{
		build:    build,
		label:    label,
		fileType: fileType,
	}
	ctx.builtinsImpl = NewBuiltinsImpl(ctx)
	// And attach it to the thread.
	core.SetContext(thread, ctx)

	return thread
}
