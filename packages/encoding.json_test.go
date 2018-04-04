package packages

import (
	"os"
	"testing"
)

func TestJson(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	var toByteSlice = func(s string) []byte { return []byte(s) }
	tests := []testStruct{
		{script: `json = import("encoding/json"); a = {"b": "b"}; c, err = json.Marshal(a); err`, output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": []byte(`{"b":"b"}`)}},
		{script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(a, &b); err`, input: map[string]interface{}{"a": []byte(`{"b": "b"}`)}, output: map[string]interface{}{"a": []byte(`{"b": "b"}`), "b": map[string]interface{}{"b": "b"}}},
		{script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(toByteSlice(a), &b); err`, input: map[string]interface{}{"a": `{"b": "b"}`, "toByteSlice": toByteSlice}, output: map[string]interface{}{"a": `{"b": "b"}`, "b": map[string]interface{}{"b": "b"}}},
		{script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(toByteSlice(a), &b); err`, input: map[string]interface{}{"a": `[["1", "2"],["3", "4"]]`, "toByteSlice": toByteSlice}, output: map[string]interface{}{"a": `[["1", "2"],["3", "4"]]`, "b": []interface{}{[]interface{}{"1", "2"}, []interface{}{"3", "4"}}}},
	}
	runTests(t, tests)
}
