// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package bazel

import (
	"go.starlark.net/starlark"
)

// NoOpBuiltinsGlobalsImpl is a no-op implementation of BuiltinsGlobals.
type NoOpBuiltinsGlobalsImpl struct {}

func (self *NoOpBuiltinsGlobalsImpl) AnalysisTestTransition(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Aspect(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Bind(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) ConfigurationField(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Depset(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) ExistingRules(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Fail(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Provider(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) RegisterExecutionPlatforms(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) RegisterToolchains(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) RepositoryRule(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Rule(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Select(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGlobalsImpl) Workspace(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsGlobals = (*NoOpBuiltinsGlobalsImpl)(nil)

type NoOpBuiltinsFunctionsImpl struct {}

func (self *NoOpBuiltinsFunctionsImpl) Package(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsFunctionsImpl) PackageGroup(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsFunctionsImpl) ExportsFiles(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsFunctionsImpl) Glob(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsFunctions = (*NoOpBuiltinsFunctionsImpl)(nil)

type NoOpBuiltinsAndroidRulesImpl struct {}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidBinary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AarImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidLibrary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidInstrumentationTest(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidLocalTest(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidDevice(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidNdkRepository(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsAndroidRulesImpl) AndroidSdkRepository(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsAndroidRules = (*NoOpBuiltinsAndroidRulesImpl)(nil)

// TODO(vtl): Split this up.
type NoOpBuiltinsImpl struct{
	NoOpBuiltinsGlobalsImpl
	NoOpBuiltinsFunctionsImpl
	NoOpBuiltinsAndroidRulesImpl
}

func (self *NoOpBuiltinsImpl) CcBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) CcImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) CcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) CcProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) FdoPrefetchHints(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) FdoProfile(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) CcTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) CcToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) CcToolchainSuite(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaLiteProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaPackageConfiguration(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaPlugin(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaRuntime(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) JavaToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) AppleBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) AppleStaticLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) J2objcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ObjcImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ObjcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ObjcProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ProtoLangToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) PyBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) PyLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) PyTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) PyRuntime(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ShBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ShLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ShTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ActionListener(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ExtraAction(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Filegroup(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Genquery(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) TestSuite(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Alias(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ConfigSetting(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Genrule(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ConstraintSetting(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) ConstraintValue(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Platform(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) Toolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) LocalRepository(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) MavenJar(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) MavenServer(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) NewLocalRepository(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) XcodeConfig(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsImpl) XcodeVersion(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ Builtins = (*NoOpBuiltinsImpl)(nil)
