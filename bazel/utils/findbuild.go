// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package utils

import (
	"os"
	"path/filepath"
)

// FindBuildFiles finds all BUILD[.bazel] files under workspaceDir, return a sorted slice of
// relative paths. It skips paths in ignorePaths. TODO(vtl): It doesn't follow/support symlinks.
func FindBuildFiles(workspaceDir string, ignorePaths []string) ([]string, error) {
	ignorePathsSet := map[string]struct{}{}
	for _, ignorePath := range ignorePaths {
		ignorePathsSet[filepath.Join(workspaceDir, ignorePath)] = struct{}{}
	}

	rv := []string{}
	err := filepath.Walk(workspaceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if _, shouldIgnore := ignorePathsSet[path]; shouldIgnore {
				return filepath.SkipDir
			}
			return nil
		}
		if info.Name() == "BUILD" || info.Name() == "BUILD.bazel" {
			rv = append(rv, path)
		}
		return nil

	})
	return rv, err
}
