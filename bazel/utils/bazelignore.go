// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package utils // import "src.tricot.io/public/bazel2x/bazel/utils"

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ReadBazelIgnore(workspaceDir string) []string {
	contents, err := ioutil.ReadFile(filepath.Join(workspaceDir, ".bazelignore"))
	if err != nil {
		return []string{}
	}

	result := strings.FieldsFunc(string(contents), func(c rune) bool {
		return c == '\n' || c == '\r'
	})

	// Be nice and clean the results (e.g., in case there are trailing slashes).
	for i := range result {
		result[i] = filepath.Clean(result[i])
	}

	return result
}
