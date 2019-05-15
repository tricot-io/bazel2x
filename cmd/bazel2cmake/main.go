// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Command bazel2cmake converts (TODO(vtl): ... or, for now, will convert) Bazel BUILD files to
// CMake CMakeLists.txt files.
package main

import (
	"fmt"
	"os"

	"bazel2cmake/utils"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	absdir, reldir, err := utils.FindWorkspaceDir(cwd)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("absdir=%q, reldir=%q\n", absdir, reldir)
}
