package vm_test

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func Example_vmFunctions() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
func a(b) {
	println(b)
}
a("b")

a = func(b) {
	println(b)
}
a("b")

func(b) {
	println(b)
}("b")

func a() {
	return "a"
}
println(a())

println("")


func fib(n) {
	if (n <= 1) {
		return n
	}
	return fib(n - 1) + fib(n - 2) 
}

println(fib(8))
 

`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// b
	// b
	// b
	// a
	//
	// 21

}
