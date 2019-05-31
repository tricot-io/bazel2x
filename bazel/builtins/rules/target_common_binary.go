// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules // import "src.tricot.io/public/bazel2x/bazel/builtins/rules"

import (
	"src.tricot.io/public/bazel2x/bazel/builtins/args"
	"src.tricot.io/public/bazel2x/bazel/core"
)

type TargetCommonBinary struct {
	Args           *[]string `bazel:"args"`
	OutputLicenses *[]string `bazel:"output_licenses"`
}

var _ args.ProcessArgsTarget = (*TargetCommonBinary)(nil)

func (self *TargetCommonBinary) DidProcessArgs(ctx core.Context) error {
	return nil
}
