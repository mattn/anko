package vm_test

import (
	"testing"
	"os"

	"github.com/mattn/anko/vm"
)

func TestBasicOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []vm.Test{
		{Script: `a=nil; print(a)`, RunOutput: []interface{}{5, nil}, Stdout: "<nil>"},
		{Script: `a=true; print(a)`, RunOutput: []interface{}{4, nil}, Stdout: "true"},
		{Script: `a=""; print(a)`, RunOutput: []interface{}{0, nil}, Stdout: ""},
		{Script: `a=2 + 1; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "3"},
		{Script: `a=2 - 1; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "1"},
		{Script: `a=2 * 1; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "2"},
		{Script: `a=4 / 2; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "2"},
		{Script: `a=1; a++; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "2"},
		{Script: `a=1; a--; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "0"},
		{Script: `a=1; a+=1; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "2"},
		{Script: `a=1; a-=1; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "0"},
		{Script: `a=2; a*=4; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "8"},
		{Script: `a=2; a/=2; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "1"},
		{Script: `a=2 ** 3; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "8"},
		{Script: `a=1 & 3; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "1"},
		{Script: `a=1 | 2; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "3"},
		{Script: `a=2 << 3; print(a)`, RunOutput: []interface{}{2, nil}, Stdout: "16"},
		{Script: `a=8 >> 2; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "2"},
		{Script: `a=7 % 3; print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "1"},
		{Script: `a=2 - (-2); print(a)`, RunOutput: []interface{}{1, nil}, Stdout: "4"},
		{Script: `a=^2; print(a)`, RunOutput: []interface{}{2, nil}, Stdout: "-3"},
		{Script: `a="a" * 4; print(a)`, RunOutput: []interface{}{4, nil}, Stdout: "aaaa"},
		{Script: `a=!true; print(a)`, RunOutput: []interface{}{5, nil}, Stdout: "false"},
	}
	vm.RunTests(t, tests, &vm.ImportTestPrint)
}
