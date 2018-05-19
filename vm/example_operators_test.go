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
		log.Fatalf("execute error: %v\n", err)
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
		log.Fatalf("execute error: %v\n", err)
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

func Example_vmIfOperators() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
a = 1
b = 2

if a == 1 {
	println(a)
}

if b == 1 {
	println(a)
} else {
	println(b)
}

if a == 3 {
	println(a)
} else if b == 3 {
	println(b)
} else {
	println(a + b)
}

println("")

if a == 2 || b == 2 {
	println(4)
}

if a == 1 && b == 2 {
	println(5)
}
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// 1
	// 2
	// 3
	//
	// 4
	// 5

}

func Example_vmForLoops() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
i = 0
for {
	println(i)
	i++
	if i > 1 {
		break
	}
}

println("")

for i in [0, 1] {
	println(i)
}

println("")

i = 0
for i < 2 {
	println(i)
	i++
}

println("")

for i = 0; i < 2; i++ {
	println(i)
}

`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// 0
	// 1
	//
	// 0
	// 1
	//
	// 0
	// 1
	//
	// 0
	// 1

}

func Example_vmSlices() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
a = "abc"
println(a[1:])
println(a[:2])
println(a[1:2])

println("")

a = [1, 2, 3]
println(a[1:])
println(a[:2])
println(a[1:2])
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// bc
	// ab
	// b
	//
	// [2 3]
	// [1 2]
	// [2]

}

func Example_vmChannels() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
a = make(chan string, 1)
a <- "a"
println(<- a)

a = make(chan string)
go func() {
	a <- "a"
}()
println(<- a)


`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// a
	// a

}
