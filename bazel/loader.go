// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"fmt"
	"path/filepath"

	"go.starlark.net/starlark"
)

type loadEntry struct {
	globals starlark.StringDict
	err     error
}

type Loader struct {
	sourceFileReader SourceFileReader
	cache            map[string]*loadEntry
}

func checkFileType(fileType FileType, label Label) error {
	switch fileType {
	case FileTypeBuild:
		if string(label.Target) != "BUILD" && string(label.Target) != "BUILD.bazel" {
			return fmt.Errorf("load not allowed: %s is not a BUILD[.bazel] file", label)
		}
	case FileTypeBzl:
		if filepath.Ext(string(label.Target)) != ".bzl" {
			return fmt.Errorf("load not allowed: %s is not a .bzl file", label)
		}
	case FileTypeWorkspace:
		if string(label.Target) != "WORKSPACE" {
			return fmt.Errorf("load not allowed: %s is not a WORKSPACE file", label)
		}
	default:
		panic(fileType)
	}
	return nil
}

func (self *Loader) Load(ctx *Context, module string) (starlark.StringDict, error) {
	moduleLabel, err := ParseLabel(ctx.Label.Workspace, ctx.Label.Package, module)
	if err != nil {
		return nil, err
	}
	err = checkFileType(ctx.FileType, moduleLabel)
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

	thread := ctx.CreateThread(moduleLabel, FileTypeBzl)
	globals, err := starlark.ExecFile(thread, module, sourceData, ctx.MakeInitialGlobals())
	self.cache[moduleLabelString] = &loadEntry{globals, err}
	return globals, err
}

func Load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	ctx := GetContext(thread)
	return ctx.Loader.Load(ctx, module)
}

func NewLoader(sourceFileReader SourceFileReader) *Loader {
	return &Loader{
		sourceFileReader: sourceFileReader,
		cache:            make(map[string]*loadEntry),
	}
}
