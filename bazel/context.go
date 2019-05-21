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
	Label    Label
	FileType FileType
	Loader   *Loader
	Thread   *starlark.Thread
}

func (ctx *Context) CreateThread(label Label, fileType FileType) *starlark.Thread {
	return CreateThread(ctx.Loader, label, fileType)
}

const contextKey = "bazel2make-bazel-context"

func CreateThread(loader *Loader, label Label, fileType FileType) *starlark.Thread {
	// Create the thread.
	thread := &starlark.Thread{Name: "exec " + label.String(), Load: Load}

	// Create a new context (with the same loader).
	ctx := &Context{
		Label:    label,
		FileType: fileType,
		Loader:   loader,
		Thread:   thread,
	}
	// And attach it to the thread.
	thread.SetLocal(contextKey, ctx)

	// Set the thread's builtins.
	// TODO(vtl): Choose builtinsImpl based on fileType.
	builtinsImpl := &NoOpBuiltinsImpl{}
	SetBuiltinsImpl(thread, builtinsImpl)

	return thread
}

func GetContext(thread *starlark.Thread) *Context {
	return thread.Local(contextKey).(*Context)
}
