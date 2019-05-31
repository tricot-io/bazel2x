// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package utils // import "src.tricot.io/public/bazel2x/bazel/utils"

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ErrNoWorkspace is the error returned when no WORKSPACE file can be found (at or above a given
// directory).
var ErrNoWorkspace = errors.New("no WORKSPACE file found")

// FindWorkspaceDir finds the first directory at or above startDir containing a WORKSPACE file. It
// returns the directory's absolute path and the relative path from that directory to startDir
// (which may be empty), or an error (which is ErrNoWorkspace if no WORKSPACE file could be found
// above startDir).
func FindWorkspaceDir(startDir string) (string, string, error) {
	const rootDir = string(os.PathSeparator)

	// Note: filepath.Abs calls filepath.Clean on its result, so there won't be any repeated
	// "/"s, any unnecessary trailing "/" (i.e., only a trailing "/" if it's the root
	// directory), or any "." or ".." components.
	startDir, err := filepath.Abs(startDir)
	if err != nil {
		return "", "", err
	}

	var components []string
	if startDir == rootDir {
		components = []string{}
	} else {
		components = strings.Split(startDir, string(os.PathSeparator))
	}

	i := len(components)
	for ; i >= 0; i-- {
		absDir := filepath.Join(rootDir, filepath.Join(components[:i]...))
		_, err := os.Lstat(filepath.Join(absDir, "WORKSPACE"))
		switch {
		case err == nil:
			relDir := filepath.Join(components[i:]...)
			return absDir, relDir, nil
		case os.IsNotExist(err):
			break
		default:
			return "", "", err
		}
	}
	return "", "", ErrNoWorkspace
}
