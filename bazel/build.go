// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"bazel2cmake/bazel/core"
)

type Build struct {
	Loader  *Loader
	Targets core.BuildTargets
}

func (self *Build) AddBuildFile(buildFileLabel core.Label) error {
	thread := CreateThread(self, buildFileLabel, FileTypeBuild)
	_, err := LoadLabel(thread, buildFileLabel)
	return err
}

// TODO(vtl): More, including impls.
func NewBuild(sourceFileReader SourceFileReader) *Build {
	return &Build{
		Loader:  NewLoader(sourceFileReader),
		Targets: make(core.BuildTargets),
	}
}
