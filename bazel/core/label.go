// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// WorkspaceName is the name of a Bazel workspace (not including leading "@"), or empty if it refers
// to the "current" workspace.
//
// A valid, non-empty workspace name must consist of only characters 'A'-'Z', 'a'-z', '0'-'9', and
// '_', and must begin with a letter.
type WorkspaceName string

const MainWorkspaceName WorkspaceName = ""

var workspaceNameRegexp = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9_]*)?$`)

// IsValid returns whether the given WorkspaceName is valid.
func (w WorkspaceName) IsValid() bool {
	return workspaceNameRegexp.MatchString(string(w))
}

// IsExternal returns whether the given WorkspaceName is external (i.e., not the main workspace).
func (w WorkspaceName) IsExternal() bool {
	return w != MainWorkspaceName
}

// String formats a workspace name as a string ("" if empty, "@<workspace name>" otherwise).
func (w WorkspaceName) String() string {
	if !w.IsExternal() {
		return ""
	}
	return "@" + string(w)
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

// String formats a package name as a string (prepends "//").
func (p PackageName) String() string {
	return "//" + string(p)
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

// String formats a target name as a string (prepends ":").
func (t TargetName) String() string {
	return ":" + string(t)
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

// IsExternal returns whether the given Label is external (i.e., not in the main workspace).
func (l Label) IsExternal() bool {
	return l.Workspace.IsExternal()
}

// String formats a label as a string. Note that it does not abbreviate "//foo/bar:bar" as
// "//foo/bar".
func (l Label) String() string {
	return l.Workspace.String() + l.Package.String() + l.Target.String()
}

// SourcePath returns the source path for the given label, assuming that it in fact does refer to a
// source file. workspaceDir is the directory for the main workspace; externalDir is the directory
// containing external workspaces (so their directories are externalDir/<workspace name>).
func (l Label) SourcePath(workspaceDir string, externalDir string) string {
	relPath := filepath.Join(string(l.Package), string(l.Target))
	if l.Workspace == "" {
		return filepath.Join(workspaceDir, relPath)
	}
	return filepath.Join(externalDir, string(l.Workspace), relPath)
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

	// Find the ":", which delimits the target.
	colon := strings.Index(s, ":")
	// If there's no ":", then the target is implicitly the same as the last component of the
	// package.
	if colon == -1 {
		lastslash := strings.LastIndex(s, "/")
		if lastslash == -1 {
			return Label{}, fmt.Errorf("invalid label: %v", s)
		}
		colon = len(s)
		s += ":" + s[lastslash+1:]
	}
	target := TargetName(s[colon+1:])

	// Relative label (i.e., only specifies target):
	if colon == 0 {
		if rv := (Label{currWorkspace, currPackage, target}); rv.IsValid() {
			return rv, nil
		}
		return Label{}, fmt.Errorf("invalid label: %v", s)
	}

	// Otherwise, it's absolute and must contain "//":
	slashslash := strings.Index(s, "//")
	if slashslash == -1 || slashslash >= colon {
		return Label{}, fmt.Errorf("invalid label: %v", s)
	}
	pkg := PackageName(s[slashslash+2 : colon])

	// Absolute label without workspace:
	if slashslash == 0 {
		if rv := (Label{currWorkspace, pkg, target}); rv.IsValid() {
			return rv, nil
		}
		return Label{}, fmt.Errorf("invalid label: %v", s)
	}

	// Absolute label with workspace (note that s must be nonempty since it contains ":", not to
	// mention "//"):
	if s[0] != '@' {
		return Label{}, fmt.Errorf("invalid label: %v", s)
	}
	workspace := WorkspaceName(s[1:slashslash])
	// Note that workspace is allowed to be empty here (you're allowed to write, e.g.,
	// "@//foo:bar"), but then we apply the current workspace.
	// TODO(vtl): Is that right? Hopefully.
	if workspace == "" {
		workspace = currWorkspace
	}

	rv := (Label{workspace, pkg, target})
	if !rv.IsValid() {
		return Label{}, fmt.Errorf("invalid label: %v", s)
	}
	return rv, nil
}
