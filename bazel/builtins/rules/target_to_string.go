// Copyright 2019 Tricot Inc.
// Use of this source code is governed by the license in the LICENSE file.

package rules // import "src.tricot.io/public/bazel2x/bazel/builtins/rules"

import (
	"fmt"
	"reflect"
	"strings"

	"src.tricot.io/public/bazel2x/bazel/builtins/args"
	"src.tricot.io/public/bazel2x/bazel/core"
)

func getAttr(attrName string, src reflect.Value) string {
	var attrValue string
	switch v := src.Interface().(type) {
	case bool:
		if v {
			attrValue = "True"
		} else {
			attrValue = "False"
		}
	case int64:
		attrValue = fmt.Sprintf("%v", v)
	case string:
		attrValue = fmt.Sprintf("%q", v)
	case core.Label:
		attrValue = fmt.Sprintf("%q", v.String())
	case []string:
		attrValue = "["
		for i := range v {
			if i > 0 {
				attrValue += ", "
			}
			attrValue += fmt.Sprintf("%q", v[i])
		}
		attrValue += "]"
	case []core.Label:
		attrValue = "["
		for i := range v {
			if i > 0 {
				attrValue += ", "
			}
			attrValue += fmt.Sprintf("%q", v[i].String())
		}
		attrValue += "]"
	default:
		panic(v)
	}
	return attrName + " = " + attrValue
}

func getAttrs(targetVp reflect.Value, attrs *[]string) {
	v := targetVp.Elem()
	if v.Kind() != reflect.Struct {
		panic(v)
	}
	typ := v.Type()

	for i := 0; i < typ.NumField(); i++ {
		typf := typ.Field(i)
		vf := v.Field(i)
		if argName, ok := typf.Tag.Lookup("bazel"); ok {
			if argName[len(argName)-1] == '!' {
				argName = argName[:len(argName)-1]
			}

			if !vf.IsNil() {
				*attrs = append(*attrs, getAttr(argName, vf.Elem()))
			}
		} else if vf.Kind() == reflect.Struct {
			getAttrs(vf.Addr(), attrs)
		}
	}
}

func targetToString(ruleName string, target args.ProcessArgsTarget) string {
	targetVp := reflect.ValueOf(target)
	if targetVp.Kind() != reflect.Ptr {
		panic(targetVp)
	}

	attrs := []string{}
	getAttrs(targetVp, &attrs)
	return ruleName + "(" + strings.Join(attrs, ", ") + ")"
}
