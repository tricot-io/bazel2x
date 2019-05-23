// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type Context interface {
	Label() core.Label
	FileType() core.FileType
}

const contextKey = "bazel2make-bazel-context"

func SetContext(thread *starlark.Thread, ctx Context) {
	thread.SetLocal(contextKey, ctx)
}

func GetContext(thread *starlark.Thread) Context {
	return thread.Local(contextKey).(Context)
}

type ContextImpl struct {
	build *Build

	label    core.Label
	fileType core.FileType

	// TODO(vtl): Remove this.
	builtinsImpl Builtins
}

var _ Context = (*ContextImpl)(nil)

func (self *ContextImpl) Label() core.Label {
	return self.label
}

func (self *ContextImpl) FileType() core.FileType {
	return self.fileType
}

func (self *ContextImpl) CreateThread(label core.Label, fileType core.FileType) *starlark.Thread {
	return CreateThread(self.build, label, fileType)
}

func (self *ContextImpl) MakeInitialGlobals() starlark.StringDict {
	return MakeInitialGlobals(self)
}

func GetContextImpl(thread *starlark.Thread) *ContextImpl {
	return GetContext(thread).(*ContextImpl)
}

func CreateThread(build *Build, label core.Label, fileType core.FileType) *starlark.Thread {
	// Create the thread.
	thread := &starlark.Thread{Name: "exec " + label.String(), Load: Load}

	// Create a new context (with the same loader).
	ctx := &ContextImpl{
		build:        build,
		label:        label,
		fileType:     fileType,
	}
	ctx.builtinsImpl = NewBuiltinsImpl(ctx)
	// And attach it to the thread.
	SetContext(thread, ctx)

	return thread
}
