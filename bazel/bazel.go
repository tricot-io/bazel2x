// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package bazel contains types used to represent Bazel concepts (and associated functions).
package bazel

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidLabel = errors.New("invalid label")

// WorkspaceName is the name of a Bazel workspace (not including leading "@"), or empty if it refers
// to the "current" workspace.
//
// A valid, non-empty workspace name must consist of only characters 'A'-'Z', 'a'-z', '0'-'9', and
// '_', and must begin with a letter.
type WorkspaceName string

var workspaceNameRegexp = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9_]*)?$`)

// IsValid returns whether the given WorkspaceName is valid.
func (w WorkspaceName) IsValid() bool {
	return workspaceNameRegexp.MatchString(string(w))
}

// PackageName is the name of a Bazel package (not including leading "//").
//
// A valid package name must consist of only characters 'A'-'Z', 'a'-'z', '0'-'9', '/', '-', '.',
// and '_'. It may neither start nor end with "/", nor may it contain "//". Note that it may be
// empty (indicating the package at the build root), though this is not advisable.
type PackageName string

var packageNameRegexp = regexp.MustCompile(`^([A-Za-z0-9\-._]+(/[A-Za-z0-9\-._]+)*)?$`)

// IsValid returns whether the given PackageName is valid.
func (p PackageName) IsValid() bool {
	return packageNameRegexp.MatchString(string(p))
}

// TargetName is the name of a Bazel target (not including leading ":").
//
// A valid target name must consist of only characters 'A'-'Z', 'a'-'z', '0'-'9', '_', '/', '.',
// '+', '-', '=', ',', '@', and '~'. It may neither start nor end with "/", nor may it contain "//".
// Its '/'-separated components may not include ".." or ".", except in the case that the label is
// just ".". A valid target name may not be empty.
type TargetName string

// Note that this regexp doesn't check for the rules involving ".." or "." components.
var targetNameRegexp = regexp.MustCompile(`^[A-Za-z0-9_.+\-=,@~]+(/[A-Za-z0-9_.+\-=,@~]+)*$`)

// IsValid returns whether the given TargetName is valid.
func (t TargetName) IsValid() bool {
	if !targetNameRegexp.MatchString(string(t)) {
		return false
	}
	if t == "." {
		return true
	}
	components := strings.Split(string(t), "/")
	for _, component := range components {
		if component == ".." || component == "." {
			return false
		}
	}
	return true
}

// Label is a Bazel label.
type Label struct {
	Workspace WorkspaceName
	Package   PackageName
	Target    TargetName
}

// IsValid returns whether the given Label is valid.
func (l Label) IsValid() bool {
	return l.Workspace.IsValid() && l.Package.IsValid() && l.Target.IsValid()
}

// ParseLabel parses a label. If the label is absolute and doesn't have a workspace, currWorkspace
// will be applied. If the label is "relative" (only contains a target name), then both
// currWorkspace and currPackage will be applied.
//
// Note that this requires that source files (which are also targets) be written with a leading ":"
// (or full specified including package name and possibly also workspace name).
func ParseLabel(currWorkspace WorkspaceName, currPackage PackageName, s string) (Label, error) {
	if !currWorkspace.IsValid() {
		panic(currWorkspace)
	}
	if !currPackage.IsValid() {
		panic(currPackage)
	}

	// Labels must always contain ":":
	colon := strings.Index(s, ":")
	if colon == -1 {
		return Label{}, ErrInvalidLabel
	}
	target := TargetName(s[colon+1:])

	// Relative label (i.e., only specifies target):
	if colon == 0 {
		if rv := (Label{currWorkspace, currPackage, target}); rv.IsValid() {
			return rv, nil
		}
		return Label{}, ErrInvalidLabel
	}

	// Otherwise, it's absolute and must contain "//":
	slashslash := strings.Index(s, "//")
	if slashslash == -1 || slashslash >= colon {
		return Label{}, ErrInvalidLabel
	}
	pkg := PackageName(s[slashslash+2 : colon])

	// Absolute label without workspace:
	if slashslash == 0 {
		if rv := (Label{currWorkspace, pkg, target}); rv.IsValid() {
			return rv, nil
		}
		return Label{}, ErrInvalidLabel
	}

	// Absolute label with workspace (note that s must be nonempty since it contains ":", not to
	// mention "//"):
	if s[0] != '@' {
		return Label{}, ErrInvalidLabel
	}
	workspace := WorkspaceName(s[1:slashslash])
	// Note that workspace is allowed to be empty here (you're allowed to write, e.g.,
	// "@//foo:bar").

	if rv := (Label{workspace, pkg, target}); rv.IsValid() {
		return rv, nil
	}
	return Label{}, ErrInvalidLabel
}
