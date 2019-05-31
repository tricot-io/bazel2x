// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel // import "src.tricot.io/public/bazel2x/bazel"

import (
	"io/ioutil"
	"path/filepath"

	"src.tricot.io/public/bazel2x/bazel/core"
)

type SourceFileReader func(sourceFileLabel core.Label) ([]byte, error)

func GetSourceFileReader(workspaceDir string, outputBase string) SourceFileReader {
	externalDir := filepath.Join(outputBase, "external")
	return func(sourceFileLabel core.Label) ([]byte, error) {
		sourceFilePath := sourceFileLabel.SourcePath(workspaceDir, externalDir)
		return ioutil.ReadFile(sourceFilePath)
	}
}
