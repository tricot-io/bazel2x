// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package workspace_rules

import (
	"fmt"

	"go.starlark.net/starlark"

	builtins_args "bazel2cmake/bazel/builtins/args"
	"bazel2cmake/bazel/core"
)

// WorkspaceArgs contains the argument for the workspace Bazel function.
type WorkspaceArgs struct {
	Name *string `bazel:"name!"`
}

var _ builtins_args.ProcessArgsTarget = (*WorkspaceArgs)(nil)

func (self *WorkspaceArgs) DidProcessArgs(ctx core.Context) error {
	if !core.WorkspaceName(*self.Name).IsValid() {
		return fmt.Errorf("%v is not a valid workspace name", *self.Name)
	}
	return nil
}

// Workspace implements the Bazel workspace function.
var Workspace = newWorkspaceRule("workspace",
	func(ctx core.Context, args starlark.Tuple, kwargs []starlark.Tuple) error {

		target := &WorkspaceArgs{}
		if err := builtins_args.ProcessArgs(args, kwargs, ctx, target); err != nil {
			return err
		}
		ctx.SetWorkspaceName(core.WorkspaceName(*target.Name))
		return nil
	})
