package vm_test

import (
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"log"
	"testing"
	"time"
)

func ExampleInterrupt() {
	env := vm.NewEnv()

	var sleepFunc = func(spec string) {
		if d, err := time.ParseDuration(spec); err != nil {
			panic(err)
		} else {
			time.Sleep(d)
		}
	}

	env.Define("println", fmt.Println)
	env.Define("sleep", sleepFunc)

	script := `
sleep("2s")
# Should interrupt here.
# The next line will not be executed.
println("<this should not be printed>")
`
	stmts, err := parser.ParseSrc(script)
	if err != nil {
		log.Fatal()
	}

	// Interrupts after 1 second.
	go func() {
		time.Sleep(time.Second)
		vm.Interrupt(env)
	}()

	// Run script
	_, err = vm.Run(stmts, env)

	//v, err := vm.Run(stmts, env)
	//fmt.Println(v, err)
	// output:
	// <nil> Execution interrupted
}

func TestExampleRaces(t *testing.T) {
	// Run example several times
	for i := 0; i < 100; i++ {
		go ExampleInterrupt()
	}
}
