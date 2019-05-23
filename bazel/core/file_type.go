// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core

type FileType int

const (
	FileTypeBuild FileType = iota
	FileTypeBzl
	FileTypeWorkspace
)
