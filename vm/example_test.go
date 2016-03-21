package vm_test

import (
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"log"
	"reflect"
	"time"
)

func ExampleInterrupt() {
	env := vm.NewEnv()

	sleepFunc := func(spec string) {
		d, err := time.ParseDuration(spec)
		if err != nil {
			panic(err)
		}

		time.Sleep(d)
	}

	env.Define("println", reflect.ValueOf(fmt.Println))
	env.Define("sleep", reflect.ValueOf(sleepFunc))

	script := `
sleep("2s")
# Should interrupt here.
# The next line will not be executed.
println("<this should not be printed>")
`

	scanner := new(parser.Scanner)
	scanner.Init(script)
	stmts, err := parser.Parse(scanner)
	if err != nil {
		log.Fatal()
	}

	// Interrupts after 1 second.
	go func() {
		time.Sleep(time.Second * 1)
		vm.Interrupt(env)
	}()

	v, err := vm.Run(stmts, env)
	fmt.Println(v, err)
	// output:
	// <nil> Execution interrupted
}
