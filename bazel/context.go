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
	FileType FileType
	Label    Label
	Thread   *starlark.Thread
}

const contextKey = "bazel2make-bazel-context"

// TODO(vtl): More args (before thread).
func InitThread(fileType FileType, label Label, builtinsImpl BuiltinsIface,
	thread *starlark.Thread) {

	ctx := &Context{
		FileType: fileType,
		Label:    label,
		Thread:   thread,
	}
	thread.SetLocal(contextKey, ctx)

	SetBuiltinsImpl(thread, builtinsImpl)
}

func GetContext(thread *starlark.Thread) *Context {
	return thread.Local(contextKey).(*Context)
}
