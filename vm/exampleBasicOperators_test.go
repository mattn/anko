package vm_test

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func Example_vmBasicOperators() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
a = nil
println(a)
a = true
println(a)

println("")

a = 2 + 1
println(a)
a = 2 - 1
println(a)
a = 2 * 1
println(a)
a = 4 / 2
println(a)

println("")

a = 1
a++
println(a)
a--
println(a)

println("")

a = 1
a += 1
println(a)
a -= 1
println(a)
a *= 4
println(a)
a /= 2
println(a)

println("")

a = 2 ** 3
println(a)
a = 1 & 3
println(a)
a = 1 | 2
println(a) 

println("")

a = 2 << 3
println(a)
a = 8 >> 2
println(a)
a = 7 % 3
println(a)

println("")

a = 2 - (-2)
println(a)
a = ^2
println(a)
a = "a" * 4
println(a)
a = !true
println(a)

`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	// output:
	// <nil>
	// true
	//
	// 3
	// 1
	// 2
	// 2
	//
	// 2
	// 1
	//
	// 2
	// 1
	// 4
	// 2
	//
	// 8
	// 1
	// 3
	//
	// 16
	// 2
	// 1
	//
	// 4
	// -3
	// aaaa
	// false

}
