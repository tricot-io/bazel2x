// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Command bazel2cmake converts (TODO(vtl): ... or, for now, will convert) Bazel BUILD files to
// CMake CMakeLists.txt files.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"bazel2cmake/bazel"
	"bazel2cmake/bazel/builtins/rules"
	"bazel2cmake/bazel/core"
	"bazel2cmake/bazel/utils"
)

var outDirFlag = flag.String("out-dir", "", "(root) output directory")

func printTargets(build *bazel.Build) {
	for workspaceName, workspaceTargets := range build.BuildTargets {
		fmt.Printf("Workspace @%v\n", string(workspaceName))
		for packageName, packageTargets := range workspaceTargets {
			fmt.Printf("  Package %v\n", packageName)
			for _, target := range packageTargets.TargetList {
				fmt.Printf("    Target %v\n", target.Label().Target)
				fmt.Printf("      %v\n", target)
			}
		}
	}
}

const cmakeMinimumVersion = "3.10.0"
//FIXME
//var cmakeIncludes = []string{"Bazel2cmakeCommon"}
var cmakeIncludes = []string{"TricotCommon"}
// TODO(vtl): my_project should come from the workspace name (which would mean we'd have to exec the
// WORKSPACE).
//FIXME
//const cmakeProjectPrefix = "bazel2cmake-my_project"
const cmakeProjectPrefix = "tricot-cpp_public"

func toDashes(s string) string {
	return strings.ReplaceAll(s, "/", "-")
}

func dashJoin(parts ...string) string {
	return strings.Join(parts, "-")
}

func writeCMakeListsHeader(packageName core.PackageName, w io.Writer) error {
	if _, err := fmt.Fprintf(w, "# Code generated by bazel2cmake. DO NOT EDIT.\n"); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "\ncmake_minimum_required(VERSION %v)\n",
		cmakeMinimumVersion); err != nil {
		return err
	}

	for _, inc := range cmakeIncludes {
		if _, err := fmt.Fprintf(w, "\ninclude(%v)\n", inc); err != nil {
			return err
		}
	}

	projectName := dashJoin(cmakeProjectPrefix, toDashes(string(packageName)))
	if _, err := fmt.Fprintf(w, "\nproject(%v LANGUAGES CXX)\n", projectName); err != nil {
		return err
	}

	return nil
}

func cMakeTargetName(l core.Label) string {
	if !l.IsExternal() {
		return dashJoin(cmakeProjectPrefix, toDashes(string(l.Package)),
			toDashes(string(l.Target)))
	} else {
		return fmt.Sprintf("# TODO (external dep): %v", l)
	}
}

//FIXME
//const cmakeCcLibraryName = "bazel2cmake_cc_library"
//const cmakeCcTestName = "bazel2cmake_cc_test"
const cmakeCcLibraryName = "tricot_cc_library"
const cmakeCcTestName = "tricot_cc_test"

func writeCMakeListsBody(targetName core.TargetName, target core.Target, w io.Writer) error {
	switch target.(type) {
	case *rules.CcLibraryTarget:
		t := target.(*rules.CcLibraryTarget)

		if _, err := fmt.Fprintf(w, "\n%v(\n", cmakeCcLibraryName); err != nil {
			return err
		}

		if _, err := fmt.Fprintf(w, "    %v\n",
			cMakeTargetName(target.Label())); err != nil {
			return err
		}

		if t.Srcs != nil {
			if _, err := fmt.Fprintf(w, "    SRCS\n"); err != nil {
				return err
			}
			for _, l := range *t.Srcs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}

		if t.Hdrs != nil {
			if _, err := fmt.Fprintf(w, "    HDRS\n"); err != nil {
				return err
			}
			for _, l := range *t.Hdrs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}

		if t.Deps != nil {
			if _, err := fmt.Fprintf(w, "    DEPS\n"); err != nil {
				return err
			}
			for _, l := range *t.Deps {
				depName := cMakeTargetName(l)
				if _, err := fmt.Fprintf(w, "        %v\n", depName); err != nil {
					return err
				}
			}
		}

		if _, err := fmt.Fprintf(w, ")\n"); err != nil {
			return err
		}
	case *rules.CcBinaryTarget:
		if _, err := fmt.Fprintf(w, "\n# TODO: cc_binary %v\n",
			string(targetName)); err != nil {
			return err
		}
	case *rules.CcTestTarget:
		t := target.(*rules.CcTestTarget)

		if _, err := fmt.Fprintf(w, "\n%v(\n", cmakeCcTestName); err != nil {
			return err
		}

		if _, err := fmt.Fprintf(w, "    %v\n",
			cMakeTargetName(target.Label())); err != nil {
			return err
		}

		if t.Srcs != nil {
			if _, err := fmt.Fprintf(w, "    SRCS\n"); err != nil {
				return err
			}
			for _, l := range *t.Srcs {
				// Assume that it's just a simple filename, so just use the target
				// part of the label.
				if _, err := fmt.Fprintf(w, "        %v\n",
					string(l.Target)); err != nil {
					return err
				}
			}
		}

		if t.Deps != nil {
			if _, err := fmt.Fprintf(w, "    DEPS\n"); err != nil {
				return err
			}
			for _, l := range *t.Deps {
				depName := cMakeTargetName(l)
				if _, err := fmt.Fprintf(w, "        %v\n", depName); err != nil {
					return err
				}
			}
		}

		if _, err := fmt.Fprintf(w, ")\n"); err != nil {
			return err
		}
	}

	return nil
}

func makeCMakeLists(packageName core.PackageName, packageTargets *core.PackageTargets,
	packagePath string) error {

	outputPath := filepath.Join(packagePath, "CMakeLists.txt")
	f, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := writeCMakeListsHeader(packageName, f); err != nil {
		return err
	}

	for _, target := range packageTargets.TargetList {
		if err := writeCMakeListsBody(target.Label().Target, target, f); err != nil {
			return err
		}
	}

	// TODO(vtl): Trailer.

	return nil
}

func makeAllCMakeLists(build *bazel.Build, outputPath string) error {
	workspaceTargets, ok := build.BuildTargets[core.MainWorkspaceName]
	if !ok {
		return fmt.Errorf("no targets in the main workspace")
	}

	for packageName, packageTargets := range workspaceTargets {
		fmt.Printf("Package %v\n", packageName)

		packagePath := filepath.Join(outputPath, string(packageName))
		if err := os.MkdirAll(packagePath, os.ModePerm); err != nil {
			return err
		}

		if err := makeCMakeLists(packageName, packageTargets, packagePath); err != nil {
			return err
		}
	}

	return nil
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

		fmt.Printf("Input BUILD file: %v\n", buildFileLabels[i])
	}

	build := bazel.NewBuild(bazel.GetSourceFileReader(workspaceDir, projectName))
	for _, buildFileLabel := range buildFileLabels {
		err := build.ExecBuildFile(buildFileLabel)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}

	// printTargets(build)

	outDir := workspaceDir
	if *outDirFlag != "" {
		outDir = *outDirFlag
	}
	err = makeAllCMakeLists(build, outDir)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
