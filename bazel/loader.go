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

type loadEntry struct {
	globals starlark.StringDict
	err     error
}

type Loader struct {
	sourceFileReader SourceFileReader
	cache            map[string]*loadEntry
}

func (self *Loader) Load(ctx *ContextImpl, moduleLabel core.Label) (starlark.StringDict, error) {
	// Only .bzl files can ever be loaded.
	if filepath.Ext(string(moduleLabel.Target)) != ".bzl" {
		return nil, fmt.Errorf("%v: load not allowed: %v is not a .bzl file", ctx.Label(),
			moduleLabel)
	}
	moduleLabelString := moduleLabel.String()

	e, ok := self.cache[moduleLabelString]
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

	self.cache[moduleLabelString] = nil

	thread := ctx.CreateThread(moduleLabel, core.FileTypeBzl)
	globals, err := starlark.ExecFile(thread, moduleLabelString, sourceData,
		builtins.InitialGlobals(ctx))
	self.cache[moduleLabelString] = &loadEntry{globals, err}
	return globals, err
}

func Load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	ctx := GetContextImpl(thread)
	moduleLabel, err := core.ParseLabel(ctx.Label().Workspace, ctx.Label().Package, module)
	if err != nil {
		return nil, err
	}
	return ctx.build.Loader.Load(ctx, moduleLabel)
}

func NewLoader(sourceFileReader SourceFileReader) *Loader {
	return &Loader{
		sourceFileReader: sourceFileReader,
		cache:            make(map[string]*loadEntry),
	}
}
