package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestJson(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	var toByteSlice = func(s string) []byte { return []byte(s) }
	tests := []vm.Test{
		{Script: `json = import("encoding/json"); a = {"b": "b"}; c, err = json.Marshal(a); err`, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": []byte(`{"b":"b"}`)}},
		{Script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(a, &b); err`, Input: map[string]interface{}{"a": []byte(`{"b": "b"}`)}, Output: map[string]interface{}{"a": []byte(`{"b": "b"}`), "b": map[string]interface{}{"b": "b"}}},
		{Script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(toByteSlice(a), &b); err`, Input: map[string]interface{}{"a": `{"b": "b"}`, "toByteSlice": toByteSlice}, Output: map[string]interface{}{"a": `{"b": "b"}`, "b": map[string]interface{}{"b": "b"}}},
		{Script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(toByteSlice(a), &b); err`, Input: map[string]interface{}{"a": `[["1", "2"],["3", "4"]]`, "toByteSlice": toByteSlice}, Output: map[string]interface{}{"a": `[["1", "2"],["3", "4"]]`, "b": []interface{}{[]interface{}{"1", "2"}, []interface{}{"3", "4"}}}},
	}
	vm.RunTests(t, tests, &testPackagesEnvSetupFunc)
}
