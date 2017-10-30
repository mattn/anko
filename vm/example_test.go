package vm_test

import (
	"fmt"
	"log"
	"time"

	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

func ExamplePrintHelloWorld() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
println("Hello World :)")
`

	// Execute script
	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("Run error: %v\n", err)
	}

	// output: Hello World :)
}

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
		log.Fatalf("ParseSrc error: %v\n", err)
	}

	// Interrupts after 1 second.
	go func() {
		time.Sleep(time.Second)
		vm.Interrupt(env)
	}()

	// Run script
	v, err := vm.Run(stmts, env)

	fmt.Println(v, err)

	// output: <nil> Execution interrupted
}
