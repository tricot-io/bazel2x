// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel_test

import (
	"testing"

	. "bazel2cmake/bazel"
)

func TestWorkspaceName_IsValid(t *testing.T) {
	valids := []WorkspaceName{"", "a", "A", "abc", "ABC", "Abc", "aBC", "a1", "ab_123", "a_b_C",
		"a_", "a_1_B_2_c_3__", "a____"}
	for _, valid := range valids {
		if !valid.IsValid() {
			t.Error(valid, " should be valid")
		}
	}

	invalids := []WorkspaceName{"!", "a#", ".", "_", "1", "..", "/", "/a", "a/", "a/b", "a//b",
		"a.b", "a-b", "a/./b", "./a", "./", "../a", "a/../b", "1a", "1A", "1_", "_a"}
	for _, invalid := range invalids {
		if invalid.IsValid() {
			t.Error(invalid, " should be invalid")
		}
	}
}

func TestPackageName_IsValid(t *testing.T) {
	valids := []PackageName{"", "a", "A", "0", "-", ".", "_", "aBc-DeF/012.gHi_JkL",
		"Ab/C-d/E_f/1.2", "1/2/3/4/5/6/7/8/9", "X/Y/Z/x/y/z"}
	for _, valid := range valids {
		if !valid.IsValid() {
			t.Error(valid, " should be valid")
		}
	}

	invalids := []PackageName{"!", "a=", "@", "A+", "/", "/a", "a/", "a//b"}
	for _, invalid := range invalids {
		if invalid.IsValid() {
			t.Error(invalid, " should be invalid")
		}
	}
}

func TestTargetName_IsValid(t *testing.T) {
	valids := []TargetName{"a", "A", "0", "_", "+", "-", "=", ",", "@", "~", "a.", "a.b",
		"a..b", "a/B", "0/1/2/A-b/C.e/f..G/h...i/_.+-=,@~789xyzXYz", "."}
	for _, valid := range valids {
		if !valid.IsValid() {
			t.Error(valid, " should be valid")
		}
	}

	invalids := []TargetName{"!", "a#", "..", "/", "/a", "a/", "a//b", "a/./b", "./a", "./",
		"../a", "a/../b", ""}
	for _, invalid := range invalids {
		if invalid.IsValid() {
			t.Error(invalid, " should be invalid")
		}
	}
}

func TestLabel_IsValid(t *testing.T) {
	valids := []Label{
		{"", "", "root-target"},
		{"", "", "."},
		{"", "my_package", "target.1"},
		{"", "my/pkg/over/here", "target/is/over.there"},
		{"my_workspace", "", "root-target"},
		{"your_workspace", "my_package", "target.1"},
		{"everyones_workspace", "my/pkg/over/here", "target/is/over.there"},
	}
	for _, valid := range valids {
		if !valid.IsValid() {
			t.Error(valid, " should be valid")
		}
	}

	invalids := []Label{
		{},
		{"", "my_package", ""},
		{"", "my_package", ".."},
		{"", "/bad/package", "target"},
		{"/bad/workspace", "good/package", "target"},
	}
	for _, invalid := range invalids {
		if invalid.IsValid() {
			t.Error(invalid, " should be invalid")
		}
	}
}

// TODO(vtl): Test ParseLabel.
