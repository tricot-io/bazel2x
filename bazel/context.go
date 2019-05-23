// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type Context struct {
	Build *Build

	Label    core.Label
	FileType core.FileType
	Thread   *starlark.Thread

	BuiltinsImpl Builtins
}

func (ctx *Context) CreateThread(label core.Label, fileType core.FileType) *starlark.Thread {
	return CreateThread(ctx.Build, label, fileType)
}

func (ctx *Context) MakeInitialGlobals() starlark.StringDict {
	return MakeInitialGlobals(ctx)
}

const contextKey = "bazel2make-bazel-context"

func CreateThread(build *Build, label core.Label, fileType core.FileType) *starlark.Thread {
	// Create the thread.
	thread := &starlark.Thread{Name: "exec " + label.String(), Load: Load}

	// Create a new context (with the same loader).
	ctx := &Context{
		Build:        build,
		Label:        label,
		FileType:     fileType,
		Thread:       thread,
	}
	ctx.BuiltinsImpl = NewBuiltinsImpl(ctx)
	// And attach it to the thread.
	thread.SetLocal(contextKey, ctx)

	return thread
}

func GetContext(thread *starlark.Thread) *Context {
	return thread.Local(contextKey).(*Context)
}
