// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package cmake implements the conversion of Bazel BUILD files to CMake CMakeLists.txt files.
package cmake // import "src.tricot.io/public/bazel2x/converters/cmake"

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"src.tricot.io/public/bazel2x/bazel"
	"src.tricot.io/public/bazel2x/bazel/builtins/rules"
	"src.tricot.io/public/bazel2x/bazel/core"
)

func toDashes(s string) string {
	return strings.ReplaceAll(s, "/", "-")
}

func dashJoin(parts ...string) string {
	return strings.Join(parts, "-")
}

type CmakeConverter struct {
	// MinimumVersion is the minimum CMake version (e.g., "3.10.0"). If empty, "3.10.0" will be
	// used.
	MinimumVersion string `json:"minimumVersion"`

	// ProjectPrefix is the prefix to prepend (not including separating '-') to all project
	// names (it is also the name of the root project). If empty, the workspace name will be
	// used.
	ProjectPrefix string `json:"projectPrefix"`

	// StartWorkspaceName is the CMake name for the macro to call in the root CMakeLists.txt
	// before emitting any of our own targets (possibly by adding our own subdirectories). If
	// empty, "bazel2cmake_start_workspace" will be used.
	StartWorkspaceName string `json:"startWorkspaceName"`

	// EndWorkspaceName is the CMake name for the macro to call in the root CMakeLists.txt
	// after emitting all of our own targets (possibly by adding our own subdirectories). If
	// empty, "bazel2cmake_end_workspace" will be used.
	EndWorkspaceName string `json:"endWorkspaceName"`

	// CcLibraryName is the CMake name to use for cc_library targets. If empty,
	// "bazel2cmake_cc_library" will be used.
	CcLibraryName string `json:"ccLibraryName"`

	// CcBinaryName is the CMake name to use for cc_binary targets. If empty,
	// "bazel2cmake_cc_binary" will be used.
	CcBinaryName string `json:"ccBinaryName"`

	// CcTestName is the CMake name to use for cc_test targets. If empty, "bazel2cmake_cc_test"
	// will be used.
	CcTestName string `json:"ccTestName"`

	// Includes are the includes for the root CMakeLists.txt. If nil, "cmake/bazel2cmake.cmake"
	// will be included.
	Includes []string `json:"includes"`

	// RootUserHeader is the custom part of the header for the root CMakeLists.txt; it is a list
	// of lines. If nil, nothing will be added.
	RootUserHeader []string `json:"rootUserHeader"`

	// ExternalTargets are external targets that may appear as dependencies; it is a map from
	// label to CMake target name. This has precedence over ExternalWorkspaces.
	// TODO(vtl): Possibly the values should be a list of targets.
	ExternalTargets map[string]string `json:"externalTargets"`

	// ExternalWorkspaces are external workspaces also converted by bazel2cmake; it is a map
	// from workspace name (not including the leading '@') to project prefix (if empty, the
	// workspace name will be used). ExternalTargets has precedence over this.
	ExternalWorkspaces map[string]string `json:"externalWorkspaces"`

	// SkipTargets are targets to skip converting; its entries should be labels, like
	// "//foo/bar:baz". "//foo/bar:all" is also supported, in which case everything in the
	// "foo/bar" package is skipped. Warning: "//foo/bar" means "//foo/bar:bar". The root
	// package ("//:all") should not be skipped. TODO(vtl): Patterns involving ... aren't
	// supported yet.
	SkipTargets []string `json:"skipTargets"`

	build *bazel.Build

	skipPackagesSet map[string]struct{}
	skipTargetsSet  map[string]struct{}
}

func (self *CmakeConverter) Init(build *bazel.Build) error {
	self.build = build

	if self.MinimumVersion == "" {
		self.MinimumVersion = "3.10.0"
	}
	if self.ProjectPrefix == "" {
		if build.WorkspaceName != "" {
			self.ProjectPrefix = string(build.WorkspaceName)
		} else {
			self.ProjectPrefix = "bazel2cmake_project"
		}
	}
	if self.StartWorkspaceName == "" {
		self.StartWorkspaceName = "bazel2cmake_start_workspace"
	}
	if self.EndWorkspaceName == "" {
		self.EndWorkspaceName = "bazel2cmake_end_workspace"
	}
	if self.CcLibraryName == "" {
		self.CcLibraryName = "bazel2cmake_cc_library"
	}
	if self.CcBinaryName == "" {
		self.CcBinaryName = "bazel2cmake_cc_binary"
	}
	if self.CcTestName == "" {
		self.CcTestName = "bazel2cmake_cc_test"
	}
	if self.Includes == nil {
		self.Includes = []string{"cmake/bazel2cmake.cmake"}
	}

	self.skipPackagesSet = make(map[string]struct{})
	self.skipTargetsSet = make(map[string]struct{})
	for _, r := range self.SkipTargets {
		l, err := core.ParseLabel("", "", r)
		if err != nil {
			return err
		}
		if l.Target == "all" {
			self.skipPackagesSet[l.Workspace.String()+l.Package.String()] = struct{}{}
		} else {
			self.skipTargetsSet[l.String()] = struct{}{}
		}
	}

	return nil
}

func (self *CmakeConverter) targetName(l core.Label) (string, error) {
	if !l.IsExternal() {
		return dashJoin(self.ProjectPrefix, toDashes(string(l.Package)),
			toDashes(string(l.Target))), nil
	}
	if rv, ok := self.ExternalTargets[l.String()]; ok {
		return rv, nil
	}
	if projectPrefix, ok := self.ExternalWorkspaces[string(l.Workspace)]; ok {
		if projectPrefix == "" {
			projectPrefix = string(l.Workspace)
		}
		return dashJoin(projectPrefix, toDashes(string(l.Package)),
			toDashes(string(l.Target))), nil
	}
	return "", fmt.Errorf("no known CMake target for label %v", l)
}

func (self *CmakeConverter) writeHeader(packageName core.PackageName, w io.Writer) error {
	if _, err := fmt.Fprintf(w, "# Code generated by bazel2cmake. DO NOT EDIT.\n"); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "\ncmake_minimum_required(VERSION %v)\n",
		self.MinimumVersion); err != nil {
		return err
	}

	var projectName string
	if packageName == "" {
		projectName = self.ProjectPrefix
	} else {
		projectName = dashJoin(self.ProjectPrefix, toDashes(string(packageName)))
	}
	if _, err := fmt.Fprintf(w, "project(%v LANGUAGES CXX)\n", projectName); err != nil {
		return err
	}

	return nil
}

func (self *CmakeConverter) writeTarget(targetName core.TargetName, target core.Target,
	w io.Writer) error {

	name, err := self.targetName(target.Label())
	if err != nil {
		return err
	}

	switch target.(type) {
	case *rules.CcLibraryTarget:
		t := target.(*rules.CcLibraryTarget)
		if _, err := fmt.Fprintf(w, "\n%v(\n", self.CcLibraryName); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "    %v\n", name); err != nil {
			return err
		}
		if t.Srcs != nil {
			if _, err := fmt.Fprintf(w, "    SRCS\n"); err != nil {
				return err
			}
			for _, l := range *t.Srcs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}
		if t.Hdrs != nil {
			if _, err := fmt.Fprintf(w, "    HDRS\n"); err != nil {
				return err
			}
			for _, l := range *t.Hdrs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}
		if t.Deps != nil {
			if _, err := fmt.Fprintf(w, "    DEPS\n"); err != nil {
				return err
			}
			for _, l := range *t.Deps {
				depName, err := self.targetName(l)
				if err != nil {
					return err
				}
				if _, err := fmt.Fprintf(w, "        %v\n", depName); err != nil {
					return err
				}
			}
		}
		if _, err := fmt.Fprintf(w, ")\n"); err != nil {
			return err
		}
	case *rules.CcBinaryTarget:
		t := target.(*rules.CcBinaryTarget)
		if _, err := fmt.Fprintf(w, "\n%v(\n", self.CcBinaryName); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "    %v\n", name); err != nil {
			return err
		}
		if t.Srcs != nil {
			if _, err := fmt.Fprintf(w, "    SRCS\n"); err != nil {
				return err
			}
			for _, l := range *t.Srcs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}
		if t.Deps != nil {
			if _, err := fmt.Fprintf(w, "    DEPS\n"); err != nil {
				return err
			}
			for _, l := range *t.Deps {
				depName, err := self.targetName(l)
				if err != nil {
					return err
				}
				if _, err := fmt.Fprintf(w, "        %v\n", depName); err != nil {
					return err
				}
			}
		}
		if _, err := fmt.Fprintf(w, ")\n"); err != nil {
			return err
		}
	case *rules.CcTestTarget:
		t := target.(*rules.CcTestTarget)
		if _, err := fmt.Fprintf(w, "\n%v(\n", self.CcTestName); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "    %v\n", name); err != nil {
			return err
		}
		if t.Srcs != nil {
			if _, err := fmt.Fprintf(w, "    SRCS\n"); err != nil {
				return err
			}
			for _, l := range *t.Srcs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}
		if t.Deps != nil {
			if _, err := fmt.Fprintf(w, "    DEPS\n"); err != nil {
				return err
			}
			for _, l := range *t.Deps {
				depName, err := self.targetName(l)
				if err != nil {
					return err
				}
				if _, err := fmt.Fprintf(w, "        %v\n", depName); err != nil {
					return err
				}
			}
		}
		if _, err := fmt.Fprintf(w, ")\n"); err != nil {
			return err
		}
	}

	return nil
}

func (self *CmakeConverter) writeTargets(packageTargets *core.PackageTargets, w io.Writer) error {
	for _, target := range packageTargets.TargetList {
		if _, skip := self.skipTargetsSet[target.Label().String()]; skip {
			continue
		}

		if err := self.writeTarget(target.Label().Target, target, w); err != nil {
			return err
		}
	}

	return nil
}

func (self *CmakeConverter) writeAddSubdirs(w io.Writer) error {
	workspaceTargets := self.build.BuildTargets[core.MainWorkspaceName]
	pkgs := make([]string, 0, len(workspaceTargets)-1)
	for packageName := range workspaceTargets {
		if packageName != "" {
			pkgs = append(pkgs, string(packageName))
		}
	}
	sort.Strings(pkgs)

	if len(pkgs) > 0 {
		if _, err := fmt.Fprintf(w, "\n"); err != nil {
			return err
		}
		for _, pkg := range pkgs {
			// TODO(vtl): This is kind of hacky, but we used string(packageName) since
			// we didn't want the leading //. If we didn't want to add "skipped"
			// comments, we could have skipped adding them to pkgs instead.
			if _, skip := self.skipPackagesSet[
				core.MainWorkspaceName.String()+"//"+pkg]; skip {
				if _, err := fmt.Fprintf(w, "# %v skipped.\n", pkg); err != nil {
					return err
				}
				continue
			}
			if _, err := fmt.Fprintf(w, "add_subdirectory(%v)\n", pkg); err != nil {
				return err
			}
		}
	}

	return nil
}

func (self *CmakeConverter) writeRootCmakeLists(packageName core.PackageName,
	packageTargets *core.PackageTargets, packagePath string) error {

	outputPath := filepath.Join(packagePath, "CMakeLists.txt")
	w, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	if err := self.writeHeader(packageName, w); err != nil {
		return err
	}

	if len(self.RootUserHeader) > 0 {
		if _, err := fmt.Fprintf(w, "\n"); err != nil {
			return err
		}
		for _, line := range self.RootUserHeader {
			if _, err := fmt.Fprintf(w, "%v\n", line); err != nil {
				return err
			}
		}
	}

	if len(self.Includes) > 0 {
		if _, err := fmt.Fprintf(w, "\n"); err != nil {
			return err
		}
		for _, inc := range self.Includes {
			if _, err := fmt.Fprintf(w, "include(%v)\n", inc); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprintf(w, "\n%v()\n", self.StartWorkspaceName); err != nil {
		return err
	}

	if err := self.writeTargets(packageTargets, w); err != nil {
		return err
	}

	if err := self.writeAddSubdirs(w); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "\n%v()\n", self.EndWorkspaceName); err != nil {
		return err
	}

	return nil
}

func (self *CmakeConverter) writeNonRootCmakeLists(packageName core.PackageName,
	packageTargets *core.PackageTargets, packagePath string) error {

	outputPath := filepath.Join(packagePath, "CMakeLists.txt")
	w, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	if err := self.writeHeader(packageName, w); err != nil {
		return err
	}

	if err := self.writeTargets(packageTargets, w); err != nil {
		return err
	}

	return nil
}

func (self *CmakeConverter) writeCmakeLists(packageName core.PackageName,
	packageTargets *core.PackageTargets, packagePath string) error {

	if packageName == "" {
		return self.writeRootCmakeLists(packageName, packageTargets, packagePath)
	} else {
		return self.writeNonRootCmakeLists(packageName, packageTargets, packagePath)
	}
}

func (self *CmakeConverter) Convert(outputPath string) error {
	workspaceTargets, ok := self.build.BuildTargets[core.MainWorkspaceName]
	if !ok {
		return fmt.Errorf("no targets in the main workspace")
	}

	for packageName, packageTargets := range workspaceTargets {
		if _, skip := self.skipPackagesSet[
			core.MainWorkspaceName.String()+packageName.String()]; skip {
			continue
		}

		packagePath := filepath.Join(outputPath, string(packageName))
		if err := os.MkdirAll(packagePath, os.ModePerm); err != nil {
			return err
		}

		if err := self.writeCmakeLists(packageName, packageTargets,
			packagePath); err != nil {
			return err
		}
	}

	return nil
}
