// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// NOTE: This file is named "target_common_tst.go", since those ending in "_test.go" are test files.

package rules

import (
	"bazel2cmake/bazel/core"
)

type TargetCommonTest struct {
	Args       *[]string `bazel:"args"`
	Size       *string   `bazel:"size"`
	Timeout    *string   `bazel:"timeout"`
	Flaky      *bool     `bazel:"flaky"`
	Local      *bool     `bazel:"local"`
	ShardCount *int64    `bazel:"shard_count"`
}

var _ ProcessRuleArgsTargetStruct = (*TargetCommonTest)(nil)

func (self *TargetCommonTest) Process(ctx core.Context) error {
	return nil
}

