// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core

import (
	"go.starlark.net/starlark"
)

// Context contains/makes accessible per-starlark-thread context data.
type Context interface {
	// Label returns a label indicating the name of the build file.
	Label() Label

	// FileType returns the build file's type.
	FileType() FileType

	BuildTargets() BuildTargets
}

const contextKey = "bazel2make-bazel-context"

// SetContext sets the Context of a starlark thread.
func SetContext(thread *starlark.Thread, ctx Context) {
	thread.SetLocal(contextKey, ctx)
}

// GetContext gets the Context of a starlark thread.
func GetContext(thread *starlark.Thread) Context {
	return thread.Local(contextKey).(Context)
}
