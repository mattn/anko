package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

var testPackagesEnvSetupFunc = func(t *testing.T, env *vm.Env) { DefineImport(env) }

func TestDefineImport(t *testing.T) {
	env := vm.NewEnv()
	DefineImport(env)
	value, err := env.Execute(`strings = import("strings"); strings.ToLower("TEST")`)
	if err != nil {
		t.Errorf("execute error - received: %v - expected: %v", err, nil)
	}
	if value != "test" {
		t.Errorf("execute value - received: %v - expected: %v", value, int(4))
	}
}

func TestDefineImportPackageNotFound(t *testing.T) {
	os.Unsetenv("ANKO_DEBUG")
	env := vm.NewEnv()
	DefineImport(env)
	value, err := env.Execute(`a = import("a")`)
	expectedError := "package 'a' not found"
	if err == nil || err.Error() != expectedError {
		t.Errorf("execute error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("execute value - received: %v - expected: %v", value, nil)
	}
}

func TestDefineImportPackageDefineError(t *testing.T) {
	os.Unsetenv("ANKO_DEBUG")
	Packages["testPackage"] = map[string]interface{}{"bad.name": testing.Coverage}
	env := vm.NewEnv()
	DefineImport(env)

	value, err := env.Execute(`a = import("testPackage")`)
	expectedError := "import Define error: unknown symbol 'bad.name'"
	if err == nil || err.Error() != expectedError {
		t.Errorf("execute error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("execute value - received: %v - expected: %v", value, nil)
	}

	Packages["testPackage"] = map[string]interface{}{"Coverage": testing.Coverage}
	PackageTypes["testPackage"] = map[string]interface{}{"bad.name": int64(1)}

	value, err = env.Execute(`a = import("testPackage")`)
	expectedError = "import DefineType error: unknown symbol 'bad.name'"
	if err == nil || err.Error() != expectedError {
		t.Errorf("execute error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("execute value - received: %v - expected: %v", value, nil)
	}
}
