package packages

import (
	"fmt"
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestStrconv(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	var toRune = func(s string) rune {
		if len(s) == 0 {
			return 0
		}
		return []rune(s)[0]
	}
	var toString = func(v interface{}) string {
		if b, ok := v.([]byte); ok {
			return string(b)
		}
		return fmt.Sprint(v)
	}
	tests := []vm.Test{
		{Script: `strconv = import("strconv"); a = true; b = strconv.FormatBool(a)`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: "true", Output: map[string]interface{}{"a": true, "b": "true"}},
		{Script: `strconv = import("strconv"); a = 1.1; b = strconv.FormatFloat(a, toRune("f"), -1, 64)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toRune": toRune}, RunOutput: "1.1", Output: map[string]interface{}{"a": float64(1.1), "b": "1.1"}},
		{Script: `strconv = import("strconv"); a = 1; b = strconv.FormatInt(a, 10)`, EnvSetupFunc: &testPackagesEnvSetupFunc, RunOutput: "1", Output: map[string]interface{}{"a": int64(1), "b": "1"}},
		{Script: `strconv = import("strconv"); b = strconv.FormatInt(a, 10)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"a": uint64(1)}, RunOutput: "1", Output: map[string]interface{}{"a": uint64(1), "b": "1"}},

		{Script: `strconv = import("strconv"); a = "true"; b, err = strconv.ParseBool(a); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "true", "b": true, "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "2"; b, err = strconv.ParseBool(a); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseBool: parsing "2": invalid syntax`, Output: map[string]interface{}{"a": "2", "b": false, "err": `strconv.ParseBool: parsing "2": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1.1"; b, err = strconv.ParseFloat(a, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1.1", "b": float64(1.1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseFloat(a, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseFloat: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": float64(0), "err": `strconv.ParseFloat: parsing "a": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1", "b": int64(1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "1.1"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseInt: parsing "1.1": invalid syntax`, Output: map[string]interface{}{"a": "1.1", "b": int64(0), "err": `strconv.ParseInt: parsing "1.1": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseInt: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": int64(0), "err": `strconv.ParseInt: parsing "a": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1", "b": uint64(1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, EnvSetupFunc: &testPackagesEnvSetupFunc, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseUint: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": uint64(0), "err": `strconv.ParseUint: parsing "a": invalid syntax`}},
	}
	vm.RunTests(t, tests)
}
