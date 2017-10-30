package vm

import (
	"reflect"
	"sync"
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

func TestEnvRaces1(t *testing.T) {
	// Test creating same variable in parallel

	env := NewEnv()

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	go func() {
		<-waitChan
		err := env.Define("a", "a")
		if err != nil {
			t.Fatalf("Unable to Define a, a")
		}
		_, err = env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, a")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("a", "b")
		if err != nil {
			t.Fatalf("Unable to Define a, b")
		}
		_, err = env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, b")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("a", "c")
		if err != nil {
			t.Fatalf("Unable to Define a, c")
		}
		_, err = env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, c")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("a", "d")
		if err != nil {
			t.Fatalf("Unable to Define a, d")
		}
		_, err = env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, d")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("a", "e")
		if err != nil {
			t.Fatalf("Unable to Define a, e")
		}
		_, err = env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, e")
		}
		waitGroup.Done()
	}()

	close(waitChan)
	waitGroup.Wait()

	_, err := env.Get("a")
	if err != nil {
		t.Fatalf("Unable to get value for a")
	}
}

func TestEnvRaces2(t *testing.T) {
	// Test creating different variables in parallel

	env := NewEnv()

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	go func() {
		<-waitChan
		err := env.Define("a", "a")
		if err != nil {
			t.Fatalf("Unable to Define a, a")
		}
		_, err = env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, a")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("b", "b")
		if err != nil {
			t.Fatalf("Unable to Define b, b")
		}
		_, err = env.Get("b")
		if err != nil {
			t.Fatalf("Unable to get value for b, b")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("c", "c")
		if err != nil {
			t.Fatalf("Unable to Define c, c")
		}
		_, err = env.Get("c")
		if err != nil {
			t.Fatalf("Unable to get value for c, c")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("d", "d")
		if err != nil {
			t.Fatalf("Unable to Define d, d")
		}
		_, err = env.Get("d")
		if err != nil {
			t.Fatalf("Unable to get value for d, d")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		err := env.Define("e", "e")
		if err != nil {
			t.Fatalf("Unable to Define e, e")
		}
		_, err = env.Get("e")
		if err != nil {
			t.Fatalf("Unable to get value for e, e")
		}
		waitGroup.Done()
	}()

	close(waitChan)
	waitGroup.Wait()

	_, err := env.Get("a")
	if err != nil {
		t.Fatalf("Unable to get value for a")
	}
	_, err = env.Get("b")
	if err != nil {
		t.Fatalf("Unable to get value for b")
	}
	_, err = env.Get("c")
	if err != nil {
		t.Fatalf("Unable to get value for c")
	}
	_, err = env.Get("d")
	if err != nil {
		t.Fatalf("Unable to get value for d")
	}
	_, err = env.Get("e")
	if err != nil {
		t.Fatalf("Unable to get value for e")
	}
}

func TestEnvRaces3(t *testing.T) {
	// Test reading different variables in parallel

	env := NewEnv()

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)

	err := env.Define("a", "a")
	if err != nil {
		t.Fatalf("Unable to Define a, a")
	}
	err = env.Define("b", "b")
	if err != nil {
		t.Fatalf("Unable to Define b, b")
	}
	err = env.Define("c", "c")
	if err != nil {
		t.Fatalf("Unable to Define c, c")
	}
	err = env.Define("d", "d")
	if err != nil {
		t.Fatalf("Unable to Define d, d")
	}
	err = env.Define("e", "e")
	if err != nil {
		t.Fatalf("Unable to Define e, e")
	}

	go func() {
		<-waitChan
		_, err := env.Get("a")
		if err != nil {
			t.Fatalf("Unable to get value for a, a")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		_, err := env.Get("b")
		if err != nil {
			t.Fatalf("Unable to get value for b, b")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		_, err := env.Get("c")
		if err != nil {
			t.Fatalf("Unable to get value for c, c")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		_, err := env.Get("d")
		if err != nil {
			t.Fatalf("Unable to get value for d, d")
		}
		waitGroup.Done()
	}()
	go func() {
		<-waitChan
		_, err := env.Get("e")
		if err != nil {
			t.Fatalf("Unable to get value for e, e")
		}
		waitGroup.Done()
	}()

	close(waitChan)
	waitGroup.Wait()

	_, err = env.Get("a")
	if err != nil {
		t.Fatalf("Unable to get value for a")
	}
	_, err = env.Get("b")
	if err != nil {
		t.Fatalf("Unable to get value for b")
	}
	_, err = env.Get("c")
	if err != nil {
		t.Fatalf("Unable to get value for c")
	}
	_, err = env.Get("d")
	if err != nil {
		t.Fatalf("Unable to get value for d")
	}
	_, err = env.Get("e")
	if err != nil {
		t.Fatalf("Unable to get value for e")
	}
}

func TestEnvRaces4(t *testing.T) {
	// Create env
	env := NewEnv()

	// Define some values
	env.Define("foo", "bar")
	env.Define("bar", "foo")
	env.Define("one", "two")
	env.Define("hello", "there")
	env.Define("hey", "ho")

	// Get some values in parallel
	go func(env *Env, t *testing.T) {
		_, err := env.Get("foo")
		if err != nil {
			t.Fatalf(`Can't Get value for "foo"`)
		}
	}(env, t)

	go func(env *Env, t *testing.T) {
		_, err := env.Get("bar")
		if err != nil {
			t.Fatalf(`Can't Get value for "bar"`)
		}
	}(env, t)

	go func(env *Env, t *testing.T) {
		_, err := env.Get("one")
		if err != nil {
			t.Fatalf(`Can't Get value for "one"`)
		}
	}(env, t)

	go func(env *Env, t *testing.T) {
		_, err := env.Get("hello")
		if err != nil {
			t.Fatalf(`Can't Get value for "hello"`)
		}
	}(env, t)

	go func(env *Env, t *testing.T) {
		_, err := env.Get("hey")
		if err != nil {
			t.Fatalf(`Can't Get value for "hey"`)
		}
	}(env, t)

	// Get subs
	go func(env *Env, t *testing.T) {
		sub := env.NewEnv()

		_, err := sub.Get("foo")
		if err != nil {
			t.Fatalf(`Can't Get value for "foo"`)
		}
	}(env, t)

	go func(env *Env, t *testing.T) {
		sub := env.NewEnv()

		_, err := sub.Get("one")
		if err != nil {
			t.Fatalf(`Can't Get value for "one"`)
		}
	}(env, t)

	go func(env *Env, t *testing.T) {
		sub := env.NewEnv()

		_, err := sub.Get("bar")
		if err != nil {
			t.Fatalf(`Can't Get value for "bar"`)
		}
	}(env, t)

	// Define some types
	go env.DefineType("int", int(0))
	go env.DefineType("str", "")

	// Define packages
	go func(env *Env, t *testing.T) {
		pkg := env.NewPackage("pkg")
		pkg.DefineType("Bool", true)
	}(env, t)

	go func(env *Env, t *testing.T) {
		pkg := env.NewPackage("pkg2")
		pkg.DefineType("Bool", true)
	}(env, t)

	// Get some types
	go env.Type("int")
	go env.Type("str")
	go env.Type("int")
	go env.Type("str")
	go env.Type("int")
	go env.Type("str")
}
