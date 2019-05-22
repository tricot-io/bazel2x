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

// NoOpBuiltinsObjCRulesImpl is a no-op implementation of BuiltinsObjCRules.
type NoOpBuiltinsObjCRulesImpl struct {}

func (self *NoOpBuiltinsObjCRulesImpl) AppleBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsObjCRulesImpl) AppleStaticLibrary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsObjCRulesImpl) J2objcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsObjCRulesImpl) ObjcImport(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsObjCRulesImpl) ObjcLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsObjCRulesImpl) ObjcProtoLibrary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsObjCRules = (*NoOpBuiltinsObjCRulesImpl)(nil)

// NoOpBuiltinsProtoBufRulesImpl is a no-op implementation of BuiltinsProtoBufRules.
type NoOpBuiltinsProtoBufRulesImpl struct {}

func (self *NoOpBuiltinsProtoBufRulesImpl) ProtoLangToolchain(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsProtoBufRulesImpl) ProtoLibrary(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsProtoBufRules = (*NoOpBuiltinsProtoBufRulesImpl)(nil)

// NoOpBuiltinsPythonRulesImpl is a no-op implementation of BuiltinsPythonRules.
type NoOpBuiltinsPythonRulesImpl struct {}

func (self *NoOpBuiltinsPythonRulesImpl) PyBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsPythonRulesImpl) PyLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsPythonRulesImpl) PyTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsPythonRulesImpl) PyRuntime(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsPythonRules = (*NoOpBuiltinsPythonRulesImpl)(nil)

// NoOpBuiltinsShellRulesImpl is a no-op implementation of BuiltinsShellRules.
type NoOpBuiltinsShellRulesImpl struct {}

func (self *NoOpBuiltinsShellRulesImpl) ShBinary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsShellRulesImpl) ShLibrary(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsShellRulesImpl) ShTest(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsShellRules = (*NoOpBuiltinsShellRulesImpl)(nil)

// NoOpBuiltinsExtraActionsRulesImpl is a no-op implementation of BuiltinsExtraActionsRules.
type NoOpBuiltinsExtraActionsRulesImpl struct {}

func (self *NoOpBuiltinsExtraActionsRulesImpl) ActionListener(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsExtraActionsRulesImpl) ExtraAction(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsExtraActionsRules = (*NoOpBuiltinsExtraActionsRulesImpl)(nil)

// NoOpBuiltinsGeneralRulesImpl is a no-op implementation of BuiltinsGeneralRules.
type NoOpBuiltinsGeneralRulesImpl struct {}

func (self *NoOpBuiltinsGeneralRulesImpl) Filegroup(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGeneralRulesImpl) Genquery(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGeneralRulesImpl) TestSuite(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGeneralRulesImpl) Alias(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGeneralRulesImpl) ConfigSetting(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsGeneralRulesImpl) Genrule(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsGeneralRules = (*NoOpBuiltinsGeneralRulesImpl)(nil)

// NoOpBuiltinsPlatformRulesImpl is a no-op implementation of BuiltinsPlatformRules.
type NoOpBuiltinsPlatformRulesImpl struct {}

func (self *NoOpBuiltinsPlatformRulesImpl) ConstraintSetting(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsPlatformRulesImpl) ConstraintValue(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsPlatformRulesImpl) Platform(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsPlatformRulesImpl) Toolchain(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsPlatformRules = (*NoOpBuiltinsPlatformRulesImpl)(nil)

// NoOpBuiltinsWorkspaceRulesImpl is a no-op implementation of BuiltinsWorkspaceRules.
type NoOpBuiltinsWorkspaceRulesImpl struct {}

func (self *NoOpBuiltinsWorkspaceRulesImpl) LocalRepository(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsWorkspaceRulesImpl) MavenJar(args starlark.Tuple, kwargs []starlark.Tuple) (
	starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsWorkspaceRulesImpl) MavenServer(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsWorkspaceRulesImpl) NewLocalRepository(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsWorkspaceRulesImpl) XcodeConfig(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

func (self *NoOpBuiltinsWorkspaceRulesImpl) XcodeVersion(args starlark.Tuple,
	kwargs []starlark.Tuple) (starlark.Value, error) {
	return starlark.None, nil
}

var _ BuiltinsWorkspaceRules = (*NoOpBuiltinsWorkspaceRulesImpl)(nil)

// NoOpBuiltins is a no-op implementation of Builtins. Note: Its implementation is split up so other
// implementations to pick and choose no-op implementations.
// TODO(vtl): Split this up.
type NoOpBuiltinsImpl struct{
	NoOpBuiltinsGlobalsImpl
	NoOpBuiltinsFunctionsImpl
	NoOpBuiltinsAndroidRulesImpl
	NoOpBuiltinsCcRulesImpl
	NoOpBuiltinsJavaRulesImpl
	NoOpBuiltinsObjCRulesImpl
	NoOpBuiltinsProtoBufRulesImpl
	NoOpBuiltinsPythonRulesImpl
	NoOpBuiltinsShellRulesImpl
	NoOpBuiltinsExtraActionsRulesImpl
	NoOpBuiltinsGeneralRulesImpl
	NoOpBuiltinsPlatformRulesImpl
	NoOpBuiltinsWorkspaceRulesImpl
}

var _ Builtins = (*NoOpBuiltinsImpl)(nil)
