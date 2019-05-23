// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core

import (
	"fmt"
)

// Target represents a Bazel target (or really potentially anything with an assigned label).
//
// TODO(vtl): Maybe add Deps() and DataDeps()?
type Target interface {
	fmt.Stringer
	Label() Label
}

// PackageTargets contains the targets in a package.
type PackageTargets map[TargetName]Target

// Add adds a target to the package.
func (self PackageTargets) Add(target Target) error {
	label := target.Label()
	if _, alreadyPresent := self[label.Target]; alreadyPresent {
		return fmt.Errorf("target %v already exists", label)
	}

	self[label.Target] = target
	return nil
}

// WorkspaceTargets contains all the targets in a workspace.
type WorkspaceTargets map[PackageName]PackageTargets

// Add adds a target to the workspace.
func (self WorkspaceTargets) Add(target Target) error {
	label := target.Label()
	packageTargets, ok := self[label.Package]
	if !ok {
		packageTargets = make(PackageTargets)
		self[label.Package] = packageTargets
	}
	return packageTargets.Add(target)
}

// BuildTargets contains all the targets in a build.
type BuildTargets map[WorkspaceName]WorkspaceTargets

// Add adds a target to the build.
func (self BuildTargets) Add(target Target) error {
	label := target.Label()
	workspaceTargets, ok := self[label.Workspace]
	if !ok {
		workspaceTargets = make(WorkspaceTargets)
		self[label.Workspace] = workspaceTargets
	}

	return workspaceTargets.Add(target)
}
