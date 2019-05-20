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

// TODO(vtl): FIXME -- should get caller label from calling thread instead?
func (self *Loader) Load(caller Label, module string) (starlark.StringDict, error) {
	moduleLabel, err := ParseLabel(caller.Workspace, caller.Package, module)
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

	load := func(_ *starlark.Thread, module2 string) (starlark.StringDict, error) {
		return self.Load(moduleLabel, module2)
	}
	// TODO(vtl): FIXME -- need to add thread local stuff to `thread` (via InitThread, or maybe
	// we should have a CreateThread).
	thread := &starlark.Thread{Name: "exec " + moduleLabelString, Load: load}
	// TODO(vtl): FIXME -- nil is wrong; need to add builtins.
	globals, err := starlark.ExecFile(thread, module, sourceData, nil)
	self.cache[moduleLabelString] = &loadEntry{globals, err}
	return globals, err
}
