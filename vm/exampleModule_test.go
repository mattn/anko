package vm_test

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func Example_vmModule() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
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
		log.Fatalf("Execute error: %v\n", err)
	}

	// output:
	// 20
	// 18
	// 8
	// 12
}
