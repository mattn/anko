package vm

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	env := NewEnv()
	env.Define("foo", "bar")

	v, err := env.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}
}

func TestDefine(t *testing.T) {
	env := NewEnv()
	env.Define("foo", "bar")
	sub := env.NewEnv()

	v, err := sub.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}
}

func TestDefineModify(t *testing.T) {
	env := NewEnv()
	env.Define("foo", "bar")
	sub := env.NewEnv()
	sub.Define("foo", true)

	v, err := sub.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.Bool {
		t.Fatalf(`Can't Get bool value for "foo"`)
	}
	if v.Bool() != true {
		t.Fatalf("Expected %v, but %v:", true, v.Bool())
	}

	v, err = env.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}
}

func TestDefineType(t *testing.T) {
	env := NewEnv()
	env.DefineType("int", int(0))
	sub := env.NewEnv()
	sub.DefineType("str", "")
	pkg := env.NewPackage("pkg")
	pkg.DefineType("Bool", true)

	for _, e := range []*Env{env, sub, pkg} {
		typ, err := e.Type("int")
		if err != nil {
			t.Fatalf(`Can't get Type for "int"`)
		}
		if typ.Kind() != reflect.Int {
			t.Fatalf(`Can't get int Type for "int"`)
		}

		typ, err = e.Type("str")
		if err != nil {
			t.Fatalf(`Can't get Type for "str"`)
		}
		if typ.Kind() != reflect.String {
			t.Fatalf(`Can't get string Type for "str"`)
		}

		typ, err = e.Type("pkg.Bool")
		if err != nil {
			t.Fatalf(`Can't get Type for "pkg.Bool"`)
		}
		if typ.Kind() != reflect.Bool {
			t.Fatalf(`Can't get bool Type for "pkg.Bool"`)
		}
	}
}
