// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

type FileType int

const (
	FileTypeBuild FileType = iota
	FileTypeBzl
	FileTypeWorkspace
)

type Context struct {
	Build    *Build

	Label    Label
	FileType FileType
	Thread   *starlark.Thread

	BuiltinsImpl BuiltinsIface
}

func (ctx *Context) CreateThread(label Label, fileType FileType) *starlark.Thread {
	return CreateThread(ctx.Build, label, fileType)
}

func (ctx *Context) MakeInitialGlobals() starlark.StringDict {
	return MakeInitialGlobals(ctx)
}

const contextKey = "bazel2make-bazel-context"

func CreateThread(build *Build, label Label, fileType FileType) *starlark.Thread {
	// Create the thread.
	thread := &starlark.Thread{Name: "exec " + label.String(), Load: Load}

	// Create a new context (with the same loader).
	ctx := &Context{
		Build:    build,
		Label:    label,
		FileType: fileType,
		Thread:   thread,
		// TODO(vtl): Choose BuiltinsImpl based on fileType.
		BuiltinsImpl: &NoOpBuiltinsImpl{},
	}
	// And attach it to the thread.
	thread.SetLocal(contextKey, ctx)

	return thread
}

func GetContext(thread *starlark.Thread) *Context {
	return thread.Local(contextKey).(*Context)
}
