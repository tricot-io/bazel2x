// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"fmt"
	"path/filepath"

	"go.starlark.net/starlark"

	"bazel2cmake/bazel/builtins"
	"bazel2cmake/bazel/core"
)

type loadCacheEntry struct {
	globals starlark.StringDict
	err     error
}

type Build struct {
	sourceFileReader SourceFileReader

	// loadCache caches the result of load statements. Its keys are labels (as strings).
	loadCache map[string]*loadCacheEntry

	// BuildTargets contains the output build targets.
	BuildTargets core.BuildTargets
}

// ExecWorkspaceFile executes the WORKSPACE file (which always has label //:WORKSPACE). It should be
// called exactly once, before ExecBuildFile is called.
func (self *Build) ExecWorkspaceFile() error {
	workspaceFileLabel := core.Label{
		Workspace: "",
		Package:   "",
		Target:    "WORKSPACE",
	}
	return self.exec(workspaceFileLabel, core.FileTypeWorkspace)
}

// ExecBuildFile executes the BUILD[.bazel] file specified by buildFileLabel. It should be called at
// most once for each BUILD[.bazel] file.
func (self *Build) ExecBuildFile(buildFileLabel core.Label) error {
	self.BuildTargets.AddPackage(buildFileLabel.Workspace, buildFileLabel.Package)
	return self.exec(buildFileLabel, core.FileTypeBuild)
}

// exec executes the file specified by moduleLabel, of the given file type (which should be
// core.FileTypeBuild or perhaps core.FileTypeWorkspace).
func (self *Build) exec(moduleLabel core.Label, fileType core.FileType) error {
	thread := createThread(self, moduleLabel, fileType)

	sourceData, err := self.sourceFileReader(moduleLabel)
	if err != nil {
		return fmt.Errorf("failed to execute %v: read failed: %v", moduleLabel, err)
	}

	ctx := core.GetContext(thread)
	_, err = starlark.ExecFile(thread, moduleLabel.String(), sourceData,
		builtins.InitialGlobals(ctx))
	return err
}

// load loads the .bzl file specified by moduleLabel and caches the result (subsequent loads of the
// same file will return the cached result).
//
// This is mainly used by the free load function, which is given to the starlark.Thread.
func (self *Build) load(ctx core.Context, moduleLabel core.Label) (starlark.StringDict, error) {
	// Only .bzl files can ever be loaded.
	if filepath.Ext(string(moduleLabel.Target)) != ".bzl" {
		return nil, fmt.Errorf("%v: load not allowed: %v is not a .bzl file", ctx.Label(),
			moduleLabel)
	}
	moduleLabelString := moduleLabel.String()

	e, ok := self.loadCache[moduleLabelString]
	if ok {
		if e == nil {
			return nil, fmt.Errorf("%v: load of %v failed: cycle in load graph",
				ctx.Label(), moduleLabel)
		}
		return e.globals, e.err
	}

	sourceData, err := self.sourceFileReader(moduleLabel)
	if err != nil {
		return nil, fmt.Errorf("%v: load of %v failed: read failed: %v", ctx.Label(),
			moduleLabel, err)
	}

	self.loadCache[moduleLabelString] = nil

	thread := createThread(self, moduleLabel, core.FileTypeBzl)
	globals, err := starlark.ExecFile(thread, moduleLabelString, sourceData,
		builtins.InitialGlobals(ctx))
	self.loadCache[moduleLabelString] = &loadCacheEntry{globals, err}
	return globals, err
}

func NewBuild(sourceFileReader SourceFileReader) *Build {
	return &Build{
		sourceFileReader: sourceFileReader,
		loadCache:        make(map[string]*loadCacheEntry),
		BuildTargets:     make(core.BuildTargets),
	}
}

// load loads the given (.bzl) file specified by module. This is meant to be given to the
// starlark.Thread.
func load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	ctx := GetContextImpl(thread)
	moduleLabel, err := core.ParseLabel(ctx.Label().Workspace, ctx.Label().Package, module)
	if err != nil {
		return nil, fmt.Errorf("%v: load of %v failed: invalid label: %v", ctx.Label(),
			module, err)
	}
	return ctx.build.load(ctx, moduleLabel)
}
