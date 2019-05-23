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
	Label() Label
}

// TODO(vtl): Move the below to another file?

type PackageTargets map[TargetName]Target

func (self PackageTargets) AddTarget(label Label, target Target) error {
	if target == nil {
		panic("invaild target")
	}

	if _, alreadyPresent := self[label.Target]; alreadyPresent {
		return fmt.Errorf("target %v already exists", label)
	}

	self[label.Target] = target
	return nil
}

type WorkspaceTargets map[PackageName]PackageTargets

func (self WorkspaceTargets) AddTarget(label Label, target Target) error {
	packageTargets, ok := self[label.Package]
	if !ok {
		packageTargets = make(PackageTargets)
		self[label.Package] = packageTargets
	}

	return packageTargets.AddTarget(label, target)
}

type BuildTargets map[WorkspaceName]WorkspaceTargets

func (self BuildTargets) AddTarget(label Label, target Target) error {
	workspaceTargets, ok := self[label.Workspace]
	if !ok {
		workspaceTargets = make(WorkspaceTargets)
		self[label.Workspace] = workspaceTargets
	}

	return workspaceTargets.AddTarget(label, target)
}
