package packages

import (
	"os"
	"testing"

	"github.com/mattn/anko/vm"
)

func TestSort(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []vm.Test{
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
	vm.RunTests(t, tests, &vm.TestingOptions{EnvSetupFunc: &testPackagesEnvSetupFunc})
}
