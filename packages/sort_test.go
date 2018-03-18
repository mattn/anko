package packages

import (
	"os"
	"testing"
)

func TestSort(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "sort = import(\"sort\"); a = make([]int); a += [5, 3, 1, 4, 2]; sort.Ints(a); a", runOutput: []int{1, 2, 3, 4, 5}, output: map[string]interface{}{"a": []int{1, 2, 3, 4, 5}}},
		{script: "sort = import(\"sort\"); a = make([]float64); a += [5.5, 3.3, 1.1, 4.4, 2.2]; sort.Float64s(a); a", runOutput: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, output: map[string]interface{}{"a": []float64{1.1, 2.2, 3.3, 4.4, 5.5}}},
		{script: "sort = import(\"sort\"); a = make([]string); a += [\"e\", \"c\", \"a\", \"d\", \"b\"]; sort.Strings(a); a", runOutput: []string{"a", "b", "c", "d", "e"}, output: map[string]interface{}{"a": []string{"a", "b", "c", "d", "e"}}},
		{script: `
sort = import("sort")
a = [5, 1.1, 3, "f", "2", "4.4"]
sortFuncs = make(sort.SortFuncsStruct)
sortFuncs.LenFunc = func() { return len(a) }
sortFuncs.LessFunc = func(i, j) { return a[i] < a[j] }
sortFuncs.SwapFunc = func(i, j) { temp = a[i]; a[i] = a[j]; a[j] = temp }
sort.Sort(sortFuncs)
a
`,
			runOutput: []interface{}{"f", float64(1.1), "2", int64(3), "4.4", int64(5)}, output: map[string]interface{}{"a": []interface{}{"f", float64(1.1), "2", int64(3), "4.4", int64(5)}}},
	}
	runTests(t, tests)
}
