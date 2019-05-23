// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package core_test

import (
	"testing"

	. "bazel2cmake/bazel/core"
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

func TestWorkspaceName_String(t *testing.T) {
	testCases := []struct {
		in  WorkspaceName
		out string
	}{
		{"", ""},
		{"a", "@a"},
		{"A", "@A"},
		{"abc", "@abc"},
		{"ABC", "@ABC"},
		{"Abc", "@Abc"},
		{"aBC", "@aBC"},
		{"a1", "@a1"},
		{"ab_123", "@ab_123"},
		{"a_b_C", "@a_b_C"},
		{"a_", "@a_"},
		{"a_1_B_2_c_3__", "@a_1_B_2_c_3__"},
		{"a____", "@a____"},
	}
	for _, testCase := range testCases {
		if out := testCase.in.String(); out != testCase.out {
			t.Error(string(testCase.in), " should have resulted in ", testCase.out,
				", but resulted in ", out)
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

func TestPackageName_String(t *testing.T) {
	testCases := []struct {
		in  PackageName
		out string
	}{
		{"", "//"},
		{"a", "//a"},
		{"A", "//A"},
		{"0", "//0"},
		{"-", "//-"},
		{".", "//."},
		{"_", "//_"},
		{"aBc-DeF/012.gHi_JkL", "//aBc-DeF/012.gHi_JkL"},
		{"Ab/C-d/E_f/1.2", "//Ab/C-d/E_f/1.2"},
		{"1/2/3/4/5/6/7/8/9", "//1/2/3/4/5/6/7/8/9"},
		{"X/Y/Z/x/y/z", "//X/Y/Z/x/y/z"},
	}
	for _, testCase := range testCases {
		if out := testCase.in.String(); out != testCase.out {
			t.Error(string(testCase.in), " should have resulted in ", testCase.out,
				", but resulted in ", out)
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

func TestTargetName_String(t *testing.T) {
	testCases := []struct {
		in  TargetName
		out string
	}{
		{"a", ":a"},
		{"A", ":A"},
		{"0", ":0"},
		{"_", ":_"},
		{"+", ":+"},
		{"-", ":-"},
		{"=", ":="},
		{",", ":,"},
		{"@", ":@"},
		{"~", ":~"},
		{"a.", ":a."},
		{"a.b", ":a.b"},
		{"a..b", ":a..b"},
		{"a/B", ":a/B"},
		{"0/1/2/A-b/C.e/f..G/h...i/_.+-=,@~789xyzXYz",
			":0/1/2/A-b/C.e/f..G/h...i/_.+-=,@~789xyzXYz"},
		{".", ":."},
	}
	for _, testCase := range testCases {
		if out := testCase.in.String(); out != testCase.out {
			t.Error(string(testCase.in), " should have resulted in ", testCase.out,
				", but resulted in ", out)
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

func TestLabel_String(t *testing.T) {
	// Note: We use ParseLabel (below).
	testCases := []struct {
		in  string
		out string
	}{
		{"//:foo", "//:foo"},
		{"//:foo.txt", "//:foo.txt"},
		{"//foo", "//foo:foo"},
		{"//foo:bar", "//foo:bar"},
		{"//foo:bar.txt", "//foo:bar.txt"},
		{"//foo/bar", "//foo/bar:bar"},
		{"//foo/bar:baz", "//foo/bar:baz"},
		{"//foo/bar:baz.txt", "//foo/bar:baz.txt"},
		{"//foo/bar:baz/quux", "//foo/bar:baz/quux"},
		{"//foo/bar:baz/quux.txt", "//foo/bar:baz/quux.txt"},
		{"//foo/bar:baz.d/quux.txt", "//foo/bar:baz.d/quux.txt"},
		{"@//:foo", "//:foo"},
		{"@//foo", "//foo:foo"},
		{"@//foo:bar", "//foo:bar"},
		{"@//foo/bar", "//foo/bar:bar"},
		{"@//foo/bar:baz", "//foo/bar:baz"},
		{"@//foo/bar:baz/quux", "//foo/bar:baz/quux"},
		{"@my_workspace//:foo", "@my_workspace//:foo"},
		{"@my_workspace//foo", "@my_workspace//foo:foo"},
		{"@my_workspace//foo:bar", "@my_workspace//foo:bar"},
		{"@my_workspace//foo/bar", "@my_workspace//foo/bar:bar"},
		{"@my_workspace//foo/bar:baz", "@my_workspace//foo/bar:baz"},
		{"@my_workspace//foo/bar:baz/quux", "@my_workspace//foo/bar:baz/quux"},
	}
	for _, testCase := range testCases {
		label, _ := ParseLabel("", "should_not_appear", testCase.in)
		if out := label.String(); out != testCase.out {
			t.Error(testCase.in, " should have resulted in ", testCase.out,
				", but resulted in ", out)
		}
	}
}

func TestLabel_SourcePath(t *testing.T) {
	workspaceDir := "/foo"
	externalDir := workspaceDir + "bazel-foo/external"

	// Note: We use ParseLabel (below).
	testCases := []struct {
		in  string
		out string
	}{
		{"//:foo.txt", workspaceDir + "/foo.txt"},
		{"//foo:bar.txt", workspaceDir + "/foo/bar.txt"},
		{"//foo/bar:baz.txt", workspaceDir + "/foo/bar/baz.txt"},
		{"//foo/bar:baz/quux.txt", workspaceDir + "/foo/bar/baz/quux.txt"},
		{"//foo/bar:baz.d/quux.txt", workspaceDir + "/foo/bar/baz.d/quux.txt"},
		{"@my_workspace//:foo.txt", externalDir + "/my_workspace/foo.txt"},
		{"@my_workspace//foo:bar.txt", externalDir + "/my_workspace/foo/bar.txt"},
		{"@my_workspace//foo/bar:baz.txt", externalDir + "/my_workspace/foo/bar/baz.txt"},
		{"@my_workspace//foo/bar:baz/quux.txt",
			externalDir + "/my_workspace/foo/bar/baz/quux.txt"},
	}
	for _, testCase := range testCases {
		label, _ := ParseLabel("", "should_not_appear", testCase.in)
		if out := label.SourcePath(workspaceDir, externalDir); out != testCase.out {
			t.Error(testCase.in, " should have resulted in ", testCase.out,
				", but resulted in ", out)
		}
	}
}

func TestParseLabel(t *testing.T) {
	const currWorkspace WorkspaceName = "default_workspace"
	const currPackage PackageName = "default/package"

	valids := []struct {
		in  string
		out Label
	}{
		{":.", Label{currWorkspace, currPackage, "."}},
		{":@-+=", Label{currWorkspace, currPackage, "@-+="}},
		{":foo", Label{currWorkspace, currPackage, "foo"}},
		{":foo.txt", Label{currWorkspace, currPackage, "foo.txt"}},
		{"//:foo", Label{currWorkspace, "", "foo"}},
		{"//:foo.txt", Label{currWorkspace, "", "foo.txt"}},
		{"//foo", Label{currWorkspace, "foo", "foo"}},
		{"//foo:bar", Label{currWorkspace, "foo", "bar"}},
		{"//foo:bar.txt", Label{currWorkspace, "foo", "bar.txt"}},
		{"//foo/bar", Label{currWorkspace, "foo/bar", "bar"}},
		{"//foo/bar:baz", Label{currWorkspace, "foo/bar", "baz"}},
		{"//foo/bar:baz.txt", Label{currWorkspace, "foo/bar", "baz.txt"}},
		{"//foo/bar:baz/quux", Label{currWorkspace, "foo/bar", "baz/quux"}},
		{"//foo/bar:baz/quux.txt", Label{currWorkspace, "foo/bar", "baz/quux.txt"}},
		{"//foo/bar:baz.d/quux.txt", Label{currWorkspace, "foo/bar", "baz.d/quux.txt"}},
		{"@//:foo", Label{currWorkspace, "", "foo"}},
		{"@//foo", Label{currWorkspace, "foo", "foo"}},
		{"@//foo:bar", Label{currWorkspace, "foo", "bar"}},
		{"@//foo/bar", Label{currWorkspace, "foo/bar", "bar"}},
		{"@//foo/bar:baz", Label{currWorkspace, "foo/bar", "baz"}},
		{"@//foo/bar:baz/quux", Label{currWorkspace, "foo/bar", "baz/quux"}},
		{"@my_workspace//:foo", Label{"my_workspace", "", "foo"}},
		{"@my_workspace//foo", Label{"my_workspace", "foo", "foo"}},
		{"@my_workspace//foo:bar", Label{"my_workspace", "foo", "bar"}},
		{"@my_workspace//foo/bar", Label{"my_workspace", "foo/bar", "bar"}},
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
