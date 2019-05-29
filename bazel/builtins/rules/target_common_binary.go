// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"bazel2cmake/bazel/core"
)

type TargetCommonBinary struct {
	Args           *[]string `bazel:"args"`
	OutputLicenses *[]string `bazel:"output_licenses"`
}

var _ ProcessArgsTarget = (*TargetCommonBinary)(nil)

func (self *TargetCommonBinary) Process(ctx core.Context) error {
	return nil
}
