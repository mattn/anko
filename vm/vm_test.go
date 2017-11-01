package vm

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/mattn/anko/parser"
)

type testStruct struct {
	script     string
	parseError error
	input      map[string]interface{}
	runError   error
	runOutput  interface{}
	ouput      map[string]interface{}
}

func TestBasicOperators(t *testing.T) {
	tests := []testStruct{
		{script: "]", parseError: fmt.Errorf("syntax error"), runOutput: (*interface{})(nil)},
		{script: "1 = 2", runError: fmt.Errorf("Invalid operation"), runOutput: []interface{}{interface{}(int64(2))}},

		{script: "2 + 1", runOutput: int64(3)},
		{script: "2 - 1", runOutput: int64(1)},
		{script: "2 * 1", runOutput: int64(2)},
		{script: "2 / 1", runOutput: float64(2)},
		{script: "2.1 + 1.1", runOutput: float64(3.2)},
		{script: "2.1 - 1.1", runOutput: float64(1)},
		{script: "2.1 * 2.0", runOutput: float64(4.2)},
		{script: "6.5 / 2.0", runOutput: float64(3.25)},

		{script: "a + b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(3)},
		{script: "a - b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(1)},
		{script: "a * b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(2)},
		{script: "a / b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: float64(2)},
		{script: "a + b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(3.2)},
		{script: "a - b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(1)},
		{script: "a * b", input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, runOutput: float64(4.2)},
		{script: "a / b", input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, runOutput: float64(3.25)},

		{script: "a + b", input: map[string]interface{}{"a": "a", "b": "b"}, runOutput: "ab"},
		{script: "a + b", input: map[string]interface{}{"a": "a", "b": int64(1)}, runOutput: "a1"},
		{script: "a + b", input: map[string]interface{}{"a": "a", "b": float64(1.1)}, runOutput: "a1.1"},
		{script: "a + z", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: (*interface{})(nil)},
		{script: "z + b", input: map[string]interface{}{"a": "a"}, runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: (*interface{})(nil)},

		{script: "a = nil", runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "a = nil; a = nil", runOutput: nil, ouput: map[string]interface{}{"a": nil}},
		{script: "a = b", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = b; b = nil", input: map[string]interface{}{"b": nil}, runOutput: nil, ouput: map[string]interface{}{"a": nil, "b": nil}},
		{script: "a = true", runOutput: true, ouput: map[string]interface{}{"a": true}},
		{script: "a = false", input: map[string]interface{}{"a": true}, runOutput: false, ouput: map[string]interface{}{"a": false}},
		{script: "a = b", input: map[string]interface{}{"b": true}, runOutput: true, ouput: map[string]interface{}{"a": true, "b": true}},
		{script: "a = b", input: map[string]interface{}{"b": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{script: "a = b", input: map[string]interface{}{"b": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2), "b": int64(2)}},
		{script: "a = b", input: map[string]interface{}{"b": float64(2.1)}, runOutput: float64(2.1), ouput: map[string]interface{}{"a": float64(2.1), "b": float64(2.1)}},
		{script: "y = z", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: (*interface{})(nil)},

		{script: "c = a + b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(3), ouput: map[string]interface{}{"c": int64(3)}},
		{script: "c = a - b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(1), ouput: map[string]interface{}{"c": int64(1)}},
		{script: "c = a * b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: int64(2), ouput: map[string]interface{}{"c": int64(2)}},
		{script: "c = a / b", input: map[string]interface{}{"a": int64(2), "b": int64(1)}, runOutput: float64(2), ouput: map[string]interface{}{"c": float64(2)}},
		{script: "c = a + b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(3.2), ouput: map[string]interface{}{"c": float64(3.2)}},
		{script: "c = a - b", input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, runOutput: float64(1), ouput: map[string]interface{}{"c": float64(1)}},
		{script: "c = a * b", input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, runOutput: float64(4.2), ouput: map[string]interface{}{"c": float64(4.2)}},
		{script: "c = a / b", input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, runOutput: float64(3.25), ouput: map[string]interface{}{"c": float64(3.25)}},

		{script: "a++", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a++", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(3.1), ouput: map[string]interface{}{"a": float64(3.1)}},
		{script: "a--", input: map[string]interface{}{"a": float64(2.1)}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "a++", input: map[string]interface{}{"a": int32(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": int32(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a++", input: map[string]interface{}{"a": float32(2.1)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a--", input: map[string]interface{}{"a": float32(2.1)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},

		{script: "1++", runError: fmt.Errorf("Invalid operation"), runOutput: (*interface{})(nil)},
		{script: "1--", runError: fmt.Errorf("Invalid operation"), runOutput: (*interface{})(nil)},
		{script: "z++", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: (*interface{})(nil)},
		{script: "z--", runError: fmt.Errorf("Undefined symbol 'z'"), runOutput: (*interface{})(nil)},

		{script: "a += 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a -= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), ouput: map[string]interface{}{"a": int64(1)}},
		{script: "a *= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), ouput: map[string]interface{}{"a": int64(4)}},
		{script: "a /= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: float64(1), ouput: map[string]interface{}{"a": float64(1)}},
		{script: "a += 1", input: map[string]interface{}{"a": 2.1}, runOutput: float64(3.1), ouput: map[string]interface{}{"a": float64(3.1)}},
		{script: "a -= 1", input: map[string]interface{}{"a": 2.1}, runOutput: float64(1.1), ouput: map[string]interface{}{"a": float64(1.1)}},
		{script: "a *= 2", input: map[string]interface{}{"a": 2.1}, runOutput: float64(4.2), ouput: map[string]interface{}{"a": float64(4.2)}},
		{script: "a /= 2", input: map[string]interface{}{"a": 6.5}, runOutput: float64(3.25), ouput: map[string]interface{}{"a": float64(3.25)}},

		{script: "a ** 3", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a ** 3", input: map[string]interface{}{"a": float64(2)}, runOutput: float64(8), ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a &= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(0)}},
		{script: "a &= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a &= 1", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(0)}},
		{script: "a &= 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a |= 1", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a |= 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a |= 1", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(3), ouput: map[string]interface{}{"a": int64(3)}},
		{script: "a |= 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a << 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a >> 2", input: map[string]interface{}{"a": int64(8)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(8)}},
		{script: "a << 2", input: map[string]interface{}{"a": float64(2)}, runOutput: int64(8), ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a >> 2", input: map[string]interface{}{"a": float64(8)}, runOutput: int64(2), ouput: map[string]interface{}{"a": float64(8)}},

		{script: "a % 2", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(0), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a % 3", input: map[string]interface{}{"a": int64(2)}, runOutput: int64(2), ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a % 2", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(0), ouput: map[string]interface{}{"a": float64(2.1)}},
		{script: "a % 3", input: map[string]interface{}{"a": float64(2.1)}, runOutput: int64(2), ouput: map[string]interface{}{"a": float64(2.1)}},
	}
	runTests(t, tests)
}

func TestComparisonOperators(t *testing.T) {
	tests := []testStruct{
		{script: "a == 1", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 1", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a != 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a == 1", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 1", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a != 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},

		{script: "a == nil", input: map[string]interface{}{"a": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil}},
		{script: "a == nil", input: map[string]interface{}{"a": nil}, runOutput: true, ouput: map[string]interface{}{"a": nil}},
		{script: "a == nil", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a == nil", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a == 2", input: map[string]interface{}{"a": nil}, runOutput: false, ouput: map[string]interface{}{"a": nil}},
		{script: "a == 2.0", input: map[string]interface{}{"a": nil}, runOutput: false, ouput: map[string]interface{}{"a": nil}},

		{script: "a > 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 1", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 2", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 3", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a > 1.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 2.0", input: map[string]interface{}{"a": int64(2)}, runOutput: false, ouput: map[string]interface{}{"a": int64(2)}},
		{script: "a < 3.0", input: map[string]interface{}{"a": int64(2)}, runOutput: true, ouput: map[string]interface{}{"a": int64(2)}},

		{script: "a > 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 1", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 2", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 3", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a > 1.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 2.0", input: map[string]interface{}{"a": float64(2)}, runOutput: false, ouput: map[string]interface{}{"a": float64(2)}},
		{script: "a < 3.0", input: map[string]interface{}{"a": float64(2)}, runOutput: true, ouput: map[string]interface{}{"a": float64(2)}},
	}
	runTests(t, tests)
}

func runTests(t *testing.T, tests []testStruct) {
	var value reflect.Value
loop:
	for _, test := range tests {
		stmts, err := parser.ParseSrc(test.script)
		if err != nil && test.parseError != nil {
			if err.Error() != test.parseError.Error() {
				t.Errorf("ParseSrc error - received: %v expected: %v - script: %v", err, test.parseError, test.script)
				continue
			}
		} else if err != test.parseError {
			t.Errorf("ParseSrc error - received: %v expected: %v - script: %v", err, test.parseError, test.script)
			continue
		}

		env := NewEnv()

		for inputName, inputValue := range test.input {
			err = env.Define(inputName, inputValue)
			if err != nil {
				t.Errorf("Define error: %v - inputName: %v - script: %v", err, inputName, test.script)
				continue loop
			}
		}

		value, err = Run(stmts, env)
		if err != nil && test.runError != nil {
			if err.Error() != test.runError.Error() {
				t.Errorf("Run error - received: %v expected: %v - script: %v", err, test.runError, test.script)
				continue
			}
		} else if err != test.runError {
			t.Errorf("Run error - received: %v expected: %v - script: %v", err, test.runError, test.script)
			continue
		}
		if !value.IsValid() || !value.CanInterface() {
			if !reflect.DeepEqual(value, test.runOutput) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
				continue
			}
		} else {
			if !reflect.DeepEqual(value.Interface(), test.runOutput) {
				t.Errorf("Run output - received: %#v expected: %#v - script: %v", value, test.runOutput, test.script)
				continue
			}
		}

		for outputName, outputValue := range test.ouput {
			value, err = env.Get(outputName)
			if err != nil {
				t.Errorf("Get error: %v - outputName: %v - script: %v", err, outputName, test.script)
				continue loop
			}
			if !value.IsValid() || !value.CanInterface() {
				if !reflect.DeepEqual(value, outputValue) {
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, outputValue, test.script)
					continue
				}
			} else {
				if !reflect.DeepEqual(value.Interface(), outputValue) {
					t.Errorf("outputName %v - received: %#v expected: %#v - script: %v", outputName, value, outputValue, test.script)
					continue
				}
			}
		}

		env.Destroy()
	}
}

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
