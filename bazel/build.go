// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"fmt"

	"go.starlark.net/starlark"

	"bazel2cmake/bazel/builtins"
	"bazel2cmake/bazel/core"
)

type Build struct {
	Loader       *Loader
	BuildTargets core.BuildTargets
}

func (self *Build) Exec(moduleLabel core.Label, fileType core.FileType) error {
	thread := CreateThread(self, moduleLabel, fileType)

	// TODO(vtl): We need to make Loader vs Build sane. (Maybe get rid of Loader.)
	sourceData, err := self.Loader.sourceFileReader(moduleLabel)
	if err != nil {
		return fmt.Errorf("failed to read %v: %v", moduleLabel, err)
	}

	ctx := core.GetContext(thread)
	_, err = starlark.ExecFile(thread, moduleLabel.String(), sourceData,
		builtins.InitialGlobals(ctx))
	return err
}

func (self *Build) AddBuildFile(buildFileLabel core.Label) error {
	return self.Exec(buildFileLabel, core.FileTypeBuild)
}

// TODO(vtl): More, including impls.
func NewBuild(sourceFileReader SourceFileReader) *Build {
	return &Build{
		Loader:       NewLoader(sourceFileReader),
		BuildTargets: make(core.BuildTargets),
	}
}
