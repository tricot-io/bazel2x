// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules

import (
	"fmt"
	"reflect"

	"go.starlark.net/starlark"

	"bazel2cmake/bazel/core"
)

type ProcessRuleArgsTargetStruct interface {
	Process(ctx core.Context) error
}

func setArg(argName string, value starlark.Value, ctx core.Context, dest reflect.Value) error {
	if !dest.CanSet() {
		panic(dest)
	}

	switch dest.Interface().(type) {
	case bool:
		// TODO(vtl): This means we accept anything for bool arguments; is this what we
		// want? (Bazel accepts at least True/False and 1/0; I'm not sure what else.)
		dest.SetBool(bool(value.Truth()))
	case string:
		if s, ok := value.(starlark.String); ok {
			dest.SetString(string(s))
		} else {
			return fmt.Errorf("argument %s should be a string", argName)
		}
	case core.Label:
		if s, ok := value.(starlark.String); ok {
			// For single labels, we don't accept filenames, which may not be right.
			label, err := core.ParseLabel(ctx.Label().Workspace, ctx.Label().Package,
				string(s))
			if err != nil {
				return err
			}
//FIXME
_ = label
		} else {
			return fmt.Errorf("argument %s should be a string (label)", argName)
		}

	default:
		// TODO

	}
	// TODO(vtl)
//	fmt.Printf("Want to set %v to %v\n", argName, value)
	return nil
}

func processRuleArgs(kwargs map[string]starlark.Value, ctx core.Context,
	targetVp reflect.Value) error {

	v := targetVp.Elem()
	if v.Kind() != reflect.Struct {
		panic(v)
	}
	typ := v.Type()

	for i := 0; i < typ.NumField(); i++ {
		typf := typ.Field(i)
		vf := v.Field(i)
		if argName, ok := typf.Tag.Lookup("bazel"); ok {
			if !vf.CanSet() {
				panic(vf)
			}

			argRequired := false
			if argName[len(argName)-1] == '!' {
				argRequired = true
				argName = argName[:len(argName)-1]
			}

			if arg, ok := kwargs[argName]; ok {
				delete(kwargs, argName)
				setArg(argName, arg, ctx, vf)
			} else if argRequired {
				return fmt.Errorf("target argument %v required", argName)
			}
		} else if vf.Kind() == reflect.Struct {
			if err := processRuleArgs(kwargs, ctx, vf.Addr()); err != nil {
				return err
			}
		}
	}
	return nil

}

func ProcessRuleArgs(args starlark.Tuple, kwargs []starlark.Tuple, ctx core.Context,
	target ProcessRuleArgsTargetStruct) error {

	if len(args) > 0 {
		return fmt.Errorf("rule arguments should be passed as kwargs")
	}
	kwargs2 := make(map[string]starlark.Value)
	for _, elem := range kwargs {
		kwargs2[string(elem[0].(starlark.String))] = elem[1]
	}

	targetVp := reflect.ValueOf(target)
	if targetVp.Kind() != reflect.Ptr {
		panic(targetVp)
	}

	err := processRuleArgs(kwargs2, ctx, targetVp)
	if err != nil {
		return err
	}

	// This is a conditional posing as a loop.
	// TODO(vtl): Map iteration order isn't deterministic, so it's weird.
	for k, _ := range kwargs2 {
		return fmt.Errorf("unknown rule argument %s", k)
	}

	return target.Process(ctx)
}
