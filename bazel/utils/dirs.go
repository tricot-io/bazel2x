// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package utils // import "src.tricot.io/public/bazel2x/bazel/utils"

import (
	"crypto/md5"
	"fmt"
	"os/user"
	"path/filepath"
)

// DefaultOutputBaseDir returns the default outputBase directory for the given workspace directory
// (and current user).
func DefaultOutputBaseDir(workspaceDir string) (string, error) {
	absWorkspaceDir, err := filepath.Abs(workspaceDir)
	if err != nil {
		return "", err
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	// The default outputBase is $HOME/.cache/bazel/_bazel_$USER/<MD5 hash of absWorkspaceDir>.
	return fmt.Sprintf("%v/.cache/bazel/_bazel_%v/%x", usr.HomeDir, usr.Username,
		md5.Sum([]byte(absWorkspaceDir))), nil
}
