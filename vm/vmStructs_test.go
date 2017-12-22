package vm

import (
	"fmt"
	"os"
	"testing"
)

func TestStructs(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: "a[\"B\"]", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			runError: fmt.Errorf("type struct does not support index operation"),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},
		{script: "a.C", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			runError: fmt.Errorf("no member named 'C' for struct"),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},

		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			runOutput: nil,
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: nil, B: nil}},
			runOutput: nil,
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: nil, B: nil}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: int32(1), B: int32(2)}},
			runOutput: int32(2),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int32(1), B: int32(2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			runOutput: int64(2),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: float32(1.1), B: float32(2.2)}},
			runOutput: float32(2.2),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: float32(1.1), B: float32(2.2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: float64(1.1), B: float64(2.2)}},
			runOutput: float64(2.2),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: float64(1.1), B: float64(2.2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: "a", B: "b"}},
			runOutput: "b",
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: "a", B: "b"}}},

		{script: "a.B", input: map[string]interface{}{"a": struct {
			A bool
			B bool
		}{}},
			runOutput: false,
			output: map[string]interface{}{"a": struct {
				A bool
				B bool
			}{}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A int32
			B int32
		}{}},
			runOutput: int32(0),
			output: map[string]interface{}{"a": struct {
				A int32
				B int32
			}{}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A int64
			B int64
		}{}},
			runOutput: int64(0),
			output: map[string]interface{}{"a": struct {
				A int64
				B int64
			}{}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A float32
			B float32
		}{}},
			runOutput: float32(0),
			output: map[string]interface{}{"a": struct {
				A float32
				B float32
			}{}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A float64
			B float64
		}{}},
			runOutput: float64(0),
			output: map[string]interface{}{"a": struct {
				A float64
				B float64
			}{}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A string
			B string
		}{}},
			runOutput: "",
			output: map[string]interface{}{"a": struct {
				A string
				B string
			}{}}},

		{script: "a.B", input: map[string]interface{}{"a": struct {
			A bool
			B bool
		}{A: true, B: true}},
			runOutput: true,
			output: map[string]interface{}{"a": struct {
				A bool
				B bool
			}{A: true, B: true}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A int32
			B int32
		}{A: int32(1), B: int32(2)}},
			runOutput: int32(2),
			output: map[string]interface{}{"a": struct {
				A int32
				B int32
			}{A: int32(1), B: int32(2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A int64
			B int64
		}{A: int64(1), B: int64(2)}},
			runOutput: int64(2),
			output: map[string]interface{}{"a": struct {
				A int64
				B int64
			}{A: int64(1), B: int64(2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A float32
			B float32
		}{A: float32(1.1), B: float32(2.2)}},
			runOutput: float32(2.2),
			output: map[string]interface{}{"a": struct {
				A float32
				B float32
			}{A: float32(1.1), B: float32(2.2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A float64
			B float64
		}{A: float64(1.1), B: float64(2.2)}},
			runOutput: float64(2.2),
			output: map[string]interface{}{"a": struct {
				A float64
				B float64
			}{A: float64(1.1), B: float64(2.2)}}},
		{script: "a.B", input: map[string]interface{}{"a": struct {
			A string
			B string
		}{A: "a", B: "b"}},
			runOutput: "b",
			output: map[string]interface{}{"a": struct {
				A string
				B string
			}{A: "a", B: "b"}}},

		{script: "a.C = 3", input: map[string]interface{}{
			"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)},
		},
			runError: fmt.Errorf("no member named 'C' for struct"),
			output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)}}},

		// TOFIX:
		//		{script: "a.B = 3", input: map[string]interface{}{
		//			"a": struct {
		//				A interface{}
		//				B interface{}
		//			}{A: int64(1), B: int64(2)},
		//		},
		//			runOutput: int64(3),
		//			output: map[string]interface{}{"a": struct {
		//				A interface{}
		//				B interface{}
		//			}{A: int64(1), B: int64(3)}}},
	}
	runTests(t, tests)
}
