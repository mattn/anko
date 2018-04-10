package vm

import (
	"fmt"
	"os"
	"testing"
)

func TestStructs(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `a["B"]`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			RunError: fmt.Errorf("type struct does not support index operation"),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},
		{Script: `a.C`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			RunError: fmt.Errorf("no member named 'C' for struct"),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},

		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: nil, B: nil}},
			RunOutput: nil,
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: nil, B: nil}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: int32(1), B: int32(2)}},
			RunOutput: int32(2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int32(1), B: int32(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			RunOutput: int64(2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: float32(1.1), B: float32(2.2)}},
			RunOutput: float32(2.2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: float32(1.1), B: float32(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: float64(1.1), B: float64(2.2)}},
			RunOutput: float64(2.2),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: float64(1.1), B: float64(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A interface{}
			B interface{}
		}{A: "a", B: "b"}},
			RunOutput: "b",
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: "a", B: "b"}}},

		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A bool
			B bool
		}{}},
			RunOutput: false,
			Output: map[string]interface{}{"a": struct {
				A bool
				B bool
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int32
			B int32
		}{}},
			RunOutput: int32(0),
			Output: map[string]interface{}{"a": struct {
				A int32
				B int32
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int64
			B int64
		}{}},
			RunOutput: int64(0),
			Output: map[string]interface{}{"a": struct {
				A int64
				B int64
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float32
			B float32
		}{}},
			RunOutput: float32(0),
			Output: map[string]interface{}{"a": struct {
				A float32
				B float32
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float64
			B float64
		}{}},
			RunOutput: float64(0),
			Output: map[string]interface{}{"a": struct {
				A float64
				B float64
			}{}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A string
			B string
		}{}},
			RunOutput: "",
			Output: map[string]interface{}{"a": struct {
				A string
				B string
			}{}}},

		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A bool
			B bool
		}{A: true, B: true}},
			RunOutput: true,
			Output: map[string]interface{}{"a": struct {
				A bool
				B bool
			}{A: true, B: true}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int32
			B int32
		}{A: int32(1), B: int32(2)}},
			RunOutput: int32(2),
			Output: map[string]interface{}{"a": struct {
				A int32
				B int32
			}{A: int32(1), B: int32(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A int64
			B int64
		}{A: int64(1), B: int64(2)}},
			RunOutput: int64(2),
			Output: map[string]interface{}{"a": struct {
				A int64
				B int64
			}{A: int64(1), B: int64(2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float32
			B float32
		}{A: float32(1.1), B: float32(2.2)}},
			RunOutput: float32(2.2),
			Output: map[string]interface{}{"a": struct {
				A float32
				B float32
			}{A: float32(1.1), B: float32(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A float64
			B float64
		}{A: float64(1.1), B: float64(2.2)}},
			RunOutput: float64(2.2),
			Output: map[string]interface{}{"a": struct {
				A float64
				B float64
			}{A: float64(1.1), B: float64(2.2)}}},
		{Script: `a.B`, Input: map[string]interface{}{"a": struct {
			A string
			B string
		}{A: "a", B: "b"}},
			RunOutput: "b",
			Output: map[string]interface{}{"a": struct {
				A string
				B string
			}{A: "a", B: "b"}}},

		{Script: `a.C = 3`, Input: map[string]interface{}{
			"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)},
		},
			RunError: fmt.Errorf("no member named 'C' for struct"),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(2)}}},

		{Script: `a.B = 3`, Input: map[string]interface{}{"a": &struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			RunOutput: int64(3),
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(3)}}},

		{Script: `a.B = 3; a = *a`, Input: map[string]interface{}{"a": &struct {
			A interface{}
			B interface{}
		}{A: int64(1), B: int64(2)}},
			RunOutput: struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(3)},
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: int64(1), B: int64(3)}}},
	}
	RunTests(t, tests, nil)
}

func TestMakeStructs(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `make(struct)`, Types: map[string]interface{}{"struct": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: &struct {
				A interface{}
				B interface{}
			}{}},

		{Script: `a = make(struct)`, Types: map[string]interface{}{"struct": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: &struct {
				A interface{}
				B interface{}
			}{},
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B interface{}
			}{}}},

		{Script: `a = make(struct); a.A = 3; a.B = 4`, Types: map[string]interface{}{"struct": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(4),
			Output: map[string]interface{}{"a": &struct {
				A interface{}
				B interface{}
			}{A: interface{}(int64(3)), B: interface{}(int64(4))}}},

		{Script: `a = make(struct); a = *a; a.A = 3; a.B = 4`, Types: map[string]interface{}{"struct": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(4),
			Output: map[string]interface{}{"a": struct {
				A interface{}
				B interface{}
			}{A: interface{}(int64(3)), B: interface{}(int64(4))}}},

		{Script: `a = make(struct); a.A = func () { return 1 }; a.A()`, Types: map[string]interface{}{"struct": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(1)},
		{Script: `a = make(struct); a.A = func () { return 1 }; a = *a; a.A()`, Types: map[string]interface{}{"struct": &struct {
			A interface{}
			B interface{}
		}{}},
			RunOutput: int64(1)},
	}
	RunTests(t, tests, nil)
}
