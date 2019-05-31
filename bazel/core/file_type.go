// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core // import "src.tricot.io/public/bazel2x/bazel/core"

// FileType indicates the type of Bazel input file.
type FileType int

const (
	// FileTypeBuild indicates a Bazel BUILD/BUILD.bazel file.
	FileTypeBuild FileType = iota

	// FileTypeBzl indicates a Bazel .bzl file.
	FileTypeBzl

	// FileTypeWorkspace indicates a Bazel WORKSPACE file.
	FileTypeWorkspace
)
