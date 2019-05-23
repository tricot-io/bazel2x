// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	// "fmt"

	// "go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type ProcessRuleArgsTargetStruct interface {
	Process(ctx core.Context) error
}

func ProcessRuleArgs(ctx core.Context, target ProcessRuleArgsTargetStruct) error {
	return nil
}
