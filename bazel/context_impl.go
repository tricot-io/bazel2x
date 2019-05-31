// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel // import "src.tricot.io/public/bazel2x/bazel"

import (
	"fmt"

	"go.starlark.net/starlark"

	"src.tricot.io/public/bazel2x/bazel/core"
)

type ContextImpl struct {
	build *Build

	label    core.Label
	fileType core.FileType
}

var _ core.Context = (*ContextImpl)(nil)

func (self *ContextImpl) WorkspaceName() core.WorkspaceName {
	return self.build.WorkspaceName
}

func (self *ContextImpl) SetWorkspaceName(workspaceName core.WorkspaceName) error {
	if self.build.WorkspaceName != "" {
		return fmt.Errorf("workspace name can only be set once")
	}
	self.build.WorkspaceName = workspaceName
	return nil
}

func (self *ContextImpl) Label() core.Label {
	return self.label
}

func (self *ContextImpl) FileType() core.FileType {
	return self.fileType
}

func (self *ContextImpl) BuildTargets() core.BuildTargets {
	return self.build.BuildTargets
}

// TODO(vtl): Maybe get rid of this. We only need this when we need to access the Build, which is
// when we need it to do a load, but maybe that should be a part of Context.
func GetContextImpl(thread *starlark.Thread) *ContextImpl {
	return core.GetContext(thread).(*ContextImpl)
}

// TODO(vtl): Move this?
func createThread(build *Build, label core.Label, fileType core.FileType) *starlark.Thread {
	// Create the thread.
	thread := &starlark.Thread{Name: "exec " + label.String(), Load: load}

	// Create a new context (with the same loader) and attach it to the thread.
	ctx := &ContextImpl{
		build:    build,
		label:    label,
		fileType: fileType,
	}
	core.SetContext(thread, ctx)

	return thread
}
