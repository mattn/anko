package core

import (
	"testing"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

type coreTest struct {
	script string
	output interface{}
	error  string
}

func runCoreTests(t *testing.T, tests []coreTest) {
	t.Helper()
	for _, test := range tests {
		e := env.NewEnv()
		Import(e)
		value, err := vm.Execute(e, nil, test.script)
		if test.error != "" {
			if err == nil || err.Error() != test.error {
				t.Errorf("script %q error - received: %v - expected: %v", test.script, err, test.error)
			}
			continue
		}
		if err != nil {
			t.Errorf("script %q unexpected error: %v", test.script, err)
			continue
		}
		if value != test.output {
			t.Errorf("script %q output - received: %#v - expected: %#v", test.script, value, test.output)
		}
	}
}

func TestTypeOf(t *testing.T) {
	tests := []coreTest{
		{script: `typeOf(nil)`, output: "nil"},
		{script: `typeOf(true)`, output: "bool"},
		{script: `typeOf(1)`, output: "int64"},
		{script: `typeOf(1.1)`, output: "float64"},
		{script: `typeOf("a")`, output: "string"},
		{script: `typeOf([1])`, output: "[]interface {}"},
	}
	runCoreTests(t, tests)
}

func TestKindOf(t *testing.T) {
	tests := []coreTest{
		{script: `kindOf(nil)`, output: "nil"},
		{script: `kindOf(true)`, output: "bool"},
		{script: `kindOf(1)`, output: "int64"},
		{script: `kindOf("a")`, output: "string"},
	}
	runCoreTests(t, tests)
}

func TestKeys(t *testing.T) {
	tests := []coreTest{
		{script: `keys({"a": 1})[0]`, output: "a"},
		{script: `len(keys({}))`, output: int64(0)},
		{script: `keys(5)`, error: "keys expected map type, got int64"},
		{script: `keys(nil)`, error: "keys expected map type, got invalid"},
	}
	runCoreTests(t, tests)
}
