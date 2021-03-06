// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Command bazel2cmake converts Bazel BUILD files to CMake CMakeLists.txt files.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"src.tricot.io/public/bazel2x/bazel"
	"src.tricot.io/public/bazel2x/bazel/core"
	"src.tricot.io/public/bazel2x/bazel/utils"
	"src.tricot.io/public/bazel2x/converters/cmake"
)

var bazelOutputBaseFlag = flag.String("bazel_output_base", "", "Bazel output base directory")
var configFileFlag = flag.String("config_file", "", "configuration file (e.g., bazel2cmake.json)")

var onlyPrintTargetsFlag = flag.Bool("only_print_targets", false, "print targets and exit")
var outDirFlag = flag.String("out_dir", "", "(root) output directory")

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

	var outputBase string
	if *bazelOutputBaseFlag == "" {
		outputBase, err = utils.DefaultOutputBaseDir(workspaceDir)
		if err != nil {
			fmt.Printf("ERROR: failed to get Bazel outputBase directory: %v\n", err)
			os.Exit(1)
		}
	} else {
		outputBase = *bazelOutputBaseFlag
	}
	fmt.Printf("Bazel output base directory: %v\n", outputBase)

	var bazel2cmakeConfig []byte
	if *configFileFlag != "" {
		// If -config_file was used, then read the config from there.
		bazel2cmakeConfig, err = ioutil.ReadFile(*configFileFlag)
		if err != nil {
			fmt.Printf("ERROR: failed to read configuration file %v: %v\n",
				*configFileFlag, err)
			os.Exit(1)
		}
		fmt.Printf("Configuration file: %v\n", *configFileFlag)
	} else {
		// Otherwise, look in several locations.
		locations := []string{
			filepath.Join(workspaceDir, "bazel2cmake.json"),
			filepath.Join(workspaceDir, "cmake", "bazel2cmake.json"),
		}

		for _, configFile := range locations {
			bazel2cmakeConfig, err = ioutil.ReadFile(configFile)
			if err == nil {
				fmt.Printf("Configuration file: %v\n", configFile)
				break
			}
			if !os.IsNotExist(err) {
				fmt.Printf("ERROR: failed to read configuration file %v: %v\n",
					configFile, err)
				os.Exit(1)
			}
			bazel2cmakeConfig = nil
		}
		if bazel2cmakeConfig == nil {
			fmt.Printf("Configuration file: n/a (using default configuration)\n")
		}
	}

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

	build := bazel.NewBuild(bazel.GetSourceFileReader(workspaceDir, outputBase))

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

	converter := cmake.CmakeConverter{}
	if bazel2cmakeConfig != nil {
		if err := json.Unmarshal(bazel2cmakeConfig, &converter); err != nil {
			fmt.Printf("ERROR: error parsing configuration: %v\n", err)
			os.Exit(1)
		}
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
