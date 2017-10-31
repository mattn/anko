package vm

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func testInterrupt(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	waitChan := make(chan struct{}, 1)

	env := NewEnv()
	sleepMillisecond := func() { time.Sleep(time.Millisecond) }

	err := env.Define("println", fmt.Println)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	err = env.Define("sleep", sleepMillisecond)
	if err != nil {
		t.Errorf("Define error: %v", err)
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
		_, err := env.Execute(script)
		if err == nil {
			t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
		} else if err.Error() != InterruptError.Error() {
			t.Errorf("Execute error - received %v expected: %v", err, InterruptError)
		}
		waitGroup.Done()
	}()

	<-waitChan
	Interrupt(env)

	waitGroup.Wait()
}

func TestInterruptRaces(t *testing.T) {
	// Run testInterrupt many times
	var waitGroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func() {
			testInterrupt(t)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}
