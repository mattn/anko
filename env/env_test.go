package env

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	env := NewEnv()
	env.Define("a", "a")
	output := env.String()
	expected := `No parent
a = "a"
`
	if output != expected {
		t.Errorf("received: %v - expected: %v", output, expected)
	}

	env = env.NewEnv()
	env.Define("b", "b")
	output = env.String()
	expected = `Has parent
b = "b"
`
	if output != expected {
		t.Errorf("received: %v - expected: %v", output, expected)
	}

	env = NewEnv()
	env.Define("c", "c")
	env.DefineType("string", "a")
	output = env.String()
	expected = `No parent
c = "c"
string = string
`
	if output != expected {
		t.Errorf("received: %v - expected: %v", output, expected)
	}
}

func TestGetEnvFromPath(t *testing.T) {
	t.Parallel()

	env := NewEnv()
	a, err := env.NewModule("a")
	if err != nil {
		t.Fatal("NewModule error:", err)
	}
	var b *Env
	b, err = a.NewModule("b")
	if err != nil {
		t.Fatal("NewModule error:", err)
	}
	var c *Env
	c, err = b.NewModule("c")
	if err != nil {
		t.Fatal("NewModule error:", err)
	}
	err = c.Define("d", "d")
	if err != nil {
		t.Fatal("Define error:", err)
	}

	e, err := env.GetEnvFromPath(nil)
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}

	e, err = env.GetEnvFromPath([]string{})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}

	e, err = env.GetEnvFromPath([]string{"a", "c"})
	expected := "no namespace called: c"
	if err == nil || err.Error() != expected {
		t.Fatalf("GetEnvFromPath error - received: %v - expected: %v", err, expected)
	}
	if e != nil {
		t.Fatal("GetEnvFromPath e not nil")
	}

	// a.b.c

	e, err = env.GetEnvFromPath([]string{"a", "b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	var value interface{}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok := value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	e, err = a.GetEnvFromPath([]string{"a", "b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	e, err = b.GetEnvFromPath([]string{"a", "b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	e, err = c.GetEnvFromPath([]string{"a", "b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	// b.c

	e, err = env.GetEnvFromPath([]string{"b", "c"})
	expected = "no namespace called: b"
	if err == nil || err.Error() != expected {
		t.Fatalf("GetEnvFromPath error - received: %v - expected: %v", err, expected)
	}
	if e != nil {
		t.Fatal("GetEnvFromPath e not nil")
	}

	e, err = a.GetEnvFromPath([]string{"b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	e, err = b.GetEnvFromPath([]string{"b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	e, err = c.GetEnvFromPath([]string{"b", "c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	// c

	e, err = env.GetEnvFromPath([]string{"c"})
	expected = "no namespace called: c"
	if err == nil || err.Error() != expected {
		t.Fatalf("GetEnvFromPath error - received: %v - expected: %v", err, expected)
	}
	if e != nil {
		t.Fatal("GetEnvFromPath e not nil")
	}

	e, err = b.GetEnvFromPath([]string{"c"})
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}

	e, err = c.GetEnvFromPath(nil)
	if err != nil {
		t.Fatal("GetEnvFromPath error:", err)
	}
	if e == nil {
		t.Fatal("GetEnvFromPath e nil")
	}
	value, err = e.Get("d")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok = value.(string)
	if !ok {
		t.Fatal("value not string")
	}
	if v != "d" {
		t.Errorf("value - received: %v - expected: %v", v, "d")
	}
}

func TestCopy(t *testing.T) {
	t.Parallel()

	parent := NewEnv()
	parent.Define("a", "a")
	parent.DefineType("b", []bool{})
	child := parent.NewEnv()
	child.Define("c", "c")
	child.DefineType("d", []int64{})
	copy := child.Copy()

	if v, e := copy.Get("a"); e != nil || v != "a" {
		t.Errorf("copy missing value")
	}
	if v, e := copy.Type("b"); e != nil || v != reflect.TypeOf([]bool{}) {
		t.Errorf("copy missing type")
	}
	if v, e := copy.Get("c"); e != nil || v != "c" {
		t.Errorf("copy missing value")
	}
	if v, e := copy.Type("d"); e != nil || v != reflect.TypeOf([]int64{}) {
		t.Errorf("copy missing type")
	}

	// TODO: add more get type tests

	copy.Set("a", "i")
	if v, e := child.Get("a"); e != nil || v != "i" {
		t.Errorf("parent was not modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "i" {
		t.Errorf("copy did not get parent value")
	}

	copy.Set("c", "j")
	if v, e := child.Get("c"); e != nil || v != "c" {
		t.Errorf("child was not modified")
	}
	if v, e := copy.Get("c"); e != nil || v != "j" {
		t.Errorf("copy child was not modified")
	}

	child.Set("a", "x")
	if v, e := child.Get("a"); e != nil || v != "x" {
		t.Errorf("parent was not modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "x" {
		t.Errorf("copy did not get parent value")
	}

	child.Set("c", "z")
	if v, e := child.Get("c"); e != nil || v != "z" {
		t.Errorf("child was not modified")
	}
	if v, e := copy.Get("c"); e != nil || v != "j" {
		t.Errorf("copy child was modified")
	}

	parent.Set("a", "m")
	if v, e := child.Get("a"); e != nil || v != "m" {
		t.Errorf("parent was not modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "m" {
		t.Errorf("copy did not get parent value")
	}

	parent.Define("x", "n")
	if v, e := child.Get("x"); e != nil || v != "n" {
		t.Errorf("child did not get parent value")
	}
	if v, e := copy.Get("x"); e != nil || v != "n" {
		t.Errorf("copy did not get parent value")
	}
}

func TestDeepCopy(t *testing.T) {
	t.Parallel()

	parent := NewEnv()
	parent.Define("a", "a")
	env := parent.NewEnv()
	copy := env.DeepCopy()

	// TODO: add more/better tests, like above
	if v, e := copy.Get("a"); e != nil || v != "a" {
		t.Errorf("copy doesn't retain original values")
	}
	parent.Set("a", "b")
	if v, e := env.Get("a"); e != nil || v != "b" {
		t.Errorf("son was not modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "a" {
		t.Errorf("copy got the new value")
	}
	parent.Set("a", "c")
	if v, e := env.Get("a"); e != nil || v != "c" {
		t.Errorf("original was not modified")
	}
	if v, e := copy.Get("a"); e != nil || v != "a" {
		t.Errorf("copy was modified")
	}
	parent.Define("b", "b")
	if _, e := copy.Get("b"); e == nil {
		t.Errorf("copy parent was modified")
	}
}
