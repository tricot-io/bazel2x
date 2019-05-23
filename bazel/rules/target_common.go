// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"fmt"

	"bazel2cmake/bazel/core"
)

type TargetCommon struct {
	Name               string       `bazel:"name!"`
	Data               []core.Label `bazel:"data"`
	Visibility         []core.Label `bazel:"visibility"`
	Toolchains         []core.Label `bazel:"toolchains"`
	Deps               []core.Label `bazel:"deps"`
	Deprecation        string       `bazel:"deprecation"`
	Tags               []string     `bazel:"tags"`
	Testonly           bool         `bazel:"testonly"`
	Features           []string     `bazel:"features"`
	Licenses           []string     `bazel:"licenses"`
	CompatibleWith     []core.Label `bazel:"compatible_with"`
	Distribs           []string     `bazel"distribs"`
	ExecCompatibleWith []core.Label `bazel:"exec_compatible_with"`
	RestrictedTo       []core.Label `bazel:"restricted_to"`

	nameLabel core.Label
}

func (self *TargetCommon) Process(ctx core.Context) error {
	self.nameLabel = core.Label{
		Workspace: ctx.Label().Workspace,
		Package:   ctx.Label().Package,
		Target:    core.TargetName(self.Name),
	}
	if !self.nameLabel.IsValid() {
		return fmt.Errorf("invalid target name %v", self.Name)
	}
	// TODO(vtl)
	return nil
}

