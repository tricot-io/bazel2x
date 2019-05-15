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

func TestParseLabel(t *testing.T) {
	const currWorkspace WorkspaceName = "default_workspace"
	const currPackage PackageName = "default/package"

	valids := []struct{
		in  string
		out Label
	}{
		{":foo", Label{currWorkspace, currPackage, "foo"}},
		{":foo.txt", Label{currWorkspace, currPackage, "foo.txt"}},
		{":.", Label{currWorkspace, currPackage, "."}},
		{":@-+=", Label{currWorkspace, currPackage, "@-+="}},
		{"//foo:bar", Label{currWorkspace, "foo", "bar"}},
		{"//foo:bar.txt", Label{currWorkspace, "foo", "bar.txt"}},
		{"//foo/bar:baz", Label{currWorkspace, "foo/bar", "baz"}},
		{"//foo/bar:baz.txt", Label{currWorkspace, "foo/bar", "baz.txt"}},
		{"//foo/bar:baz/quux", Label{currWorkspace, "foo/bar", "baz/quux"}},
		{"//foo/bar:baz/quux.txt", Label{currWorkspace, "foo/bar", "baz/quux.txt"}},
		{"//foo/bar:baz.d/quux.txt", Label{currWorkspace, "foo/bar", "baz.d/quux.txt"}},
		{"@//foo:bar", Label{currWorkspace, "foo", "bar"}},
		{"@//foo/bar:baz", Label{currWorkspace, "foo/bar", "baz"}},
		{"@//foo/bar:baz/quux", Label{currWorkspace, "foo/bar", "baz/quux"}},
		{"@my_workspace//foo:bar", Label{"my_workspace", "foo", "bar"}},
		{"@my_workspace//foo/bar:baz", Label{"my_workspace", "foo/bar", "baz"}},
		{"@my_workspace//foo/bar:baz/quux", Label{"my_workspace", "foo/bar", "baz/quux"}},
	}
	for _, valid := range valids {
		label, err := ParseLabel(currWorkspace, currPackage, valid.in)
		if err != nil {
			t.Error(valid.in, " should not have resulted in error: ", err)
		} else if label != valid.out {
			t.Error(valid.in, " should have resulted in ", valid.out,
				", but resulted in ", label)
		}

	}

	invalids := []string{"", ":", "//", "@", "foo:bar", "./foo:bar", "foo//bar", ":!",
		"foo//@bar:baz", "f@@//bar:baz", "@_foo//bar:baz", "@foo//b+r:baz", "@foo//bar:b!z"}
	for _, invalid := range invalids {
		label, err := ParseLabel(currWorkspace, currPackage, invalid)
		if err == nil {
			t.Error(invalid, " should have resulted in error, but resulted in ", label)
		}
	}
}
