package vm_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func Example_vmExecuteContext() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	waitChan := make(chan struct{}, 1)

	e := env.NewEnv()
	sleepMillisecond := func() { time.Sleep(time.Millisecond) }

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}
	err = e.Define("sleep", sleepMillisecond)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
# sleep for 10 seconds
for i = 0; i < 10000; i++ {
	sleep()
}
# the context should cancel before printing the next line
println("this line should not be printed")
`

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		close(waitChan)
		v, err := vm.ExecuteContext(ctx, e, nil, script)
		fmt.Println(v, err)
		waitGroup.Done()
	}()

	<-waitChan
	cancel()

	waitGroup.Wait()

	// output: <nil> execution interrupted
}

func Example_vmEnvDefine() {
	// "github.com/mattn/anko/env"

	e := env.NewEnv()

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	err = e.Define("a", true)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}
	err = e.Define("b", int64(1))
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}
	err = e.Define("c", float64(1.1))
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}
	err = e.Define("d", "d")
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}
	err = e.Define("e", []interface{}{true, int64(1), float64(1.1), "d"})
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}
	err = e.Define("f", map[string]interface{}{"a": true})
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
println(a)
println(b)
println(c)
println(d)
println(e)
println(f)
`

	_, err = vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// true
	// 1
	// 1.1
	// d
	// [true 1 1.1 d]
	// map[a:true]
}

func Example_vmEnv() {
	// "github.com/mattn/anko/env"

	e := env.NewEnv()

	err := e.Define("a", "a")
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	_, err = e.Get("a")
	if err != nil {
		log.Fatalf("get error: %v\n", err)
	}

	fmt.Println(e)

	// output:
	// No parent
	// a = "a"
}

func Example_vmHelloWorld() {
	// "github.com/mattn/anko/env"

	e := env.NewEnv()

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
println("Hello World :)")
`

	_, err = vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output: Hello World :)
}

func Example_vmQuickStart() {
	// "github.com/mattn/anko/env"

	e := env.NewEnv()

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("define error: %v\n", err)
	}

	script := `
// declare variables
x = 1
y = x + 1

// print using outside the script defined println function
println(x + y) // 3

// if else statement
if x < 1 || y < 1 {
	println(x)
} else if x < 1 && y < 1 {
	println(y)
} else {
	println(x + y)
}

// slice
a = []interface{1, 2, 3}
println(a) // [1 2 3]
println(a[0]) // 1

// map
a = map[interface]interface{"x": 1}
println(a) // map[x:1]
a.b = 2
a["c"] = 3
println(a["b"]) // 2
println(a.c) // 3

// struct
a = make(struct {
	A int64,
	B float64
})
a.A = 4
a.B = 5.5
println(a.A) // 4
println(a.B) // 5.5

// function
func a (x) {
	println(x + 1)
}
a(5) // 6
`

	_, err = vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// 3
	// 3
	// [1 2 3]
	// 1
	// map[x:1]
	// 2
	// 3
	// 4
	// 5.5
	// 6
}
