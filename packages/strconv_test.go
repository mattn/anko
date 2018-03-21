package packages

import (
	"fmt"
	"os"
	"testing"
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
	tests := []testStruct{
		{script: `strconv = import("strconv"); a = true; b = strconv.FormatBool(a)`, runOutput: "true", output: map[string]interface{}{"a": true, "b": "true"}},
		{script: `strconv = import("strconv"); a = 1.1; b = strconv.FormatFloat(a, toRune("f"), -1, 64)`, input: map[string]interface{}{"toRune": toRune}, runOutput: "1.1", output: map[string]interface{}{"a": float64(1.1), "b": "1.1"}},
		{script: `strconv = import("strconv"); a = 1; b = strconv.FormatInt(a, 10)`, runOutput: "1", output: map[string]interface{}{"a": int64(1), "b": "1"}},
		{script: `strconv = import("strconv"); b = strconv.FormatInt(a, 10)`, input: map[string]interface{}{"a": uint64(1)}, runOutput: "1", output: map[string]interface{}{"a": uint64(1), "b": "1"}},

		{script: `strconv = import("strconv"); a = "true"; b, err = strconv.ParseBool(a); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: "<nil>", output: map[string]interface{}{"a": "true", "b": true, "err": "<nil>"}},
		{script: `strconv = import("strconv"); a = "2"; b, err = strconv.ParseBool(a); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: `strconv.ParseBool: parsing "2": invalid syntax`, output: map[string]interface{}{"a": "2", "b": false, "err": `strconv.ParseBool: parsing "2": invalid syntax`}},
		{script: `strconv = import("strconv"); a = "1.1"; b, err = strconv.ParseFloat(a, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: "<nil>", output: map[string]interface{}{"a": "1.1", "b": float64(1.1), "err": "<nil>"}},
		{script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseFloat(a, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: `strconv.ParseFloat: parsing "a": invalid syntax`, output: map[string]interface{}{"a": "a", "b": float64(0), "err": `strconv.ParseFloat: parsing "a": invalid syntax`}},
		{script: `strconv = import("strconv"); a = "1"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: "<nil>", output: map[string]interface{}{"a": "1", "b": int64(1), "err": "<nil>"}},
		{script: `strconv = import("strconv"); a = "1.1"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: `strconv.ParseInt: parsing "1.1": invalid syntax`, output: map[string]interface{}{"a": "1.1", "b": int64(0), "err": `strconv.ParseInt: parsing "1.1": invalid syntax`}},
		{script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: `strconv.ParseInt: parsing "a": invalid syntax`, output: map[string]interface{}{"a": "a", "b": int64(0), "err": `strconv.ParseInt: parsing "a": invalid syntax`}},
		{script: `strconv = import("strconv"); a = "1"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: "<nil>", output: map[string]interface{}{"a": "1", "b": uint64(1), "err": "<nil>"}},
		{script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, input: map[string]interface{}{"toString": toString}, runOutput: `strconv.ParseUint: parsing "a": invalid syntax`, output: map[string]interface{}{"a": "a", "b": uint64(0), "err": `strconv.ParseUint: parsing "a": invalid syntax`}},
	}
	runTests(t, tests)
}
