package env

import (
	"reflect"
	"testing"
)

func TestDefinePackage(t *testing.T) {
	e := NewEnv()

	if custom := e.Package("custom"); custom != nil {
		t.Fatal("unexpected \"custom\" package", custom)
	}

	// Define the method "hello" for the package "custom" using the global variable.
	Packages["custom"] = map[string]reflect.Value{
		"hello": reflect.ValueOf(func() string { return "hello" }),
	}

	if custom := e.Package("custom"); custom == nil {
		t.Fatal("expected \"custom\" package")
	} else if hello, ok := custom["hello"]; !ok {
		t.Fatal("expected method definition \"hello\" in package \"custom\"")
	} else if fn, ok := hello.Interface().(func() string); !ok {
		t.Fatal("expected method func() string {}")
	} else if res := fn(); res != "hello" {
		t.Fatalf("expected return value \"hello\" but got %v", res)
	}

	// Re-define the method for the package "custom" at the environment level.
	e.DefinePackage("custom", map[string]reflect.Value{
		"hello": reflect.ValueOf(func() string { return "hallo" }),
	})

	if custom := e.Package("custom"); custom == nil {
		t.Fatal("expected \"custom\" package")
	} else if hello, ok := custom["hello"]; !ok {
		t.Fatal("expected method definition \"hello\" in package \"custom\"")
	} else if fn, ok := hello.Interface().(func() string); !ok {
		t.Fatal("expected method func() string {}")
	} else if res := fn(); res != "hallo" {
		t.Fatalf("expected return value \"hallo\" but got %v", res)
	}
}

func TestDefinePackageTypes(t *testing.T) {
	e := NewEnv()

	if custom := e.PackageTypes("custom"); custom != nil {
		t.Fatal("unexpected \"custom\" package", custom)
	}

	type A struct {
		a string
	}

	// Define the type "A" for the package "custom" using the global variable.
	PackageTypes["custom"] = map[string]reflect.Type{
		"A": reflect.TypeOf(A{}),
	}

	if custom := e.PackageTypes("custom"); custom == nil {
		t.Fatal("expected \"custom\" package")
	} else if typeA, ok := custom["A"]; !ok {
		t.Fatal("expected type definition \"A\" in package \"custom\"")
	} else {
		a := reflect.New(typeA).Interface()
		if is, want := a, (&A{}); !reflect.DeepEqual(is, want) {
			t.Fatalf("%T != %T", is, want)
		}
	}

	type B struct {
		b int
	}

	// Re-define the type for the package "custom" at the environment level.
	e.DefinePackageTypes("custom", map[string]reflect.Type{
		"A": reflect.TypeOf(B{}),
	})

	if custom := e.PackageTypes("custom"); custom == nil {
		t.Fatal("expected \"custom\" package")
	} else if typeA, ok := custom["A"]; !ok {
		t.Fatal("expected type definition \"A\" in package \"custom\"")
	} else {
		a := reflect.New(typeA).Interface()
		if is, want := a, (&B{}); !reflect.DeepEqual(is, want) {
			t.Fatalf("%T != %T", is, want)
		}
	}
}
