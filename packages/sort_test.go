package packages

import (
	"os"
	"testing"
)

func TestSort(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "sort = import(\"sort\"); a = make([]\"int\"); a += [5, 3, 1, 4, 2]; sort.Ints(a); a", runOutput: []int{1, 2, 3, 4, 5}, output: map[string]interface{}{"a": []int{1, 2, 3, 4, 5}}},
		{script: "sort = import(\"sort\"); a = make([]\"float64\"); a += [5.5, 3.3, 1.1, 4.4, 2.2]; sort.Float64s(a); a", runOutput: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, output: map[string]interface{}{"a": []float64{1.1, 2.2, 3.3, 4.4, 5.5}}},
		{script: "sort = import(\"sort\"); a = make([]\"string\"); a += [\"e\", \"c\", \"a\", \"d\", \"b\"]; sort.Strings(a); a", runOutput: []string{"a", "b", "c", "d", "e"}, output: map[string]interface{}{"a": []string{"a", "b", "c", "d", "e"}}},
	}
	runTests(t, tests)
}
