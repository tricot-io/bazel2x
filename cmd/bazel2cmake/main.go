// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Command bazel2cmake converts Bazel BUILD files to CMake CMakeLists.txt files.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"bazel2cmake/bazel"
	"bazel2cmake/bazel/core"
	"bazel2cmake/bazel/utils"
	"bazel2cmake/converters/cmake"
)

var onlyPrintTargetsFlag = flag.Bool("only-print-targets", false, "print targets and exit")
var outDirFlag = flag.String("out-dir", "", "(root) output directory")

func printTargets(build *bazel.Build) {
	for workspaceName, workspaceTargets := range build.BuildTargets {
		fmt.Printf("Workspace @%v\n", string(workspaceName))
		// TODO(vtl): This order isn't deterministic.
		for packageName, packageTargets := range workspaceTargets {
			fmt.Printf("  Package %v\n", packageName)
			for _, target := range packageTargets.TargetList {
				fmt.Printf("    Target %v\n", target.Label().Target)
				fmt.Printf("      %v\n", target)
			}
		}
	}
}

func main() {
	flag.Parse()

	var workspaceDir string
	args := flag.Args()
	switch {
	case len(args) == 0:
		// The default is to find the workspace root at or above the working directory.
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("ERROR: failed to get working directory: %v\n", err)
			os.Exit(1)
		}
		workspaceDir, _, err = utils.FindWorkspaceDir(cwd)
		if err != nil {
			fmt.Printf("ERROR: failed to find workspace root: %v\n", err)
			os.Exit(1)
		}
	case len(args) == 1:
		workspaceDir = args[0]
	default:
		fmt.Printf("ERROR: usage: %v [workspace-dir]\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("Workspace root: %v\n", workspaceDir)

	bazelIgnore := utils.ReadBazelIgnore(workspaceDir)
	buildFiles, err := utils.FindBuildFiles(workspaceDir, bazelIgnore)
	if err != nil {
		fmt.Printf("ERROR: failed to find BUILD[.bazel] files: %v\n", err)
		os.Exit(1)
	}

	// The "project name" is the name of the directory. This is a bit odd, but Bazel seems to
	// prefer to put things in bazel-<project name>, instead of bazel-<workspace-name>.
	projectName := filepath.Base(workspaceDir)
	if projectName == string(filepath.Separator) {
		fmt.Printf("ERROR: unable to determine project name\n")
		os.Exit(1)
	}
	fmt.Printf("Project name: %v\n", projectName)

	buildFileLabels := make([]core.Label, len(buildFiles))
	for i, buildFile := range buildFiles {
		dir := filepath.Dir(buildFile)
		if dir == "." {
			dir = ""
		}

		buildFileLabels[i] = core.Label{
			Workspace: "",
			Package:   core.PackageName(dir),
			Target:    core.TargetName(filepath.Base(buildFile)),
		}
	}

	build := bazel.NewBuild(bazel.GetSourceFileReader(workspaceDir, projectName))

	err = build.ExecWorkspaceFile()
	if err != nil {
		fmt.Printf("ERROR: failed to execute WORKSPACE file: %v\n", err)
		os.Exit(1)
	}

	workspaceName := string(build.WorkspaceName)
	if workspaceName == "" {
		workspaceName = "<unset>"
	}
	fmt.Printf("Workspace name: %v\n", workspaceName)

	for _, buildFileLabel := range buildFileLabels {
		err := build.ExecBuildFile(buildFileLabel)
		if err != nil {
			fmt.Printf("ERROR: failed to execute BUILD[.bzel] file %v: %v\n",
				buildFileLabel, err)
			os.Exit(1)
		}
	}

	if *onlyPrintTargetsFlag {
		printTargets(build)
		os.Exit(0)
	}

	outDir := workspaceDir
	if *outDirFlag != "" {
		outDir = *outDirFlag
	}

	// TODO(vtl)
	//converter := &cmake.CmakeConverter{
	//	MinimumVersion:     "3.10.0",
	//	ProjectPrefix:      "",  // Use default.
	//	StartWorkspaceName: "",  // Use default.
	//	EndWorkspaceName:   "",  // Use default.
	//	CcLibraryName:      "",  // Use default.
	//	CcBinaryName:       "",  // Use default.
	//	CcTestName:         "",  // Use default.
	//	Includes:           nil, // Use default.
	//	ExternalTargets:    nil, // Use default.
	//	ExternalWorkspaces: nil, // Use default.
	//}
	converter := &cmake.CmakeConverter{
		MinimumVersion:     "3.10.0",
		ProjectPrefix:      "", // Use default.
		StartWorkspaceName: "tricot_start_workspace",
		EndWorkspaceName:   "tricot_end_workspace",
		CcLibraryName:      "tricot_cc_library",
		CcBinaryName:       "tricot_cc_binary",
		CcTestName:         "tricot_cc_test",
		Includes:           []string{"TricotCommon"},
		ExternalTargets: map[string]string{
			"@googletest//:gtest":                  "gtest",
			"@googletest//:gtest_main":             "gtest_main",
			"@mpark_variant//:mpark-variant":       "mpark_variant",
			"@optional_lite//:optional-lite":       "optional-lite",
			"@string_view_lite//:string_view-lite": "string_view-lite",
			// TODO(vtl): I just made these up.
			"@boost//:asio":  "boost-asio",
			"@boost//:beast": "boost-beast",
		},
		ExternalWorkspaces: map[string]string{
			"tricot_tid_public_output_cpp": "",
		},
	}

	err = converter.Init(build)
	if err != nil {
		fmt.Printf("ERROR: failed to initialize converter: %v\n", err)
		os.Exit(1)
	}

	err = converter.Convert(outDir)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
