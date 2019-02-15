package vm_test

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func Example_vmArrays() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
a = []interface{1, 2}
println(a)

a += 3
println(a)

a = []interface{}
// this automatically appends to array
a[0] = 1
println(a)

println("")

a = []interface{}
// this would give an index out of range error
// a[1] = 1

a = []interface{1, 2}
b = []interface{3, 4}
c = a + b
println(c)

c = []interface{1, 2} + []interface{3, 4}
println(c)

println("")

c = []interface{a} + b
println(c)

c = []interface{a} + []interface{b}
println(c)

c = []interface{[]interface{1, 2}} + []interface{[]interface{3, 4}}
println(c)

println("")

a = []interface{1, 2}

println(len(a))

println(a[1])

a = [1, 2]
println(a)
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// [1 2]
	// [1 2 3]
	// [1]
	//
	// [1 2 3 4]
	// [1 2 3 4]
	//
	// [[1 2] 3 4]
	// [[1 2] [3 4]]
	// [[1 2] [3 4]]
	//
	// 2
	// 2
	// [1 2]
}

func Example_vmMaps() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
a = map[interface]interface{}
println(a)

a.b = 1
println(a)
println(a.b)

a["b"] = 2
println(a["b"])

println(len(a))

println("")

b, ok = a["b"]
println(b)
println(ok)

delete(a, "b")

_, ok = a["b"]
println(ok)

println("")

a = {}
println(a)

a.b = 1
println(a)
println(a.b)

a["b"] = 2
println(a["b"])

`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// map[]
	// map[b:1]
	// 1
	// 2
	// 1
	//
	// 2
	// true
	// false
	//
	// map[]
	// map[b:1]
	// 1
	// 2
}

func Example_vmModules() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
module rectangle {
	_length = 1
	_width = 1
	
	func setLength (length) {
		if length <= 0 {
			return
		}
		_length = length
	}
	
	func setWidth (width) {
		if width <= 0 {
			return
		}
		_width = width
	}
	
	func area () {
		return _length * _width
	}
	
	func perimeter () {
		return 2 * (_length + _width)
	}
 }

rectangle1 = rectangle

rectangle1.setLength(4)
rectangle1.setWidth(5)

println(rectangle1.area())
println(rectangle1.perimeter())

rectangle2 = rectangle

rectangle2.setLength(2)
rectangle2.setWidth(4)

println(rectangle2.area())
println(rectangle2.perimeter())
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// 20
	// 18
	// 8
	// 12
}
