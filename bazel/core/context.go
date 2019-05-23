// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core

import (
	"go.starlark.net/starlark"
)

type Context interface {
	Label() Label
	FileType() FileType
}

const contextKey = "bazel2make-bazel-context"

func SetContext(thread *starlark.Thread, ctx Context) {
	thread.SetLocal(contextKey, ctx)
}

func GetContext(thread *starlark.Thread) Context {
	return thread.Local(contextKey).(Context)
}
