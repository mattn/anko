package env

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestSetError(t *testing.T) {
	envParent := NewEnv()
	envChild := envParent.NewEnv()
	err := envChild.Set("a", "a")
	if err == nil {
		t.Errorf("Set error - received: %v - expected: %v", err, fmt.Errorf("undefined symbol 'a'"))
	} else if err.Error() != "undefined symbol 'a'" {
		t.Errorf("Set error - received: %v - expected: %v", err, fmt.Errorf("undefined symbol 'a'"))
	}
}

func TestAddrError(t *testing.T) {
	envParent := NewEnv()
	envChild := envParent.NewEnv()
	_, err := envChild.Addr("a")
	if err == nil {
		t.Errorf("Addr error - received: %v - expected: %v", err, fmt.Errorf("undefined symbol 'a'"))
	} else if err.Error() != "undefined symbol 'a'" {
		t.Errorf("Addr error - received: %v - expected: %v", err, fmt.Errorf("undefined symbol 'a'"))
	}
}

func TestDefineGlobalValue(t *testing.T) {
	envParent := NewEnv()
	envChild := envParent.NewEnv()
	err := envChild.DefineGlobalValue("a", reflect.ValueOf("a"))
	if err != nil {
		t.Fatal("DefineGlobalValue error:", err)
	}

	var value interface{}
	value, err = envParent.Get("a")
	if err != nil {
		t.Fatal("Get error:", err)
	}
	v, ok := value.(string)
	if !ok {
		t.Fatalf("value - received: %T - expected: %T", value, "a")
	}
	if v != "a" {
		t.Fatalf("value - received: %v - expected: %v", v, "a")
	}
}

func TestDefineAndGet(t *testing.T) {
	var err error
	var value interface{}
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: reflect.Value{}, varGetValue: reflect.Value{}, varKind: reflect.Invalid},
		{testInfo: "nil", varName: "a", varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface},
		{testInfo: "bool", varName: "a", varDefineValue: true, varGetValue: true, varKind: reflect.Bool},
		{testInfo: "int16", varName: "a", varDefineValue: int16(1), varGetValue: int16(1), varKind: reflect.Int16},
		{testInfo: "int32", varName: "a", varDefineValue: int32(1), varGetValue: int32(1), varKind: reflect.Int32},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64},
		{testInfo: "uint32", varName: "a", varDefineValue: uint32(1), varGetValue: uint32(1), varKind: reflect.Uint32},
		{testInfo: "uint64", varName: "a", varDefineValue: uint64(1), varGetValue: uint64(1), varKind: reflect.Uint64},
		{testInfo: "float32", varName: "a", varDefineValue: float32(1), varGetValue: float32(1), varKind: reflect.Float32},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64},
		{testInfo: "string", varName: "a", varDefineValue: "a", varGetValue: "a", varKind: reflect.String},

		{testInfo: "string with dot", varName: "a.a", varDefineValue: "a", varGetValue: nil, varKind: reflect.Interface, defineError: ErrSymbolContainsDot, getError: fmt.Errorf("undefined symbol 'a.a'")},
		{testInfo: "string with quotes", varName: "a", varDefineValue: `"a"`, varGetValue: `"a"`, varKind: reflect.String},
	}

	// DefineAndGet
	for _, test := range tests {
		env := NewEnv()

		err = env.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = env.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}
	}

	// DefineAndGet NewEnv
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet NewEnv %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet NewEnv %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envChild.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet NewEnv %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet NewEnv %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet NewEnv %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}
	}

	// DefineAndGet DefineGlobal
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envChild.DefineGlobal(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineAndGet DefineGlobal %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineAndGet DefineGlobal %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envParent.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineAndGet DefineGlobal %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineAndGet DefineGlobal %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineAndGet DefineGlobal %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}
	}

}

func TestGetValueSymbols(t *testing.T) {
	var symbols []string
	values := map[string]interface{}{
		"a": int64(1),
		"b": float64(1),
		"c": true,
		"d": "a",
	}

	env := NewEnv()
	for s, v := range values {
		env.Define(s, v)
	}

	symbols = env.GetValueSymbols()
	if len(symbols) != len(values) {
		t.Errorf("Expected %d symbols, received %d", len(values), len(symbols))
	}

	for _, symbol := range symbols {
		_, ok := values[symbol]
		if !ok {
			t.Errorf("Missing %s symbol", symbol)
		}
	}
}

func TestDefineModify(t *testing.T) {
	var err error
	var value interface{}
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface},
		{testInfo: "bool", varName: "a", varDefineValue: true, varGetValue: true, varKind: reflect.Bool},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64},
		{testInfo: "string", varName: "a", varDefineValue: "a", varGetValue: "a", varKind: reflect.String},
	}
	changeTests := []struct {
		varDefineValue interface{}
		varGetValue    interface{}
		varKind        reflect.Kind
		defineError    error
		getError       error
	}{
		{varDefineValue: nil, varGetValue: nil, varKind: reflect.Interface},
		{varDefineValue: "a", varGetValue: "a", varKind: reflect.String},
		{varDefineValue: int64(1), varGetValue: int64(1), varKind: reflect.Int64},
		{varDefineValue: float64(1), varGetValue: float64(1), varKind: reflect.Float64},
		{varDefineValue: true, varGetValue: true, varKind: reflect.Bool},
	}

	// DefineModify
	for _, test := range tests {
		env := NewEnv()

		err = env.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineModify %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineModify %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = env.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineModify %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineModify %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineModify %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		// DefineModify changeTest
		for _, changeTest := range changeTests {
			err = env.Set(test.varName, changeTest.varDefineValue)
			if err != nil && changeTest.defineError != nil {
				if err.Error() != changeTest.defineError.Error() {
					t.Errorf("DefineModify changeTest %v - Set error - received: %v - expected: %v", test.testInfo, err, changeTest.defineError)
					continue
				}
			} else if err != changeTest.defineError {
				t.Errorf("DefineModify changeTest %v - Set error - received: %v - expected: %v", test.testInfo, err, changeTest.defineError)
				continue
			}

			value, err = env.Get(test.varName)
			if err != nil && changeTest.getError != nil {
				if err.Error() != changeTest.getError.Error() {
					t.Errorf("DefineModify changeTest  %v - Get error - received: %v - expected: %v", test.testInfo, err, changeTest.getError)
					continue
				}
			} else if err != changeTest.getError {
				t.Errorf("DefineModify changeTest  %v - Get error - received: %v - expected: %v", test.testInfo, err, changeTest.getError)
				continue
			}
			if value != changeTest.varGetValue {
				t.Errorf("DefineModify changeTest  %v - value check - received %#v expected: %#v", test.testInfo, value, changeTest.varGetValue)
			}
		}
	}

	// DefineModify envParent
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineModify envParent %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineModify envParent %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envChild.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineModify envParent  %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineModify envParent  %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineModify envParent  %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		for _, changeTest := range changeTests {
			err = envParent.Set(test.varName, changeTest.varDefineValue)
			if err != nil && changeTest.defineError != nil {
				if err.Error() != changeTest.defineError.Error() {
					t.Errorf("DefineModify envParent changeTest %v - Set error - received: %v - expected: %v", test.testInfo, err, changeTest.defineError)
					continue
				}
			} else if err != changeTest.defineError {
				t.Errorf("DefineModify envParent changeTest %v - Set error - received: %v - expected: %v", test.testInfo, err, changeTest.defineError)
				continue
			}

			value, err = envChild.Get(test.varName)
			if err != nil && changeTest.getError != nil {
				if err.Error() != changeTest.getError.Error() {
					t.Errorf("DefineModify envParent changeTest %v - Get error - received: %v - expected: %v", test.testInfo, err, changeTest.getError)
					continue
				}
			} else if err != changeTest.getError {
				t.Errorf("ChanDefineModify envParent changeTestgeTest %v - Get error - received: %v - expected: %v", test.testInfo, err, changeTest.getError)
				continue
			}
			if value != changeTest.varGetValue {
				t.Errorf("DefineModify envParent changeTest %v - value check - received %#v expected: %#v", test.testInfo, value, changeTest.varGetValue)
			}
		}
	}

	// DefineModify envChild
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("DefineModify envChild %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("DefineModify envChild %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		value, err = envChild.Get(test.varName)
		if err != nil && test.getError != nil {
			if err.Error() != test.getError.Error() {
				t.Errorf("DefineModify envChild  %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
				continue
			}
		} else if err != test.getError {
			t.Errorf("DefineModify envChild  %v - Get error - received: %v - expected: %v", test.testInfo, err, test.getError)
			continue
		}
		if value != test.varGetValue {
			t.Errorf("DefineModify envChild  %v - value check - received %#v expected: %#v", test.testInfo, value, test.varGetValue)
		}

		for _, changeTest := range changeTests {
			err = envChild.Set(test.varName, changeTest.varDefineValue)
			if err != nil && changeTest.defineError != nil {
				if err.Error() != changeTest.defineError.Error() {
					t.Errorf("DefineModify envChild changeTest %v - Set error - received: %v - expected: %v", test.testInfo, err, changeTest.defineError)
					continue
				}
			} else if err != changeTest.defineError {
				t.Errorf("DefineModify envChild changeTest %v - Set error - received: %v - expected: %v", test.testInfo, err, changeTest.defineError)
				continue
			}

			value, err = envChild.Get(test.varName)
			if err != nil && changeTest.getError != nil {
				if err.Error() != changeTest.getError.Error() {
					t.Errorf("DefineModify envChild changeTest %v - Get error - received: %v - expected: %v", test.testInfo, err, changeTest.getError)
					continue
				}
			} else if err != changeTest.getError {
				t.Errorf("ChanDefineModify envChild changeTestgeTest %v - Get error - received: %v - expected: %v", test.testInfo, err, changeTest.getError)
				continue
			}
			if value != changeTest.varGetValue {
				t.Errorf("DefineModify envChild changeTest %v - value check - received %#v expected: %#v", test.testInfo, value, changeTest.varGetValue)
			}
		}
	}
}

func TestAddr(t *testing.T) {
	var err error
	tests := []struct {
		testInfo       string
		varName        string
		varDefineValue interface{}
		defineError    error
		addrError      error
	}{
		{testInfo: "nil", varName: "a", varDefineValue: nil, addrError: nil},
		{testInfo: "string", varName: "a", varDefineValue: "a", addrError: fmt.Errorf("unaddressable")},
		{testInfo: "int64", varName: "a", varDefineValue: int64(1), addrError: fmt.Errorf("unaddressable")},
		{testInfo: "float64", varName: "a", varDefineValue: float64(1), addrError: fmt.Errorf("unaddressable")},
		{testInfo: "bool", varName: "a", varDefineValue: true, addrError: fmt.Errorf("unaddressable")},
	}

	// TestAddr
	for _, test := range tests {
		envParent := NewEnv()
		envChild := envParent.NewEnv()

		err = envParent.Define(test.varName, test.varDefineValue)
		if err != nil && test.defineError != nil {
			if err.Error() != test.defineError.Error() {
				t.Errorf("TestAddr %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
				continue
			}
		} else if err != test.defineError {
			t.Errorf("TestAddr %v - Define error - received: %v - expected: %v", test.testInfo, err, test.defineError)
			continue
		}

		_, err = envChild.Addr(test.varName)
		if err != nil && test.addrError != nil {
			if err.Error() != test.addrError.Error() {
				t.Errorf("TestAddr %v - Addr error - received: %v - expected: %v", test.testInfo, err, test.addrError)
				continue
			}
		} else if err != test.addrError {
			t.Errorf("TestAddr %v - Addr error - received: %v - expected: %v", test.testInfo, err, test.addrError)
			continue
		}
	}
}

func TestDelete(t *testing.T) {
	// empty
	env := NewEnv()
	env.Delete("a")

	// add & delete a
	env.Define("a", "a")
	env.Delete("a")

	value, err := env.Get("a")
	expectedError := "undefined symbol 'a'"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Get error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("Get value - received: %#v - expected: %#v", value, nil)
	}
}

func TestDeleteGlobal(t *testing.T) {
	// empty
	env := NewEnv()
	env.DeleteGlobal("a")

	// add & delete a
	env.Define("a", "a")
	env.DeleteGlobal("a")

	value, err := env.Get("a")
	expectedError := "undefined symbol 'a'"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Get error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("Get value - received: %#v - expected: %#v", value, nil)
	}

	// parent & child, var in child, delete in parent
	envChild := env.NewEnv()
	envChild.Define("a", "a")
	env.DeleteGlobal("a")

	value, err = envChild.Get("a")
	if err != nil {
		t.Errorf("Get error - received: %v - expected: %v", err, nil)
	}
	if value != "a" {
		t.Errorf("Get value - received: %#v - expected: %#v", value, "a")
	}

	// parent & child, var in child, delete in child
	envChild.DeleteGlobal("a")

	value, err = envChild.Get("a")
	if err == nil || err.Error() != expectedError {
		t.Errorf("Get error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("Get value - received: %#v - expected: %#v", value, nil)
	}

	// parent & child, var in parent, delete in child
	env.Define("a", "a")
	envChild.DeleteGlobal("a")

	value, err = envChild.Get("a")
	if err == nil || err.Error() != expectedError {
		t.Errorf("Get error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("Get value - received: %#v - expected: %#v", value, nil)
	}

	// parent & child, var in parent, delete in parent
	env.Define("a", "a")
	env.DeleteGlobal("a")

	value, err = envChild.Get("a")
	if err == nil || err.Error() != expectedError {
		t.Errorf("Get error - received: %v - expected: %v", err, expectedError)
	}
	if value != nil {
		t.Errorf("Get value - received: %#v - expected: %#v", value, nil)
	}
}

func TestRaceCreateSameVariable(t *testing.T) {
	// Test creating same variable in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			err := env.Define("a", i)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			_, err = env.Get("a")
			if err != nil {
				t.Errorf("Get error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()

	_, err := env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func TestRaceCreateDifferentVariables(t *testing.T) {
	// Test creating different variables in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			err := env.Define(fmt.Sprint(i), i)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			_, err = env.Get(fmt.Sprint(i))
			if err != nil {
				t.Errorf("Get error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()

	for i := 0; i < 100; i++ {
		_, err := env.Get(fmt.Sprint(i))
		if err != nil {
			t.Errorf("Get error: %v", err)
		}
	}
}

func TestRaceReadDifferentVariables(t *testing.T) {
	// Test reading different variables in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	for i := 0; i < 100; i++ {
		err := env.Define(fmt.Sprint(i), i)
		if err != nil {
			t.Errorf("Define error: %v", err)
		}
		_, err = env.Get(fmt.Sprint(i))
		if err != nil {
			t.Errorf("Get error: %v", err)
		}
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			_, err := env.Get(fmt.Sprint(i))
			if err != nil {
				t.Errorf("Get error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()
}

func TestRaceSetSameVariable(t *testing.T) {
	// Test setting same variable in parallel

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	err := env.Define("a", 0)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	_, err = env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			err := env.Set("a", i)
			if err != nil {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}

	close(waitChan)
	waitGroup.Wait()

	_, err = env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func TestRaceSetSameVariableNewEnv(t *testing.T) {
	// Test setting same variable in parallel with NewEnv

	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	env := NewEnv()

	err := env.Define("a", 0)
	if err != nil {
		t.Errorf("Define error: %v", err)
	}
	_, err = env.Get("a")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go func(i int) {
			<-waitChan
			env = env.NewEnv().NewEnv()
			err := env.Set("a", i)
			if err != nil {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}(i)
	}
}

func TestRaceDefineAndSetSameVariable(t *testing.T) {
	// Test defining and setting same variable in parallel
	for i := 0; i < 100; i++ {
		raceDefineAndSetSameVariable(t)
	}
}

func raceDefineAndSetSameVariable(t *testing.T) {
	waitChan := make(chan struct{}, 1)
	var waitGroup sync.WaitGroup

	envParent := NewEnv()
	envChild := envParent.NewEnv()

	for i := 0; i < 2; i++ {
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envParent.Set("a", 1)
			if err != nil && err.Error() != "undefined symbol 'a'" {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}()
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envParent.Define("a", 2)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			waitGroup.Done()
		}()
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envChild.Set("a", 3)
			if err != nil && err.Error() != "undefined symbol 'a'" {
				t.Errorf("Set error: %v", err)
			}
			waitGroup.Done()
		}()
		waitGroup.Add(1)
		go func() {
			<-waitChan
			err := envChild.Define("a", 4)
			if err != nil {
				t.Errorf("Define error: %v", err)
			}
			waitGroup.Done()
		}()
	}

	close(waitChan)
	waitGroup.Wait()

	_, err := envParent.Get("a") // value of a could be 1, 2, or 3
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	_, err = envChild.Get("a") // value of a could be 3 or 4
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
}

func BenchmarkDefine(b *testing.B) {
	var err error
	env := NewEnv()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := env.Define("a", 1)
		if err != nil {
			b.Errorf("Set error: %v", err)
		}
	}
	b.StopTimer()
	_, err = env.Get("a")
	if err != nil {
		b.Errorf("Get error: %v", err)
	}
}

func BenchmarkSet(b *testing.B) {
	env := NewEnv()
	err := env.Define("a", 1)
	if err != nil {
		b.Errorf("Define error: %v", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = env.Set("a", 1)
		if err != nil {
			b.Errorf("Set error: %v", err)
		}
	}
	b.StopTimer()
	_, err = env.Get("a")
	if err != nil {
		b.Errorf("Get error: %v", err)
	}
}
