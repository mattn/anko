package vm_test

import (
	"log"

	"github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
)

func Example_vmSort() {
	env := vm.NewEnv()
	packages.DefineImport(env)

	script := `
fmt = import("fmt")
sort = import("sort")
a = [5, 1.1, 3, "f", "2", "4.4"]
sortFuncs = make(sort.SortFuncsStruct)
sortFuncs.LenFunc = func() { return len(a) }
sortFuncs.LessFunc = func(i, j) { return a[i] < a[j] }
sortFuncs.SwapFunc = func(i, j) { temp = a[i]; a[i] = a[j]; a[j] = temp }
sort.Sort(sortFuncs)
fmt.Println(a)
`

	_, err := env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// [f 1.1 2 3 4.4 5]
}
