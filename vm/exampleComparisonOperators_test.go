package vm_test

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func Example_vmComparisonOperators() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
a = nil == nil
println(a)
a = "a" != "a" 
println(a)
a = 1 == 1.0
println(a)
a = !true
println(a)

println("")

a = 1 < 2
println(a)
a = 1 > 3
println(a)
a = 2 <= 2
println(a)
a = 2 >= 3
println(a)

println("")
a = 1 == 2 && 1 == 1
println(a)
a = 1 == 2 || 1 == 1
println(a)
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	// output:
	// true
	// false
	// true
	// false
	//
	// true
	// false
	// true
	// false
	//
	// false
	// true

}
