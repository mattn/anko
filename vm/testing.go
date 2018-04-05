package vm

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/mattn/anko/parser"
)

type envSetupFunc *func(*testing.T, *Env)

// Test is a struct use for testing VM
type Test struct {
	Script         string
	ParseError     error
	ParseErrorFunc *func(*testing.T, error)
	EnvSetupFunc   envSetupFunc
	Types          map[string]interface{}
	Input          map[string]interface{}
	RunError       error
	RunErrorFunc   *func(*testing.T, error)
	RunOutput      interface{}
	Output         map[string]interface{}
	Stdout         string
}

// RunTests runs VM tests
func RunTests(t *testing.T, tests []Test, options ...envSetupFunc) {
	for _, test := range tests {
		RunTest(t, test, options...)
	}
}

// RunTest runs a VM test
func RunTest(t *testing.T, test Test, options ...envSetupFunc) {
	stmts, err := parser.ParseSrc(test.Script)
	if test.ParseErrorFunc != nil {
		(*test.ParseErrorFunc)(t, err)
	} else if err != nil && test.ParseError != nil {
		if err.Error() != test.ParseError.Error() {
			t.Errorf("ParseSrc error - received: %v - expected: %v - script: %v", err, test.ParseError, test.Script)
			return
		}
	} else if err != test.ParseError {
		t.Errorf("ParseSrc error - received: %v - expected: %v - script: %v", err, test.ParseError, test.Script)
		return
	}

	env := NewEnv()
	if test.EnvSetupFunc != nil {
		(*test.EnvSetupFunc)(t, env)
	}
	for _, envSetupFunc := range options {
		(*envSetupFunc)(t, env)
	}

	for typeName, typeValue := range test.Types {
		err = env.DefineType(typeName, typeValue)
		if err != nil {
			t.Errorf("DefineType error: %v - typeName: %v - script: %v", err, typeName, test.Script)
			return
		}
	}

	for inputName, inputValue := range test.Input {
		err = env.Define(inputName, inputValue)
		if err != nil {
			t.Errorf("Define error: %v - inputName: %v - script: %v", err, inputName, test.Script)
			return
		}
	}

	var value interface{}
	value, err = Run(stmts, env)
	if test.RunErrorFunc != nil {
		(*test.RunErrorFunc)(t, err)
	} else if err != nil && test.RunError != nil {
		if err.Error() != test.RunError.Error() {
			t.Errorf("Run error - received: %v - expected: %v - script: %v", err, test.RunError, test.Script)
			return
		}
	} else if err != test.RunError {
		t.Errorf("Run error - received: %v - expected: %v - script: %v", err, test.RunError, test.Script)
		return
	}

	if !ValueEqual(value, test.RunOutput) {
		t.Errorf("Run output - received: %#v - expected: %#v - script: %v", value, test.RunOutput, test.Script)
		t.Errorf("received type: %T - expected: %T", value, test.RunOutput)
		return
	}

	for outputName, outputValue := range test.Output {
		value, err = env.Get(outputName)
		if err != nil {
			t.Errorf("Get error: %v - outputName: %v - script: %v", err, outputName, test.Script)
			return
		}

		if !ValueEqual(value, outputValue) {
			t.Errorf("outputName %v - received: %#v - expected: %#v - script: %v", outputName, value, outputValue, test.Script)
			t.Errorf("received type: %T - expected: %T", value, outputValue)
			continue
		}
	}

	v, err := env.Get("TestStdout")
	if err == nil && v != test.Stdout {
		t.Errorf("Stdout error - received: %#v - expected: %#v - script: %v", v, test.Stdout, test.Script)
		return
	} else if err != nil && err.Error() != "undefined symbol 'TestStdout'" {
		t.Errorf("Get error: %v - TestStdout - script: %v", err, test.Script)
		return
	}

	env.Destroy()
}

var tee = func(env *Env, n int, b *bytes.Buffer, err error) (int, error) {
	if err != nil {
		return n, err
	}

	oldOut, err := env.Get("TestStdout")
	if err != nil {
		return n, err
	}

	err = env.Set("TestStdout", fmt.Sprintf("%s%s", oldOut, b.String()))
	if err != nil {
		return n, err
	}

	fmt.Print(b.String())

	return n, nil
}
var ImportTestPrint = func(t *testing.T, env *Env) {
	err := env.Define("TestStdout", "")
	if err != nil {
		t.Fatalf("Define error: %v\n", err)
	}

	err = env.Define("print", func(a ...interface{}) (n int, err error) {
		var b = new(bytes.Buffer)
		n, err = fmt.Fprint(b, a...)
		return tee(env, n, b, err)
	})
	if err != nil {
		t.Fatalf("Define error: %v\n", err)
	}

	err = env.Define("println", func(a ...interface{}) (n int, err error) {
		var b = new(bytes.Buffer)
		n, err = fmt.Fprintln(b, a...)
		return tee(env, n, b, err)
	})
	if err != nil {
		t.Fatalf("Define error: %v\n", err)
	}

	err = env.Define("printf", func(format string, a ...interface{}) (n int, err error) {
		var b = new(bytes.Buffer)
		n, err = fmt.Fprintf(b, format, a...)
		return tee(env, n, b, err)
	})
	if err != nil {
		t.Fatalf("Define error: %v\n", err)
	}
}

// ValueEqual checks the values and returns true if equal
// If passed function, does extra checks otherwise just doing reflect.DeepEqual
func ValueEqual(v1 interface{}, v2 interface{}) bool {
	v1RV := reflect.ValueOf(v1)
	switch v1RV.Kind() {
	case reflect.Func:
		// This is best effort to check if functions match, but it could be wrong
		v2RV := reflect.ValueOf(v2)
		if !v1RV.IsValid() || !v2RV.IsValid() {
			if v1RV.IsValid() != !v2RV.IsValid() {
				return false
			}
			return true
		} else if v1RV.Kind() != v2RV.Kind() {
			return false
		} else if v1RV.Type() != v2RV.Type() {
			return false
		} else if v1RV.Pointer() != v2RV.Pointer() {
			// From reflect: If v's Kind is Func, the returned pointer is an underlying code pointer, but not necessarily enough to identify a single function uniquely.
			return false
		}
		return true
	}
	return reflect.DeepEqual(v1, v2)
}
