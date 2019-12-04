package vm

import (
	"fmt"
	"testing"

	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
)

func TestPackagesBytes(t *testing.T) {
	tests := []Test{
		{Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("a"); if err != nil { return err }; n`, RunOutput: 1},
		{Script: `bytes = import("bytes"); a = make(bytes.Buffer); n, err = a.WriteString("a"); if err != nil { return err }; a.String()`, RunOutput: "a"},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesJson(t *testing.T) {
	tests := []Test{
		{Script: `json = import("encoding/json"); a = make(map[string]interface); a["b"] = "b"; c, err = json.Marshal(a); err`, Output: map[string]interface{}{"a": map[string]interface{}{"b": "b"}, "c": []byte(`{"b":"b"}`)}},
		{Script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(a, &b); err`, Input: map[string]interface{}{"a": []byte(`{"b": "b"}`)}, Output: map[string]interface{}{"a": []byte(`{"b": "b"}`), "b": map[string]interface{}{"b": "b"}}},
		{Script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(a, &b); err`, Input: map[string]interface{}{"a": `{"b": "b"}`}, Output: map[string]interface{}{"a": `{"b": "b"}`, "b": map[string]interface{}{"b": "b"}}},
		{Script: `json = import("encoding/json"); b = 1; err = json.Unmarshal(a, &b); err`, Input: map[string]interface{}{"a": `[["1", "2"],["3", "4"]]`}, Output: map[string]interface{}{"a": `[["1", "2"],["3", "4"]]`, "b": []interface{}{[]interface{}{"1", "2"}, []interface{}{"3", "4"}}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesRegexp(t *testing.T) {
	tests := []Test{
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^simple$"); re.MatchString("simple")`, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^simple$"); re.MatchString("no match")`, RunOutput: false},

		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^simple$", "b": "simple"}, RunOutput: true, Output: map[string]interface{}{"a": "^simple$", "b": "simple"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^simple$", "b": "no match"}, RunOutput: false, Output: map[string]interface{}{"a": "^simple$", "b": "no match"}},

		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.String()`, RunOutput: "^a\\.\\d+\\.b$"},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a.1.b")`, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a.22.b")`, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a.333.b")`, RunOutput: true},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("no match")`, RunOutput: false},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile("^a\\.\\d+\\.b$"); re.MatchString("a+1+b")`, RunOutput: false},

		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.String()`, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$"}, RunOutput: "^a\\.\\d+\\.b$", Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.1.b"}, RunOutput: true, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.1.b"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.22.b"}, RunOutput: true, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.22.b"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.333.b"}, RunOutput: true, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a.333.b"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "no match"}, RunOutput: false, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "no match"}},
		{Script: `regexp = import("regexp"); re = regexp.MustCompile(a); re.MatchString(b)`, Input: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a+1+b"}, RunOutput: false, Output: map[string]interface{}{"a": "^a\\.\\d+\\.b$", "b": "a+1+b"}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesSort(t *testing.T) {
	tests := []Test{
		{Script: `sort = import("sort"); a = make([]int); a += [5, 3, 1, 4, 2]; sort.Ints(a); a`, RunOutput: []int{1, 2, 3, 4, 5}, Output: map[string]interface{}{"a": []int{1, 2, 3, 4, 5}}},
		{Script: `sort = import("sort"); a = make([]float64); a += [5.5, 3.3, 1.1, 4.4, 2.2]; sort.Float64s(a); a`, RunOutput: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, Output: map[string]interface{}{"a": []float64{1.1, 2.2, 3.3, 4.4, 5.5}}},
		{Script: `sort = import("sort"); a = make([]string); a += ["e", "c", "a", "d", "b"]; sort.Strings(a); a`, RunOutput: []string{"a", "b", "c", "d", "e"}, Output: map[string]interface{}{"a": []string{"a", "b", "c", "d", "e"}}},
		{Script: `
sort = import("sort")
a = [5, 1.1, 3, "f", "2", "4.4"]
sortFuncs = make(sort.SortFuncsStruct)
sortFuncs.LenFunc = func() { return len(a) }
sortFuncs.LessFunc = func(i, j) { return a[i] < a[j] }
sortFuncs.SwapFunc = func(i, j) { temp = a[i]; a[i] = a[j]; a[j] = temp }
sort.Sort(sortFuncs)
a
`,
			RunOutput: []interface{}{"f", float64(1.1), "2", int64(3), "4.4", int64(5)}, Output: map[string]interface{}{"a": []interface{}{"f", float64(1.1), "2", int64(3), "4.4", int64(5)}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesStrconv(t *testing.T) {
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
	tests := []Test{
		{Script: `strconv = import("strconv"); a = true; b = strconv.FormatBool(a)`, RunOutput: "true", Output: map[string]interface{}{"a": true, "b": "true"}},
		{Script: `strconv = import("strconv"); a = 1.1; b = strconv.FormatFloat(a, toRune("f"), -1, 64)`, Input: map[string]interface{}{"toRune": toRune}, RunOutput: "1.1", Output: map[string]interface{}{"a": float64(1.1), "b": "1.1"}},
		{Script: `strconv = import("strconv"); a = 1; b = strconv.FormatInt(a, 10)`, RunOutput: "1", Output: map[string]interface{}{"a": int64(1), "b": "1"}},
		{Script: `strconv = import("strconv"); b = strconv.FormatInt(a, 10)`, Input: map[string]interface{}{"a": uint64(1)}, RunOutput: "1", Output: map[string]interface{}{"a": uint64(1), "b": "1"}},

		{Script: `strconv = import("strconv"); a = "true"; b, err = strconv.ParseBool(a); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "true", "b": true, "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "2"; b, err = strconv.ParseBool(a); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseBool: parsing "2": invalid syntax`, Output: map[string]interface{}{"a": "2", "b": false, "err": `strconv.ParseBool: parsing "2": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1.1"; b, err = strconv.ParseFloat(a, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1.1", "b": float64(1.1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseFloat(a, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseFloat: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": float64(0), "err": `strconv.ParseFloat: parsing "a": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1", "b": int64(1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "1.1"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseInt: parsing "1.1": invalid syntax`, Output: map[string]interface{}{"a": "1.1", "b": int64(0), "err": `strconv.ParseInt: parsing "1.1": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseInt: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": int64(0), "err": `strconv.ParseInt: parsing "a": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1", "b": uint64(1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "a"; b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseUint: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": uint64(0), "err": `strconv.ParseUint: parsing "a": invalid syntax`}},

		{Script: `strconv = import("strconv"); a = "true"; var b, err = strconv.ParseBool(a); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "true", "b": true, "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "2"; var b, err = strconv.ParseBool(a); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseBool: parsing "2": invalid syntax`, Output: map[string]interface{}{"a": "2", "b": false, "err": `strconv.ParseBool: parsing "2": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1.1"; var b, err = strconv.ParseFloat(a, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1.1", "b": float64(1.1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "a"; var b, err = strconv.ParseFloat(a, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseFloat: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": float64(0), "err": `strconv.ParseFloat: parsing "a": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1"; var b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1", "b": int64(1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "1.1"; var b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseInt: parsing "1.1": invalid syntax`, Output: map[string]interface{}{"a": "1.1", "b": int64(0), "err": `strconv.ParseInt: parsing "1.1": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "a"; var b, err = strconv.ParseInt(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseInt: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": int64(0), "err": `strconv.ParseInt: parsing "a": invalid syntax`}},
		{Script: `strconv = import("strconv"); a = "1"; var b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: "<nil>", Output: map[string]interface{}{"a": "1", "b": uint64(1), "err": "<nil>"}},
		{Script: `strconv = import("strconv"); a = "a"; var b, err = strconv.ParseUint(a, 10, 64); err = toString(err)`, Input: map[string]interface{}{"toString": toString}, RunOutput: `strconv.ParseUint: parsing "a": invalid syntax`, Output: map[string]interface{}{"a": "a", "b": uint64(0), "err": `strconv.ParseUint: parsing "a": invalid syntax`}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesStrings(t *testing.T) {
	tests := []Test{
		{Script: `strings = import("strings"); a = " one two "; b = strings.TrimSpace(a)`, RunOutput: "one two", Output: map[string]interface{}{"a": " one two ", "b": "one two"}},
		{Script: `strings = import("strings"); a = "a b c d"; b = strings.Split(a, " ")`, RunOutput: []string{"a", "b", "c", "d"}, Output: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c", "d"}}},
		{Script: `strings = import("strings"); a = "a b c d"; b = strings.SplitN(a, " ", 3)`, RunOutput: []string{"a", "b", "c d"}, Output: map[string]interface{}{"a": "a b c d", "b": []string{"a", "b", "c d"}}},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesSync(t *testing.T) {
	tests := []Test{
		{Script: `sync = import("sync"); once = make(sync.Once); a = []; func add() { a += "a" }; once.Do(add); once.Do(add); a`, RunOutput: []interface{}{"a"}, Output: map[string]interface{}{"a": []interface{}{"a"}}},
		{Script: `sync = import("sync"); waitGroup = make(sync.WaitGroup); waitGroup.Add(2);  func done() { waitGroup.Done() }; go done(); go done(); waitGroup.Wait(); "a"`, RunOutput: "a"},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesTime(t *testing.T) {
	tests := []Test{
		{Script: `time = import("time"); a = make(time.Time); a.IsZero()`, RunOutput: true},
	}
	runTests(t, tests, nil, &Options{Debug: true})
}

func TestPackagesURL(t *testing.T) {
	e := env.NewEnv()
	value, err := Execute(e, nil, `
url = import("net/url")
v1 = make(url.Values)
v1.Set("a", "a")
if v1.Get("a") != "a" {
	return "value a not set"
} 
v2 = make(url.Values) 
v2.Set("b", "b")
if v2.Get("b") != "b" {
	return "value b not set"
} 
v2.Get("a")
`)
	if err != nil {
		t.Errorf("execute error - received: %v expected: %v", err, nil)
	}
	if value != "" {
		t.Errorf("execute value - received: %#v expected: %#v", value, "")
	}
}
