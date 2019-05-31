// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

// Package args contains various helpers for extracting/processing starlark arguments.
package args

import (
	"fmt"
	"reflect"

	"go.starlark.net/starlark"

	"src.tricot.io/public/bazel2x/bazel/core"
)

// ProcessArgsTarget is implemented by structs that are targets for ProcessArgs. ProcessArgs will
// fill in a target's fields based on its field tags, and once complete it will call DidProcessArgs.
type ProcessArgsTarget interface {
	DidProcessArgs(ctx core.Context) error
}

func toLabel(value starlark.Value, ctx core.Context) (core.Label, error) {
	s, ok := value.(starlark.String)
	if !ok {
		return core.Label{}, fmt.Errorf("label value is not a string")
	}

	// If it's a valid target name (e.g., a filename), then we accept it as such.
	// TODO(vtl): This means that we always accept filenames as labels, which is too lax.
	if core.TargetName(s).IsValid() {
		label := core.Label{
			Workspace: ctx.Label().Workspace,
			Package:   ctx.Label().Package,
			Target:    core.TargetName(s),
		}
		if !label.IsValid() {
			panic(label)
		}
		return label, nil
	}

	label, err := core.ParseLabel(ctx.Label().Workspace, ctx.Label().Package, string(s))
	if err != nil {
		return core.Label{}, err
	}
	return label, nil
}

func setArg(argName string, value starlark.Value, ctx core.Context, dest reflect.Value) error {
	if !dest.CanSet() {
		panic(dest)
	}

	switch dest.Interface().(type) {
	case *bool:
		// TODO(vtl): This means we accept anything for bool arguments; is this what we
		// want? (Bazel accepts at least True/False and 1/0; I'm not sure what else.)
		boolValue := bool(value.Truth())
		dest.Set(reflect.ValueOf(&boolValue))
	case *int64:
		n, ok := value.(starlark.Int)
		if !ok {
			return fmt.Errorf("argument %v invalid: value is not an integer", argName)
		}
		intValue, ok := n.Int64()
		if !ok {
			return fmt.Errorf("argument %v invalid: integer out of range", argName)
		}
		dest.Set(reflect.ValueOf(&intValue))
	case *string:
		s, ok := value.(starlark.String)
		if !ok {
			return fmt.Errorf("argument %v invalid: value is not a string", argName)
		}
		stringValue := string(s)
		dest.Set(reflect.ValueOf(&stringValue))
	case *core.Label:
		labelValue, err := toLabel(value, ctx)
		if err != nil {
			return fmt.Errorf("argument %v invalid: %v", argName, err)
		}
		dest.Set(reflect.ValueOf(&labelValue))
	case *[]string:
		l, ok := value.(*starlark.List)
		if !ok {
			return fmt.Errorf("argument %v invalid: value is not a list", argName)
		}
		listValue := make([]string, l.Len())
		for i := 0; i < l.Len(); i++ {
			s, ok := l.Index(i).(starlark.String)
			if !ok {
				return fmt.Errorf("argument %v invalid: invalid element: value is "+
					"not a string", argName)
			}
			listValue[i] = string(s)
		}
		dest.Set(reflect.ValueOf(&listValue))
	case *[]core.Label:
		l, ok := value.(*starlark.List)
		if !ok {
			return fmt.Errorf("argument %v invalid: value is not a list", argName)
		}
		listValue := make([]core.Label, l.Len())
		for i := 0; i < l.Len(); i++ {
			labelValue, err := toLabel(l.Index(i), ctx)
			if err != nil {
				return fmt.Errorf("argument %v invalid: invalid element: %v",
					argName, err)
			}
			listValue[i] = labelValue
		}
		dest.Set(reflect.ValueOf(&listValue))
	default:
		panic(dest)
	}
	return nil
}

var processArgsTargetType = reflect.TypeOf((*ProcessArgsTarget)(nil)).Elem()

func processRuleArgsHelper(kwargs map[string]starlark.Value, ctx core.Context,
	targetVp reflect.Value) error {

	if !targetVp.Type().Implements(processArgsTargetType) {
		panic(targetVp)
	}

	didProcessArgs := targetVp.MethodByName("DidProcessArgs")
	if !didProcessArgs.IsValid() {
		panic(targetVp)
	}

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
			vfa := vf.Addr()
			if vfa.Type().Implements(processArgsTargetType) {
				if err := processRuleArgsHelper(kwargs, ctx, vfa); err != nil {
					return err
				}
			}
		}
	}

	errorValue := didProcessArgs.Call([]reflect.Value{reflect.ValueOf(ctx)})[0]
	if errorValue.IsNil() {
		return nil
	}
	return errorValue.Interface().(error)
}

// ProcessArgs processes kwargs for the given target. Currently, it always requires that all
// arguments be kwargs.
func ProcessArgs(args starlark.Tuple, kwargs []starlark.Tuple, ctx core.Context,
	target ProcessArgsTarget) error {

	if len(args) > 0 {
		return fmt.Errorf("all arguments should be passed as keyword arguments")
	}

	kwargs2 := make(map[string]starlark.Value)
	for _, elem := range kwargs {
		kwargs2[string(elem[0].(starlark.String))] = elem[1]
	}

	targetVp := reflect.ValueOf(target)
	if targetVp.Kind() != reflect.Ptr {
		panic(targetVp)
	}

	err := processRuleArgsHelper(kwargs2, ctx, targetVp)
	if err != nil {
		return err
	}

	// This is a conditional posing as a loop.
	// TODO(vtl): Map iteration order isn't deterministic, so it's weird.
	for k := range kwargs2 {
		return fmt.Errorf("unknown keyword argument %v", k)
	}

	return nil
}
