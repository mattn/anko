package core

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
)

var testCoreEnvSetupFunc = func(t *testing.T, env corelib.Env) { Import(env.(*vm.Env)) }

func init() {
	corelib.NewEnv = func() corelib.Env {
		return vm.NewEnv()
	}
}

func TestKeys(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `a = {}; b = keys(a)`, RunOutput: []interface{}{}, Output: map[string]interface{}{"a": map[interface{}]interface{}{}}},
		{Script: `a = {"a": nil}; b = keys(a)`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"a": nil}}},
		{Script: `a = {"a": 1}; b = keys(a)`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": map[interface{}]interface{}{"a": int64(1)}}},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testCoreEnvSetupFunc})
}

func TestKindOf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": reflect.Value{}}, RunOutput: "struct", Output: map[string]interface{}{"a": reflect.Value{}}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": nil}, RunOutput: "nil", Output: map[string]interface{}{"a": nil}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": true}, RunOutput: "bool", Output: map[string]interface{}{"a": true}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: "int32", Output: map[string]interface{}{"a": int32(1)}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: "int64", Output: map[string]interface{}{"a": int64(1)}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": float32(1.1)}, RunOutput: "float32", Output: map[string]interface{}{"a": float32(1.1)}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": float64(1.1)}, RunOutput: "float64", Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": "a"}, RunOutput: "string", Output: map[string]interface{}{"a": "a"}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": 'a'}, RunOutput: "int32", Output: map[string]interface{}{"a": 'a'}},

		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}(nil)}, RunOutput: "nil", Output: map[string]interface{}{"a": interface{}(nil)}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}(true)}, RunOutput: "bool", Output: map[string]interface{}{"a": interface{}(true)}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}(int32(1))}, RunOutput: "int32", Output: map[string]interface{}{"a": interface{}(int32(1))}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}(int64(1))}, RunOutput: "int64", Output: map[string]interface{}{"a": interface{}(int64(1))}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}(float32(1))}, RunOutput: "float32", Output: map[string]interface{}{"a": interface{}(float32(1))}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}(float64(1))}, RunOutput: "float64", Output: map[string]interface{}{"a": interface{}(float64(1))}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": interface{}("a")}, RunOutput: "string", Output: map[string]interface{}{"a": interface{}("a")}},

		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": []interface{}{}}, RunOutput: "slice", Output: map[string]interface{}{"a": []interface{}{}}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": []interface{}{nil}}, RunOutput: "slice", Output: map[string]interface{}{"a": []interface{}{nil}}},

		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": map[string]interface{}{}}, RunOutput: "map", Output: map[string]interface{}{"a": map[string]interface{}{}}},
		{Script: `kindOf(a)`, Input: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}, RunOutput: "map", Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}}},

		{Script: `a = make(interface); kindOf(a)`, RunOutput: "nil", Output: map[string]interface{}{"a": interface{}(nil)}},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testCoreEnvSetupFunc})
}

func TestRange(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	tests := []testlib.Test{
		// 0 arguments
		{Script: `range()`, RunError: fmt.Errorf("range expected at least 1 argument, got 0")},
		// 1 arguments(step == 1, start == 0)
		{Script: `range(-1)`, RunOutput: []int64{}},
		{Script: `range(0)`, RunOutput: []int64{}},
		{Script: `range(1)`, RunOutput: []int64{0}},
		{Script: `range(2)`, RunOutput: []int64{0, 1}},
		{Script: `range(10)`, RunOutput: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		// 2 arguments(step == 1)
		{Script: `range(-5,-1)`, RunOutput: []int64{-5, -4, -3, -2}},
		{Script: `range(-1,1)`, RunOutput: []int64{-1, 0}},
		{Script: `range(1,5)`, RunOutput: []int64{1, 2, 3, 4}},
		// 3 arguments
		// step == 2
		{Script: `range(-5,-1,2)`, RunOutput: []int64{-5, -3}},
		{Script: `range(1,5,2)`, RunOutput: []int64{1, 3}},
		{Script: `range(-1,5,2)`, RunOutput: []int64{-1, 1, 3}},
		// step < 0 and from small to large
		{Script: `range(-5,-1,-1)`, RunOutput: []int64{}},
		{Script: `range(1,5,-1)`, RunOutput: []int64{}},
		{Script: `range(-1,5,-1)`, RunOutput: []int64{}},
		// step < 0 and from large to small
		{Script: `range(-1,-5,-1)`, RunOutput: []int64{-1, -2, -3, -4}},
		{Script: `range(5,1,-1)`, RunOutput: []int64{5, 4, 3, 2}},
		{Script: `range(5,-1,-1)`, RunOutput: []int64{5, 4, 3, 2, 1, 0}},
		// 4,5 arguments
		{Script: `range(1,5,1,1)`, RunError: fmt.Errorf("range expected at most 3 arguments, got 4")},
		{Script: `range(1,5,1,1,1)`, RunError: fmt.Errorf("range expected at most 3 arguments, got 5")},
		// more 0 test
		{Script: `range(0,1,2)`, RunOutput: []int64{0}},
		{Script: `range(1,0,2)`, RunOutput: []int64{}},
		{Script: `range(1,2,0)`, RunError: fmt.Errorf("range argument 3 must not be zero")},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testCoreEnvSetupFunc})
}

func TestLoad(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	notFoundRunErrorFunc := func(t *testing.T, err error) {
		if err == nil || !strings.HasPrefix(err.Error(), "open testdata/not-found.ank:") {
			t.Errorf("load not-found.ank failed - received: %v", err)
		}
	}
	tests := []testlib.Test{
		{Script: `load('testdata/test.ank'); X(1)`, RunOutput: int64(2)},
		{Script: `load('testdata/not-found.ank'); X(1)`, RunErrorFunc: &notFoundRunErrorFunc},
		{Script: `load('testdata/broken.ank'); X(1)`, RunError: fmt.Errorf("syntax error")},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testCoreEnvSetupFunc})
}

func TestAnk(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	var testEnvSetupFunc = func(t *testing.T, env corelib.Env) {
		e := env.(*vm.Env)
		Import(e)
		packages.DefineImport(e)
	}
	tests := []testlib.Test{
		{Script: `load('testdata/testing.ank'); load('testdata/let.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/toString.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/op.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/func.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/len.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/for.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/switch.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/if.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/toBytes.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/toRunes.ank')`},
		{Script: `load('testdata/testing.ank'); load('testdata/chan.ank')`},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testEnvSetupFunc})
}

func TestDefined(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "")
	tests := []testlib.Test{
		{Script: `var a = 1; defined("a")`, RunOutput: true},
		{Script: `defined("a")`, RunOutput: false},
		{Script: `func(){ var a = 1 }(); defined("a")`, RunOutput: false},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testCoreEnvSetupFunc})
}
