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

// NoOpBuiltinsFunctionsImpl is a no-op implementation of BuiltinsFunctions.
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

// NoOpBuiltinsAndroidRulesImpl is a no-op implementation of BuiltinsAndroidRules.
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

// NoOpBuiltinsCcRulesImpl is a no-op implementation of BuiltinsCcRules.
type NoOpBuiltinsCcRulesImpl struct {}

func (self *NoOpBuiltinsCcRulesImpl) CcBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) CcImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) CcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) CcProtoLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) FdoPrefetchHints(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) FdoProfile(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) CcTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) CcToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsCcRulesImpl) CcToolchainSuite(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsCcRules = (*NoOpBuiltinsCcRulesImpl)(nil)

// NoOpBuiltinsJavaRulesImpl is a no-op implementation of BuiltinsJavaRules.
type NoOpBuiltinsJavaRulesImpl struct {}

func (self *NoOpBuiltinsJavaRulesImpl) JavaBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaLiteProtoLibrary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaProtoLibrary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaPackageConfiguration(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaPlugin(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaRuntime(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsJavaRulesImpl) JavaToolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsJavaRules = (*NoOpBuiltinsJavaRulesImpl)(nil)

// NoOpBuiltins is a no-op implementation of Builtins. Note: Its implementation being split up
// allows other implementations to pick and choose no-op implementations.
// TODO(vtl): Split this up.
type NoOpBuiltinsImpl struct{
	NoOpBuiltinsGlobalsImpl
	NoOpBuiltinsFunctionsImpl
	NoOpBuiltinsAndroidRulesImpl
	NoOpBuiltinsCcRulesImpl
	NoOpBuiltinsJavaRulesImpl
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
