// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core // import "src.tricot.io/public/bazel2x/bazel/core"

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
type PackageTargets struct {
	TargetList    []Target
	TargetsByName map[TargetName]Target
}

// Add adds a target to the package.
func (self *PackageTargets) Add(target Target) error {
	label := target.Label()
	if _, alreadyExists := self.TargetsByName[label.Target]; alreadyExists {
		return fmt.Errorf("target %v already exists", label)
	}
	self.TargetList = append(self.TargetList, target)
	self.TargetsByName[label.Target] = target
	return nil
}

// WorkspaceTargets contains all the targets in a workspace.
type WorkspaceTargets map[PackageName]*PackageTargets

// AddPackage adds a package to the workspace.
func (self WorkspaceTargets) AddPackage(packageName PackageName) {
	if _, alreadyExists := self[packageName]; alreadyExists {
		panic(packageName)
	}
	self[packageName] = &PackageTargets{[]Target{}, make(map[TargetName]Target)}
}

// Add adds a target to the workspace.
func (self WorkspaceTargets) Add(target Target) error {
	return self[target.Label().Package].Add(target)
}

// BuildTargets contains all the targets in a build.
type BuildTargets map[WorkspaceName]WorkspaceTargets

func (self BuildTargets) AddPackage(workspaceName WorkspaceName, packageName PackageName) {
	workspaceTargets, ok := self[workspaceName]
	if !ok {
		workspaceTargets = make(WorkspaceTargets)
		self[workspaceName] = workspaceTargets
	}
	workspaceTargets.AddPackage(packageName)
}

// Add adds a target to the build.
func (self BuildTargets) Add(target Target) error {
	return self[target.Label().Workspace].Add(target)
}
