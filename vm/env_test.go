package vm

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	env := NewEnv()
	env.Define("foo", reflect.ValueOf("bar"))
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
	env.Define("foo", reflect.ValueOf("bar"))

	env = env.NewEnv()
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

func TestDefineModify(t *testing.T) {
	env := NewEnv()
	env.Define("foo", reflect.ValueOf("bar"))
	orig := env

	env = env.NewEnv()
	env.Define("foo", reflect.ValueOf(true))
	v, err := env.Get("foo")
	if err != nil {
		t.Fatalf(`Can't Get value for "foo"`)
	}
	if v.Kind() != reflect.Bool {
		t.Fatalf(`Can't Get bool value for "foo"`)
	}
	if v.Bool() != true {
		t.Fatalf("Expected %v, but %v:", true, v.Bool())
	}

	v, err = orig.Get("foo")
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
