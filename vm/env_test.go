package vm

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	env := NewEnv()
	env.Define("foo", reflect.ValueOf("bar"))
	v, ok := env.Get("foo")
	if !ok {
		t.Fatalf(`Can't Get valeu for "foo"`)
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

	env = env.New()
	v, ok := env.Get("foo")
	if !ok {
		t.Fatalf(`Can't Get valeu for "foo"`)
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

	env = env.New()
	env.Define("foo", reflect.ValueOf(true))
	v, ok := env.Get("foo")
	if !ok {
		t.Fatalf(`Can't Get valeu for "foo"`)
	}
	if v.Kind() != reflect.Bool {
		t.Fatalf(`Can't Get bool value for "foo"`)
	}
	if v.Bool() != true {
		t.Fatalf("Expected %v, but %v:", true, v.Bool())
	}

	v, ok = orig.Get("foo")
	if !ok {
		t.Fatalf(`Can't Get valeu for "foo"`)
	}
	if v.Kind() != reflect.String {
		t.Fatalf(`Can't Get string value for "foo"`)
	}
	if v.String() != "bar" {
		t.Fatalf("Expected %v, but %v:", "bar", v.String())
	}

}
