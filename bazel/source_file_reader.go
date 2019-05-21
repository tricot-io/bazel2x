// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"io/ioutil"
	"path/filepath"
)

type SourceFileReader func(sourceFileLabel Label) ([]byte, error)

func GetSourceFileReader(workspaceDir string, projectName string) SourceFileReader {
	externalDir := filepath.Join(workspaceDir, "bazel-" + projectName, "external")
	return func(sourceFileLabel Label) ([]byte, error) {
		sourceFilePath := sourceFileLabel.SourcePath(workspaceDir, externalDir)
		return ioutil.ReadFile(sourceFilePath)
	}
}