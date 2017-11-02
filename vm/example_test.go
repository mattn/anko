package vm_test

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mattn/anko/vm"
)

func ExampleInterrupt() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	waitChan := make(chan struct{}, 1)

	env := vm.NewEnv()
	sleepMillisecond := func() { time.Sleep(time.Millisecond) }

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}
	err = env.Define("sleep", sleepMillisecond)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
# sleep for 10 seconds
for i = 0; i < 10000; i++ {
	sleep()
}
# Should interrupt before printing the next line
println("this line should not be printed")
`

	go func() {
		close(waitChan)
		v, err := env.Execute(script)
		fmt.Println(v, err)
		waitGroup.Done()
	}()

	<-waitChan
	vm.Interrupt(env)

	waitGroup.Wait()

	// output: <nil> Execution interrupted
}

func ExampleEnv_Dump() {
	env := vm.NewEnv()

	err := env.Define("a", "a")
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	_, err = env.Get("a")
	if err != nil {
		log.Fatalf("Get error: %v\n", err)
	}

	env.Dump()

	// output: a = "a"
}
