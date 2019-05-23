// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"fmt"
	"path/filepath"

	"go.starlark.net/starlark"

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

func checkFileType(fileType core.FileType, label core.Label) error {
	switch fileType {
	case core.FileTypeBuild:
		if string(label.Target) != "BUILD" && string(label.Target) != "BUILD.bazel" {
			return fmt.Errorf("load not allowed: %s is not a BUILD[.bazel] file", label)
		}
	case core.FileTypeBzl:
		if filepath.Ext(string(label.Target)) != ".bzl" {
			return fmt.Errorf("load not allowed: %s is not a .bzl file", label)
		}
	case core.FileTypeWorkspace:
		if string(label.Target) != "WORKSPACE" {
			return fmt.Errorf("load not allowed: %s is not a WORKSPACE file", label)
		}
	default:
		panic(fileType)
	}
	return nil
}

func (self *Loader) Load(ctx *Context, moduleLabel core.Label) (starlark.StringDict, error) {
	err := checkFileType(ctx.FileType, moduleLabel)
	if err != nil {
		return nil, err
	}
	moduleLabelString := moduleLabel.String()

	e, ok := self.cache[moduleLabelString]
	if ok {
		if e == nil {
			return nil, fmt.Errorf("cycle in load graph (involving %s)", moduleLabel)
		}
		return e.globals, e.err
	}

	sourceData, err := self.sourceFileReader(moduleLabel)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %s", moduleLabel, err)
	}

	self.cache[moduleLabelString] = nil

	thread := ctx.CreateThread(moduleLabel, core.FileTypeBzl)
	globals, err := starlark.ExecFile(thread, moduleLabelString, sourceData,
		ctx.MakeInitialGlobals())
	self.cache[moduleLabelString] = &loadEntry{globals, err}
	return globals, err
}

func LoadLabel(thread *starlark.Thread, moduleLabel core.Label) (starlark.StringDict, error) {
	ctx := GetContextImpl(thread)
	return ctx.build.Loader.Load(ctx, moduleLabel)
}

func Load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	ctx := GetContextImpl(thread)
	moduleLabel, err := core.ParseLabel(ctx.Label.Workspace, ctx.Label.Package, module)
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
