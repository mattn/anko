package testlib

import (
	"testing"

	"github.com/mattn/anko/internal/corelib"
	"github.com/mattn/anko/parser"
)

// Test is utility struct to make tests easy.
type Test struct {
	Script         string
	ParseError     error
	ParseErrorFunc *func(*testing.T, error)
	EnvSetupFunc   *func(*testing.T, corelib.Env)
	Types          map[string]interface{}
	Input          map[string]interface{}
	RunError       error
	RunErrorFunc   *func(*testing.T, error)
	RunOutput      interface{}
	Output         map[string]interface{}
}

// Options is utility struct to pass options to the test.
type Options struct {
	EnvSetupFunc *func(*testing.T, corelib.Env)
}

// Run runs VM tests
func Run(t *testing.T, tests []Test, testingOptions *Options) {
	for _, test := range tests {
		run(t, test, testingOptions)
	}
}

func run(t *testing.T, test Test, testingOptions *Options) {
	// parser.EnableErrorVerbose()
	stmt, err := parser.ParseSrc(test.Script)
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
	// Note: Still want to run the code even after a parse error to see what happens

	env := corelib.NewEnv()
	if testingOptions != nil && testingOptions.EnvSetupFunc != nil {
		(*testingOptions.EnvSetupFunc)(t, env)
	}
	if test.EnvSetupFunc != nil {
		(*test.EnvSetupFunc)(t, env)
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
	value, err = env.Run(stmt)
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

	if !corelib.ValueEqual(value, test.RunOutput) {
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

		if !corelib.ValueEqual(value, outputValue) {
			t.Errorf("outputName %v - received: %#v - expected: %#v - script: %v", outputName, value, outputValue, test.Script)
			t.Errorf("received type: %T - expected: %T", value, outputValue)
			continue
		}
	}
}
