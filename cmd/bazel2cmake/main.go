// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Command bazel2cmake converts (TODO(vtl): ... or, for now, will convert) Bazel BUILD files to
// CMake CMakeLists.txt files.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"bazel2cmake/bazel"
	"bazel2cmake/bazel/core"
	"bazel2cmake/utils"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Print("ERROR: BUILD[.bazel] argument(s) required\n")
		os.Exit(1)
	}

	// Use the first arg to determine the workspace root.
	workspaceDir, _, err := utils.FindWorkspaceDir(filepath.Dir(args[0]))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Workspace root: %s\n", workspaceDir)

	projectName := filepath.Base(workspaceDir)
	if projectName == string(filepath.Separator) {
		fmt.Printf("ERROR: unable to determine project name\n")
		os.Exit(1)
	}
	fmt.Printf("Project name: %s\n", projectName)

	// Preprocess all the arguments to get relative paths to the workspace dir, since we'll
	// Chdir to the workspace dir.
	buildFileLabels := make([]core.Label, len(args))
	for i, arg := range args {
		wsDir, relDir, err := utils.FindWorkspaceDir(filepath.Dir(arg))
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(1)
		}

		if wsDir != workspaceDir {
			fmt.Printf("ERROR: %s not in same workspace as %s\n", arg, args[0])
			os.Exit(1)
		}

		buildFileLabels[i] = core.Label{
			Workspace: "",
			Package:   core.PackageName(relDir),
			Target:    core.TargetName(filepath.Base(arg)),
		}

		fmt.Printf("Input BUILD file: %s\n", buildFileLabels[i])
	}

	err = os.Chdir(workspaceDir)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	build := bazel.NewBuild(bazel.GetSourceFileReader(workspaceDir, projectName))
	for _, buildFileLabel := range buildFileLabels {
		err := build.AddBuildFile(buildFileLabel)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			os.Exit(1)
		}
	}
}
