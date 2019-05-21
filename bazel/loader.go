// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"fmt"

	"go.starlark.net/starlark"
)

type loadEntry struct {
	globals starlark.StringDict
	err     error
}

type Loader struct {
	SourceFileReader func(sourceFileLabel Label) ([]byte, error)

	cache map[string]*loadEntry
}

func (self *Loader) Load(ctx *Context, module string) (starlark.StringDict, error) {
	moduleLabel, err := ParseLabel(ctx.Label.Workspace, ctx.Label.Package, module)
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

	sourceData, err := self.SourceFileReader(moduleLabel)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %s", moduleLabel, err)
	}

	self.cache[moduleLabelString] = nil

	thread := ctx.CreateThread(moduleLabel, FileTypeBzl)
	globals, err := starlark.ExecFile(thread, module, sourceData, Builtins)
	self.cache[moduleLabelString] = &loadEntry{globals, err}
	return globals, err
}

func Load(thread *starlark.Thread, module string) (starlark.StringDict, error) {
	ctx := GetContext(thread)
	return ctx.Loader.Load(ctx, module)
}
