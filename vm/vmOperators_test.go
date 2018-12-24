package vm

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/mattn/anko/internal/testlib"
)

func TestBasicOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `]`, ParseError: fmt.Errorf("syntax error")},

		{Script: `2 + 1`, RunOutput: int64(3)},
		{Script: `2 - 1`, RunOutput: int64(1)},
		{Script: `2 * 1`, RunOutput: int64(2)},
		{Script: `2 / 1`, RunOutput: float64(2)},
		{Script: `2.1 + 1.1`, RunOutput: float64(3.2)},
		{Script: `2.1 - 1.1`, RunOutput: float64(1)},
		{Script: `2.1 * 2.0`, RunOutput: float64(4.2)},
		{Script: `6.5 / 2.0`, RunOutput: float64(3.25)},

		{Script: `a + b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(3)},
		{Script: `a - b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(1)},
		{Script: `a * b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(2)},
		{Script: `a / b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: float64(2)},
		{Script: `a + b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(3.2)},
		{Script: `a - b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(1)},
		{Script: `a * b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, RunOutput: float64(4.2)},
		{Script: `a / b`, Input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, RunOutput: float64(3.25)},

		{Script: `a + b`, Input: map[string]interface{}{"a": "a", "b": "b"}, RunOutput: "ab"},
		{Script: `a + b`, Input: map[string]interface{}{"a": "a", "b": int64(1)}, RunOutput: "a1"},
		{Script: `a + b`, Input: map[string]interface{}{"a": "a", "b": float64(1.1)}, RunOutput: "a1.1"},
		{Script: `a + z`, Input: map[string]interface{}{"a": "a"}, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},
		{Script: `z + b`, Input: map[string]interface{}{"a": "a"}, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},

		{Script: `c = a + b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(3), Output: map[string]interface{}{"c": int64(3)}},
		{Script: `c = a - b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(1), Output: map[string]interface{}{"c": int64(1)}},
		{Script: `c = a * b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"c": int64(2)}},
		{Script: `c = a / b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: float64(2), Output: map[string]interface{}{"c": float64(2)}},
		{Script: `c = a + b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(3.2), Output: map[string]interface{}{"c": float64(3.2)}},
		{Script: `c = a - b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(1.1)}, RunOutput: float64(1), Output: map[string]interface{}{"c": float64(1)}},
		{Script: `c = a * b`, Input: map[string]interface{}{"a": float64(2.1), "b": float64(2)}, RunOutput: float64(4.2), Output: map[string]interface{}{"c": float64(4.2)}},
		{Script: `c = a / b`, Input: map[string]interface{}{"a": float64(6.5), "b": float64(2)}, RunOutput: float64(3.25), Output: map[string]interface{}{"c": float64(3.25)}},

		{Script: `a = nil; a++`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = false; a++`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = true; a++`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; a++`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1.5; a++`, RunOutput: float64(2.5), Output: map[string]interface{}{"a": float64(2.5)}},
		{Script: `a = "1"; a++`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = "a"; a++`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a = nil; a--`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a = false; a--`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a = true; a--`, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a = 2; a--`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2.5; a--`, RunOutput: float64(1.5), Output: map[string]interface{}{"a": float64(1.5)}},
		{Script: `a = "2"; a--`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "a"; a--`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},

		{Script: `a++`, Input: map[string]interface{}{"a": nil}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a++`, Input: map[string]interface{}{"a": false}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a++`, Input: map[string]interface{}{"a": true}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a++`, Input: map[string]interface{}{"a": int32(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a++`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a++`, Input: map[string]interface{}{"a": float32(1.5)}, RunOutput: float64(2.5), Output: map[string]interface{}{"a": float64(2.5)}},
		{Script: `a++`, Input: map[string]interface{}{"a": float64(1.5)}, RunOutput: float64(2.5), Output: map[string]interface{}{"a": float64(2.5)}},
		{Script: `a++`, Input: map[string]interface{}{"a": "2"}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a++`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a--`, Input: map[string]interface{}{"a": nil}, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": false}, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": true}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a--`, Input: map[string]interface{}{"a": int32(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": float32(2.5)}, RunOutput: float64(1.5), Output: map[string]interface{}{"a": float64(1.5)}},
		{Script: `a--`, Input: map[string]interface{}{"a": float64(2.5)}, RunOutput: float64(1.5), Output: map[string]interface{}{"a": float64(1.5)}},
		{Script: `a--`, Input: map[string]interface{}{"a": "2"}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a--`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},

		{Script: `1++`, RunError: fmt.Errorf("invalid operation"), RunOutput: nil},
		{Script: `1--`, RunError: fmt.Errorf("invalid operation"), RunOutput: nil},
		{Script: `z++`, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},
		{Script: `z--`, RunError: fmt.Errorf("undefined symbol 'z'"), RunOutput: nil},
		{Script: `!(1++)`, RunError: fmt.Errorf("invalid operation"), RunOutput: nil},

		{Script: `a += 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a -= 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a *= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},
		{Script: `a /= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: float64(1), Output: map[string]interface{}{"a": float64(1)}},
		{Script: `a += 1`, Input: map[string]interface{}{"a": 2.1}, RunOutput: float64(3.1), Output: map[string]interface{}{"a": float64(3.1)}},
		{Script: `a -= 1`, Input: map[string]interface{}{"a": 2.1}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `a *= 2`, Input: map[string]interface{}{"a": 2.1}, RunOutput: float64(4.2), Output: map[string]interface{}{"a": float64(4.2)}},
		{Script: `a /= 2`, Input: map[string]interface{}{"a": 6.5}, RunOutput: float64(3.25), Output: map[string]interface{}{"a": float64(3.25)}},

		{Script: `a ** 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(8), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a ** 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: float64(8), Output: map[string]interface{}{"a": float64(2)}},

		{Script: `a &= 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a &= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a &= 1`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(0)}},
		{Script: `a &= 2`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a |= 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a |= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a |= 1`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a |= 2`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a << 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(8), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >> 2`, Input: map[string]interface{}{"a": int64(8)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(8)}},
		{Script: `a << 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: int64(8), Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >> 2`, Input: map[string]interface{}{"a": float64(8)}, RunOutput: int64(2), Output: map[string]interface{}{"a": float64(8)}},

		{Script: `a % 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(0), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a % 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a % 2`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(0), Output: map[string]interface{}{"a": float64(2.1)}},
		{Script: `a % 3`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": float64(2.1)}},

		{Script: `a * 4`, Input: map[string]interface{}{"a": "a"}, RunOutput: "aaaa", Output: map[string]interface{}{"a": "a"}},
		{Script: `a * 4.0`, Input: map[string]interface{}{"a": "a"}, RunOutput: float64(0), Output: map[string]interface{}{"a": "a"}},

		{Script: `-a`, Input: map[string]interface{}{"a": nil}, RunOutput: float64(-0), Output: map[string]interface{}{"a": nil}},
		{Script: `-a`, Input: map[string]interface{}{"a": "a"}, RunOutput: float64(-0), Output: map[string]interface{}{"a": "a"}},
		{Script: `-a`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(-2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `-a`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: float64(-2.1), Output: map[string]interface{}{"a": float64(2.1)}},

		{Script: `^a`, Input: map[string]interface{}{"a": nil}, RunOutput: int64(-1), Output: map[string]interface{}{"a": nil}},
		{Script: `^a`, Input: map[string]interface{}{"a": "a"}, RunOutput: int64(-1), Output: map[string]interface{}{"a": "a"}},
		{Script: `^a`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(-3), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `^a`, Input: map[string]interface{}{"a": float64(2.1)}, RunOutput: int64(-3), Output: map[string]interface{}{"a": float64(2.1)}},

		{Script: `!true`, RunOutput: false},
		{Script: `!1`, RunOutput: false},
	}
	testlib.Run(t, tests, nil)
}

func TestComparisonOperators(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `a == 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == 1.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 1.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a != 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a == 1`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 1`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 1.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 1.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a != 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},

		{Script: `a == nil`, Input: map[string]interface{}{"a": nil}, RunOutput: true, Output: map[string]interface{}{"a": nil}},
		{Script: `a == nil`, Input: map[string]interface{}{"a": nil}, RunOutput: true, Output: map[string]interface{}{"a": nil}},
		{Script: `a == nil`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a == nil`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a == 2`, Input: map[string]interface{}{"a": nil}, RunOutput: false, Output: map[string]interface{}{"a": nil}},
		{Script: `a == 2.0`, Input: map[string]interface{}{"a": nil}, RunOutput: false, Output: map[string]interface{}{"a": nil}},

		{Script: `1 == 1.0`, RunOutput: true},
		{Script: `1 != 1.0`, RunOutput: false},
		{Script: `"a" != "a"`, RunOutput: false},

		{Script: `a > 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a > 1`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a > 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a > 1.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a < 3.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a > 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a > 1`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a > 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a > 1.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a < 3.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},

		{Script: `a >= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >= 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 2`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 3`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >= 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a >= 3.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 2.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a <= 3.0`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2)}},

		{Script: `a >= 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >= 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 2`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 3`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >= 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a >= 3.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 2.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},
		{Script: `a <= 3.0`, Input: map[string]interface{}{"a": float64(2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2)}},

		{Script: `1 == 1 && 1  == 1`, RunOutput: true},
		{Script: `1 == 1 && 1  == 2`, RunOutput: false},
		{Script: `1 == 2 && 1  == 1`, RunOutput: false},
		{Script: `1 == 2 && 1  == 2`, RunOutput: false},
		{Script: `1 == 1 || 1  == 1`, RunOutput: true},
		{Script: `1 == 1 || 1  == 2`, RunOutput: true},
		{Script: `1 == 2 || 1  == 1`, RunOutput: true},
		{Script: `1 == 2 || 1  == 2`, RunOutput: false},

		{Script: `true && func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},
		{Script: `false && func(){throw('abcde')}()`, RunOutput: false},
		{Script: `true || func(){throw('abcde')}()`, RunOutput: true},
		{Script: `false || func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},
		{Script: `true && true && func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},
		{Script: `true && false && func(){throw('abcde')}()`, RunOutput: false},
		{Script: `true && func(){throw('abcde')}() && true`, RunError: fmt.Errorf("abcde")},
		{Script: `false && func(){throw('abcde')}() && func(){throw('abcde')}() `, RunOutput: false},

		{Script: `true && func(){throw('abcde')}() || false`, RunError: fmt.Errorf("abcde")},
		{Script: `true && false || func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},
		{Script: `true && true || func(){throw('abcde')}()`, RunOutput: true},

		{Script: `true || func(){throw('abcde')}() || func(){throw('abcde')}()`, RunOutput: true},
		{Script: `false || func(){throw('abcde')}() || true`, RunError: fmt.Errorf("abcde")},
		{Script: `false || true || func(){throw('abcde')}()`, RunOutput: true},
		{Script: `false || false || func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},

		{Script: `false || false && func(){throw('abcde')}()`, RunOutput: false},
		{Script: `false || true && func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},
		{Script: `false || func(){throw('abcde')}() || true`, RunError: fmt.Errorf("abcde")},

		{Script: `1 == 1 && func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},
		{Script: `1 == 2 && func(){throw('abcde')}()`, RunOutput: false},
		{Script: `1 == 1 || func(){throw('abcde')}()`, RunOutput: true},
		{Script: `1 == 2 || func(){throw('abcde')}()`, RunError: fmt.Errorf("abcde")},

		{Script: `(true || func(){throw('abcde')}()) && (true || func(){throw('hello')}())`, RunOutput: true},
		{Script: `(true || func(){throw('abcde')}()) && (true && func(){throw('hello')}())`, RunError: fmt.Errorf("hello")},
		{Script: `(true || func(){throw('abcde')}()) || (true && func(){throw('hello')}())`, RunOutput: true},
		{Script: `(true && func(){throw('abcde')}()) && (true && func(){throw('hello')}())`, RunError: fmt.Errorf("abcde")},
		{Script: `(true || func(){throw('abcde')}()) && (false || func(){throw('hello')}())`, RunError: fmt.Errorf("hello")},

		{Script: `true == "1"`, RunOutput: true},
		{Script: `true == "t"`, RunOutput: true},
		{Script: `true == "T"`, RunOutput: true},
		{Script: `true == "true"`, RunOutput: true},
		{Script: `true == "TRUE"`, RunOutput: true},
		{Script: `true == "True"`, RunOutput: true},
		{Script: `true == "false"`, RunOutput: false},
		{Script: `false == "0"`, RunOutput: true},
		{Script: `false == "f"`, RunOutput: true},
		{Script: `false == "F"`, RunOutput: true},
		{Script: `false == "false"`, RunOutput: true},
		{Script: `false == "false"`, RunOutput: true},
		{Script: `false == "FALSE"`, RunOutput: true},
		{Script: `false == "False"`, RunOutput: true},
		{Script: `false == "true"`, RunOutput: false},
		{Script: `false == "foo"`, RunOutput: false},
		{Script: `true == "foo"`, RunOutput: true},

		{Script: `0 == "0"`, RunOutput: true},
		{Script: `"1.0" == 1`, RunOutput: true},
		{Script: `1 == "1"`, RunOutput: true},
		{Script: `0.0 == "0"`, RunOutput: true},
		{Script: `0.0 == "0.0"`, RunOutput: true},
		{Script: `1.0 == "1.0"`, RunOutput: true},
		{Script: `1.2 == "1.2"`, RunOutput: true},
		{Script: `"7" == 7.2`, RunOutput: false},
		{Script: `1.2 == "1"`, RunOutput: false},
		{Script: `"1.1" == 1`, RunOutput: false},
		{Script: `0 == "1"`, RunOutput: false},

		{Script: `a == b`, Input: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}, RunOutput: true, Output: map[string]interface{}{"a": reflect.Value{}, "b": reflect.Value{}}},
		{Script: `a == b`, Input: map[string]interface{}{"a": reflect.Value{}, "b": true}, RunOutput: false, Output: map[string]interface{}{"a": reflect.Value{}, "b": true}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": reflect.Value{}}, RunOutput: false, Output: map[string]interface{}{"a": true, "b": reflect.Value{}}},

		{Script: `a == b`, Input: map[string]interface{}{"a": nil, "b": nil}, RunOutput: true, Output: map[string]interface{}{"a": nil, "b": nil}},
		{Script: `a == b`, Input: map[string]interface{}{"a": nil, "b": true}, RunOutput: false, Output: map[string]interface{}{"a": nil, "b": true}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": nil}, RunOutput: false, Output: map[string]interface{}{"a": true, "b": nil}},

		{Script: `a == b`, Input: map[string]interface{}{"a": false, "b": false}, RunOutput: true, Output: map[string]interface{}{"a": false, "b": false}},
		{Script: `a == b`, Input: map[string]interface{}{"a": false, "b": true}, RunOutput: false, Output: map[string]interface{}{"a": false, "b": true}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": false}, RunOutput: false, Output: map[string]interface{}{"a": true, "b": false}},
		{Script: `a == b`, Input: map[string]interface{}{"a": true, "b": true}, RunOutput: true, Output: map[string]interface{}{"a": true, "b": true}},

		{Script: `a == b`, Input: map[string]interface{}{"a": int32(1), "b": int32(1)}, RunOutput: true, Output: map[string]interface{}{"a": int32(1), "b": int32(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int32(1), "b": int32(2)}, RunOutput: false, Output: map[string]interface{}{"a": int32(1), "b": int32(2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int32(2), "b": int32(1)}, RunOutput: false, Output: map[string]interface{}{"a": int32(2), "b": int32(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int32(2), "b": int32(2)}, RunOutput: true, Output: map[string]interface{}{"a": int32(2), "b": int32(2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": int64(1), "b": int64(1)}, RunOutput: true, Output: map[string]interface{}{"a": int64(1), "b": int64(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int64(1), "b": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int64(2), "b": int64(1)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2), "b": int64(1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": int64(2), "b": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": int64(2), "b": int64(2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}, RunOutput: true, Output: map[string]interface{}{"a": float32(1.1), "b": float32(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}, RunOutput: false, Output: map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}, RunOutput: false, Output: map[string]interface{}{"a": float32(2.2), "b": float32(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}, RunOutput: true, Output: map[string]interface{}{"a": float32(2.2), "b": float32(2.2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}, RunOutput: true, Output: map[string]interface{}{"a": float64(1.1), "b": float64(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}, RunOutput: false, Output: map[string]interface{}{"a": float64(1.1), "b": float64(2.2)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}, RunOutput: false, Output: map[string]interface{}{"a": float64(2.2), "b": float64(1.1)}},
		{Script: `a == b`, Input: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}, RunOutput: true, Output: map[string]interface{}{"a": float64(2.2), "b": float64(2.2)}},

		{Script: `a == b`, Input: map[string]interface{}{"a": 'a', "b": 'a'}, RunOutput: true, Output: map[string]interface{}{"a": 'a', "b": 'a'}},
		{Script: `a == b`, Input: map[string]interface{}{"a": 'a', "b": 'b'}, RunOutput: false, Output: map[string]interface{}{"a": 'a', "b": 'b'}},
		{Script: `a == b`, Input: map[string]interface{}{"a": 'b', "b": 'a'}, RunOutput: false, Output: map[string]interface{}{"a": 'b', "b": 'a'}},
		{Script: `a == b`, Input: map[string]interface{}{"a": 'b', "b": 'b'}, RunOutput: true, Output: map[string]interface{}{"a": 'b', "b": 'b'}},

		{Script: `a == b`, Input: map[string]interface{}{"a": "a", "b": "a"}, RunOutput: true, Output: map[string]interface{}{"a": "a", "b": "a"}},
		{Script: `a == b`, Input: map[string]interface{}{"a": "a", "b": "b"}, RunOutput: false, Output: map[string]interface{}{"a": "a", "b": "b"}},
		{Script: `a == b`, Input: map[string]interface{}{"a": "b", "b": "a"}, RunOutput: false, Output: map[string]interface{}{"a": "b", "b": "a"}},
		{Script: `a == b`, Input: map[string]interface{}{"a": "b", "b": "b"}, RunOutput: true, Output: map[string]interface{}{"a": "b", "b": "b"}},

		{Script: `b = "\"a\""; a == b`, Input: map[string]interface{}{"a": `"a"`}, RunOutput: true, Output: map[string]interface{}{"a": `"a"`, "b": `"a"`}},

		{Script: `a = "test"; a == "test"`, RunOutput: true},
		{Script: `a = "test"; a[0:1] == "t"`, RunOutput: true},
		{Script: `a = "test"; a[0:2] == "te"`, RunOutput: true},
		{Script: `a = "test"; a[1:3] == "es"`, RunOutput: true},
		{Script: `a = "test"; a[0:4] == "test"`, RunOutput: true},

		{Script: `a = "a b"; a[1] == ' '`, RunOutput: true},
		{Script: `a = "test"; a[0] == 't'`, RunOutput: true},
		{Script: `a = "test"; a[1] == 'e'`, RunOutput: true},
		{Script: `a = "test"; a[3] == 't'`, RunOutput: true},

		{Script: `a = "a b"; a[1] != ' '`, RunOutput: false},
		{Script: `a = "test"; a[0] != 't'`, RunOutput: false},
		{Script: `a = "test"; a[1] != 'e'`, RunOutput: false},
		{Script: `a = "test"; a[3] != 't'`, RunOutput: false},
	}
	testlib.Run(t, tests, nil)
}

func TestTernaryOperator(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `a = 1 ? 2 : panic(2)`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = c ? a : b`, RunError: fmt.Errorf("undefined symbol 'c'")},
		{Script: `a = a ? a : b`, RunError: fmt.Errorf("undefined symbol 'a'")},
		{Script: `a = 0; a = a ? a : b`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `a = -1 ? 2 : 1`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = true ? 2 : 1`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = false ? 2 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "true" ? 2 : 1`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = "false" ? 2 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "-1" ? 2 : 1`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = "0" ? 2 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "0.0" ? 2 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "2" ? 2 : 1`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": int64(0)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": float64(0.0)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": float64(2.0)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b ? 2 : 1.0`, Input: map[string]interface{}{"b": float64(0.0)}, RunOutput: float64(1.0), Output: map[string]interface{}{"a": float64(1.0)}},
		{Script: `a = b ? 2 : 1.0`, Input: map[string]interface{}{"b": float64(0.1)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b ? 2 : 1.0`, Input: map[string]interface{}{"b": nil}, RunOutput: float64(1.0), Output: map[string]interface{}{"a": float64(1.0)}},
		{Script: `a = nil ? 2 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": []interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": map[string]interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b[1] ? 2 : 1`, Input: map[string]interface{}{"b": []interface{}{}}, RunError: fmt.Errorf("index out of range")},
		{Script: `a = b[1][2] ? 2 : 1`, Input: map[string]interface{}{"b": []interface{}{}}, RunError: fmt.Errorf("index out of range")},
		{Script: `a = [] ? 2 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = [2] ? 2 : 1`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b ? 2 : 1`, Input: map[string]interface{}{"b": map[string]interface{}{"test": int64(2)}}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b["test"] ? 2 : 1`, Input: map[string]interface{}{"b": map[string]interface{}{"test": int64(2)}}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b["test"][1] ? 2 : 1`, Input: map[string]interface{}{"b": map[string]interface{}{"test": 2}}, RunError: fmt.Errorf("type int does not support index operation")},
		{Script: `b = "test"; a = b ? 2 : "empty"`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `b = "test"; a = b[1:3] ? 2 : "empty"`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `b = "test"; a = b[2:2] ? 2 : "empty"`, RunOutput: "empty", Output: map[string]interface{}{"a": "empty"}},
		{Script: `b = "0.0"; a = false ? 2 : b ? 3 : 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `b = "true"; a = false ? 2 : b ? 3 : 1`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
	}
	testlib.Run(t, tests, nil)
}

func TestNilCoalescingOperator(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `a = 1 ?? panic(2)`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = c ?? b`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `a = -1 ?? 1`, RunOutput: int64(-1), Output: map[string]interface{}{"a": int64(-1)}},
		{Script: `a = true ?? 1`, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `a = false ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "true" ?? 1`, RunOutput: "true", Output: map[string]interface{}{"a": "true"}},
		{Script: `a = "false" ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "-1" ?? 1`, RunOutput: "-1", Output: map[string]interface{}{"a": "-1"}},
		{Script: `a = "0" ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "0.0" ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = "2" ?? 1`, RunOutput: "2", Output: map[string]interface{}{"a": "2"}},
		{Script: `a = b ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": int64(0)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": int64(2)}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": float64(0.0)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": float64(2.0)}, RunOutput: float64(2.0), Output: map[string]interface{}{"a": float64(2.0)}},
		{Script: `a = b ?? 1.0`, Input: map[string]interface{}{"b": float64(0.0)}, RunOutput: float64(1.0), Output: map[string]interface{}{"a": float64(1.0)}},
		{Script: `a = b ?? 1.0`, Input: map[string]interface{}{"b": float64(0.1)}, RunOutput: float64(0.1), Output: map[string]interface{}{"a": float64(0.1)}},
		{Script: `a = b ?? 1.0`, Input: map[string]interface{}{"b": nil}, RunOutput: float64(1.0), Output: map[string]interface{}{"a": float64(1.0)}},
		{Script: `a = nil ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": []interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": map[string]interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b[1] ?? 1`, Input: map[string]interface{}{"b": []interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = b[1][2] ?? 1`, Input: map[string]interface{}{"b": []interface{}{}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = [] ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = [2] ?? 1`, RunOutput: []interface{}{int64(2)}, Output: map[string]interface{}{"a": []interface{}{int64(2)}}},
		{Script: `a = b ?? 1`, Input: map[string]interface{}{"b": map[string]interface{}{"test": int64(2)}}, RunOutput: map[string]interface{}{"test": int64(2)}, Output: map[string]interface{}{"a": map[string]interface{}{"test": int64(2)}}},
		{Script: `a = b["test"] ?? 1`, Input: map[string]interface{}{"b": map[string]interface{}{"test": int64(2)}}, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = b["test"][1] ?? 1`, Input: map[string]interface{}{"b": map[string]interface{}{"test": 2}}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `b = "test"; a = b ?? "empty"`, RunOutput: "test", Output: map[string]interface{}{"a": "test"}},
		{Script: `b = "test"; a = b[1:3] ?? "empty"`, RunOutput: "es", Output: map[string]interface{}{"a": "es"}},
		{Script: `b = "test"; a = b[2:2] ?? "empty"`, RunOutput: "empty", Output: map[string]interface{}{"a": "empty"}},
		{Script: `b = "0.0"; a = false ?? b ?? 1`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
	}
	testlib.Run(t, tests, nil)
}

func TestIf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `if 1++ {}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `if false {} else if 1++ {}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `if false {} else if true { 1++ }`, RunError: fmt.Errorf("invalid operation")},

		{Script: `if true {}`, Input: map[string]interface{}{"a": nil}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `if true {}`, Input: map[string]interface{}{"a": true}, RunOutput: nil, Output: map[string]interface{}{"a": true}},
		{Script: `if true {}`, Input: map[string]interface{}{"a": int64(1)}, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `if true {}`, Input: map[string]interface{}{"a": float64(1.1)}, RunOutput: nil, Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `if true {}`, Input: map[string]interface{}{"a": "a"}, RunOutput: nil, Output: map[string]interface{}{"a": "a"}},

		{Script: `if true {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `if true {a = true}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: true, Output: map[string]interface{}{"a": true}},
		{Script: `if true {a = 1}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `if true {a = 1.1}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: float64(1.1), Output: map[string]interface{}{"a": float64(1.1)}},
		{Script: `if true {a = "a"}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: "a", Output: map[string]interface{}{"a": "a"}},

		{Script: `if a == 1 {a = 1}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `if a == 2 {a = 1}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `if a == 1 {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `if a == 2 {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},

		{Script: `if a == 1 {a = 1} else {a = 3}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `if a == 2 {a = 1} else {a = 3}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `if a == 1 {a = 1} else if a == 2 {a = 3}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3} else {a = 4}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},

		{Script: `if a == 1 {a = 1} else {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `if a == 2 {a = nil} else {a = 3}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `if a == 1 {a = nil} else if a == 3 {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: false, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `if a == 1 {a = 1} else if a == 2 {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3} else {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},

		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = 5}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(5)}},
		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = nil}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},
		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = 4} else {a = 5}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},
		{Script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = nil} else {a = 5}`, Input: map[string]interface{}{"a": int64(2)}, RunOutput: nil, Output: map[string]interface{}{"a": nil}},

		// check scope
		{Script: `a = 1; if a == 1 { b = 2 }; b`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; if a == 2 { b = 3 } else { b = 4 }; b`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; if a == 2 { b = 3 } else if a == 1 { b = 4 }; b`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; if a == 2 { b = 4 } else if a == 5 { b = 6 } else if a == 1 { c = b }`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; if a == 2 { b = 4 } else if a == 5 { b = 6 } else if a == 1 { b = 7 }; b`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": int64(1)}},
	}
	testlib.Run(t, tests, nil)
}

func TestSwitch(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		// test parse errors
		{Script: `switch {}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a = 1; switch a; {}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a = 1; switch a = 2 {}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a = 1; switch a {default: return 6; default: return 7}`, ParseError: fmt.Errorf("multiple default statement"), RunOutput: int64(7)},
		{Script: `a = 1; switch a {case 1: return 5; default: return 6; default: return 7}`, ParseError: fmt.Errorf("multiple default statement"), RunOutput: int64(5)},

		// test run errors
		{Script: `a = 1; switch 1++ {}`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a = 1; switch a {case 1++: return 2}`, RunError: fmt.Errorf("invalid operation")},

		// test no or empty cases
		{Script: `a = 1; switch a {}`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; switch a {case: return 2}`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; switch a {case: return 2; case: return 3}`, RunOutput: int64(1), Output: map[string]interface{}{"a": int64(1)}},

		// test 1 case
		{Script: `a = 1; switch a {case 1: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: return 5}`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; switch a {case 1,2: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1,2: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1,2: return 5}`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; switch a {case 1,2,3: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1,2,3: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1,2,3: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 4; switch a {case 1,2,3: return 5}`, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},
		{Script: `a = func() { return 1 }; switch a() {case 1: return 5}`, RunOutput: int64(5)},
		{Script: `a = func() { return 2 }; switch a() {case 1: return 5}`, RunOutput: int64(2)},
		{Script: `a = func() { return 5 }; b = 1; switch b {case 1: return a() }`, RunOutput: int64(5), Output: map[string]interface{}{"b": int64(1)}},
		{Script: `a = func() { return 6 }; b = 2; switch b {case 1: return a() }`, RunOutput: int64(2), Output: map[string]interface{}{"b": int64(2)}},

		// test 2 cases
		{Script: `a = 1; switch a {case 1: return 5; case 2: return 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: return 5; case 2: return 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1: return 5; case 2: return 6}`, RunOutput: int64(3), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; switch a {case 1: return 5; case 2,3: return 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: return 5; case 2,3: return 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1: return 5; case 2,3: return 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 4; switch a {case 1: return 5; case 2,3: return 6}`, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},

		// test 3 cases
		{Script: `a = 1; switch a {case 1,2: return 5; case 3: return 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1,2: return 5; case 3: return 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1,2: return 5; case 3: return 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 4; switch a {case 1,2: return 5; case 3: return 6}`, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},
		{Script: `a = 1; switch a {case 1,2: return 5; case 2,3: return 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1,2: return 5; case 2,3: return 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1,2: return 5; case 2,3: return 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 4; switch a {case 1,2: return 5; case 2,3: return 6}`, RunOutput: int64(4), Output: map[string]interface{}{"a": int64(4)}},

		// test default
		{Script: `a = 1; switch a {default: return 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; switch a {case 1: return 5; case 2: return 6; default: return 7}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: return 5; case 2: return 6; default: return 7}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1: return 5; case 2: return 6; default: return 7}`, RunOutput: int64(7), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; switch a {case 1: return 5; case 2,3: return 6; default: return 7}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: return 5; case 2,3: return 6; default: return 7}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1: return 5; case 2,3: return 6; default: return 7}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 4; switch a {case 1: return 5; case 2,3: return 6; default: return 7}`, RunOutput: int64(7), Output: map[string]interface{}{"a": int64(4)}},
		{Script: `a = 1; switch a {case 1,2: return 5; case 3: return 6; default: return 7}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1,2: return 5; case 3: return 6; default: return 7}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 3; switch a {case 1,2: return 5; case 3: return 6; default: return 7}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 4; switch a {case 1,2: return 5; case 3: return 6; default: return 7}`, RunOutput: int64(7), Output: map[string]interface{}{"a": int64(4)}},

		// test scope
		{Script: `a = 1; switch a {case 1: a = 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(5)}},
		{Script: `a = 2; switch a {case 1: a = 5}`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; b = 5; switch a {case 1: b = 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(1), "b": int64(6)}},
		{Script: `a = 2; b = 5; switch a {case 1: b = 6}`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2), "b": int64(5)}},
		{Script: `a = 1; b = 5; switch a {case 1: b = 6; default: b = 7}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(1), "b": int64(6)}},
		{Script: `a = 2; b = 5; switch a {case 1: b = 6; default: b = 7}`, RunOutput: int64(7), Output: map[string]interface{}{"a": int64(2), "b": int64(7)}},

		// test scope without define b
		{Script: `a = 1; switch a {case 1: b = 5}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: b = 5}`, RunOutput: int64(2), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; switch a {case 1: b = 5}; b`, RunError: fmt.Errorf("undefined symbol 'b'"), RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: b = 5}; b`, RunError: fmt.Errorf("undefined symbol 'b'"), RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; switch a {case 1: b = 5; default: b = 6}`, RunOutput: int64(5), Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: b = 5; default: b = 6}`, RunOutput: int64(6), Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; switch a {case 1: b = 5; default: b = 6}; b`, RunError: fmt.Errorf("undefined symbol 'b'"), RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 2; switch a {case 1: b = 5; default: b = 6}; b`, RunError: fmt.Errorf("undefined symbol 'b'"), RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},

		// test new lines
		{Script: `
a = 1;
switch a {
	case 1:
		return 1
}`, RunOutput: int64(1)},
	}
	testlib.Run(t, tests, nil)
}

func TestForLoop(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `break`, RunError: fmt.Errorf("unexpected break statement")},
		{Script: `continue`, RunError: fmt.Errorf("unexpected continue statement")},
		{Script: `for 1++ { }`, RunError: fmt.Errorf("invalid operation")},
		{Script: `for { 1++ }`, RunError: fmt.Errorf("invalid operation")},
		{Script: `for a in 1++ { }`, RunError: fmt.Errorf("invalid operation")},

		{Script: `for { break }`, RunOutput: nil},
		{Script: `for {a = 1; if a == 1 { break } }`, RunOutput: nil},
		{Script: `a = 1; for { if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for { if a == 1 { break }; a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},

		{Script: `a = 1; for { a++; if a == 2 { continue } else { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; for { a++; if a == 2 { continue }; if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `for a in [1] { if a == 1 { break } }`, RunOutput: nil},
		{Script: `for a in [1, 2] { if a == 2 { break } }`, RunOutput: nil},
		{Script: `for a in [1, 2, 3] { if a == 3 { break } }`, RunOutput: nil},

		{Script: `a = [1]; for b in a { if b == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1)}}},
		{Script: `a = [1, 2]; for b in a { if b == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}}},
		{Script: `a = [1, 2, 3]; for b in a { if b == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}}},

		{Script: `a = [1]; b = 0; for c in a { b = c }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{Script: `a = [1, 2]; b = 0; for c in a { b = c }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2)}, "b": int64(2)}},
		{Script: `a = [1, 2, 3]; b = 0; for c in a { b = c }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1), int64(2), int64(3)}, "b": int64(3)}},

		{Script: `a = 1; for a < 2 { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for a < 3 { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = 1; for nil { a++; if a > 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for nil { a++; if a > 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for true { a++; if a > 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},
		{Script: `a = 1; for true { a++; if a > 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(4)}},

		{Script: `func x() { return [1] }; for b in x() { if b == 1 { break } }`, RunOutput: nil},
		{Script: `func x() { return [1, 2] }; for b in x() { if b == 2 { break } }`, RunOutput: nil},
		{Script: `func x() { return [1, 2, 3] }; for b in x() { if b == 3 { break } }`, RunOutput: nil},

		{Script: `func x() { a = 1; for { if a == 1 { return } } }; x()`, RunOutput: nil},
		{Script: `func x() { a = 1; for { if a == 1 { return nil } } }; x()`, RunOutput: nil},
		{Script: `func x() { a = 1; for { if a == 1 { return true } } }; x()`, RunOutput: true},
		{Script: `func x() { a = 1; for { if a == 1 { return 1 } } }; x()`, RunOutput: int64(1)},
		{Script: `func x() { a = 1; for { if a == 1 { return 1.1 } } }; x()`, RunOutput: float64(1.1)},
		{Script: `func x() { a = 1; for { if a == 1 { return "a" } } }; x()`, RunOutput: "a"},

		{Script: `func x() { for a in [1, 2, 3] { if a == 3 { return } } }; x()`, RunOutput: nil},
		{Script: `func x() { for a in [1, 2, 3] { if a == 1 { continue } } }; x()`, RunOutput: nil},
		{Script: `func x() { for a in [1, 2, 3] { if a == 1 { continue };  if a == 3 { return } } }; x()`, RunOutput: nil},

		{Script: `func x() { return [1, 2] }; a = 1; for i in x() { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},
		{Script: `func x() { return [1, 2, 3] }; a = 1; for i in x() { a++ }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(4)}},

		{Script: `for a = 1; nil; nil { return }`},
		// TOFIX:
		// {Script: `for a, b = 1; nil; nil { return }`},
		// {Script: `for a, b = 1, 2; nil; nil { return }`},

		{Script: `for var a = 1; nil; nil { return }`},
		{Script: `for var a = 1, 2; nil; nil { return }`},
		{Script: `for var a, b = 1; nil; nil { return }`},
		{Script: `for var a, b = 1, 2; nil; nil { return }`},

		{Script: `for a.b = 1; nil; nil { return }`, RunError: fmt.Errorf("undefined symbol 'a'")},

		{Script: `for a = 1; nil; nil { if a == 1 { break } }`, RunOutput: nil},
		{Script: `for a = 1; nil; nil { if a == 2 { break }; a++ }`, RunOutput: nil},
		{Script: `for a = 1; nil; nil { a++; if a == 3 { break } }`, RunOutput: nil},

		{Script: `for a = 1; a < 1; nil { }`, RunOutput: nil},
		{Script: `for a = 1; a > 1; nil { }`, RunOutput: nil},
		{Script: `for a = 1; a == 1; nil { break }`, RunOutput: nil},

		{Script: `for a = 1; a == 1; a++ { }`, RunOutput: nil},
		{Script: `for a = 1; a < 2; a++ { }`, RunOutput: nil},
		{Script: `for a = 1; a < 3; a++ { }`, RunOutput: nil},

		{Script: `a = 1; for b = 1; a < 1; a++ { }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ { }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ { }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 1 { continue } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 1 { continue } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 1 { continue } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 1 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 2 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 1; a++ {  if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(1)}},
		{Script: `a = 1; for b = 1; a < 2; a++ {  if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(2)}},
		{Script: `a = 1; for b = 1; a < 3; a++ {  if a == 3 { break } }`, RunOutput: nil, Output: map[string]interface{}{"a": int64(3)}},

		{Script: `a = ["123", "456", "789"]; b = ""; for i = 0; i < len(a); i++ { b += a[i][len(a[i]) - 2:]; b += a[i][:len(a[i]) - 2] }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{"123", "456", "789"}, "b": "231564897"}},
		{Script: `a = [[["123"], ["456"]], [["789"]]]; b = ""; for i = 0; i < len(a); i++ { for j = 0; j < len(a[i]); j++ {  for k = 0; k < len(a[i][j]); k++ { for l = 0; l < len(a[i][j][k]); l++ { b += a[i][j][k][l] + "-" } } } }`,
			RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{[]interface{}{[]interface{}{"123"}, []interface{}{"456"}}, []interface{}{[]interface{}{"789"}}}, "b": "1-2-3-4-5-6-7-8-9-"}},

		{Script: `func x() { for a = 1; a < 3; a++ { if a == 1 { return a } } }; x()`, RunOutput: int64(1)},
		{Script: `func x() { for a = 1; a < 3; a++ { if a == 2 { return a } } }; x()`, RunOutput: int64(2)},
		{Script: `func x() { for a = 1; a < 3; a++ { if a == 3 { return a } } }; x()`, RunOutput: nil},
		{Script: `func x() { for a = 1; a < 3; a++ { if a == 4 { return a } } }; x()`, RunOutput: nil},

		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 1 { continue } }; return a }; x()`, RunOutput: int64(3)},
		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 2 { continue } }; return a }; x()`, RunOutput: int64(3)},
		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 3 { continue } }; return a }; x()`, RunOutput: int64(3)},
		{Script: `func x() { a = 1; for b = 1; a < 3; a++ { if a == 4 { continue } }; return a }; x()`, RunOutput: int64(3)},

		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{reflect.Value{}}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{reflect.Value{}}, "b": reflect.Value{}}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{nil}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{nil}, "b": nil}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{true}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{true}, "b": true}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{int32(1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int32(1)}, "b": int32(1)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{int64(1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{int64(1)}, "b": int64(1)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{float32(1.1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{float32(1.1)}, "b": float32(1.1)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{float64(1.1)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{float64(1.1)}, "b": float64(1.1)}},

		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(reflect.Value{})}, "b": interface{}(reflect.Value{})}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(nil)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(nil)}, "b": interface{}(nil)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(true)}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(true)}, "b": interface{}(true)}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(int32(1))}, "b": interface{}(int32(1))}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(int64(1))}, "b": interface{}(int64(1))}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(float32(1.1))}, "b": interface{}(float32(1.1))}},
		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}}, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{interface{}(float64(1.1))}, "b": interface{}(float64(1.1))}},

		{Script: `b = 0; for i in a { b = i }`, Input: map[string]interface{}{"a": interface{}([]interface{}{nil})}, RunOutput: nil, Output: map[string]interface{}{"a": interface{}([]interface{}{nil}), "b": nil}},

		{Script: `for i in nil { }`, RunError: fmt.Errorf("for cannot loop over type interface")},
		{Script: `for i in true { }`, RunError: fmt.Errorf("for cannot loop over type bool")},
		{Script: `for i in a { }`, Input: map[string]interface{}{"a": reflect.Value{}}, RunError: fmt.Errorf("for cannot loop over type struct"), Output: map[string]interface{}{"a": reflect.Value{}}},
		{Script: `for i in a { }`, Input: map[string]interface{}{"a": interface{}(nil)}, RunError: fmt.Errorf("for cannot loop over type interface"), Output: map[string]interface{}{"a": interface{}(nil)}},
		{Script: `for i in a { }`, Input: map[string]interface{}{"a": interface{}(true)}, RunError: fmt.Errorf("for cannot loop over type bool"), Output: map[string]interface{}{"a": interface{}(true)}},
		{Script: `for i in [1, 2, 3] { b++ }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = 1; a < 3; a++ { b++ }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = b; a < 3; a++ { }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = 1; b < 3; a++ { }`, RunError: fmt.Errorf("undefined symbol 'b'")},
		{Script: `for a = 1; a < 3; b++ { }`, RunError: fmt.Errorf("undefined symbol 'b'")},

		{Script: `a = 1; b = [{"c": "c"}]; for i in b { a = i }`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"c": "c"}, "b": []interface{}{map[interface{}]interface{}{"c": "c"}}}},
		{Script: `a = 1; b = {"x": [{"y": "y"}]};  for i in b.x { a = i }`, RunOutput: nil, Output: map[string]interface{}{"a": map[interface{}]interface{}{"y": "y"}, "b": map[interface{}]interface{}{"x": []interface{}{map[interface{}]interface{}{"y": "y"}}}}},

		{Script: `a = {}; b = 1; for i in a { b = i }; b`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{}, "b": int64(1)}},
		{Script: `a = {"x": 2}; b = 1; for i in a { b = i }; b`, RunOutput: "x", Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2)}, "b": "x"}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { b++ }; b`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(2)}},
		{Script: `a = {"x": 2, "y": 3}; for i in a { b++ }`, RunError: fmt.Errorf("undefined symbol 'b'"), Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}}},

		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "x" { continue }; b = i }; b`, RunOutput: "y", Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}, "b": "y"}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "y" { continue }; b = i }; b`, RunOutput: "x", Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}, "b": "x"}},
		{Script: `a = {"x": 2, "y": 3}; for i in a { if i ==  "x" { return 1 } }`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; for i in a { if i ==  "y" { return 2 } }`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "x" { break }; b++ }; if b > 1 { return false } else { return true }`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; b = 0; for i in a { if i ==  "y" { break }; b++ }; if b > 1 { return false } else { return true }`, RunOutput: true, Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}}},
		{Script: `a = {"x": 2, "y": 3}; b = 1; for i in a { if (i ==  "x" || i ==  "y") { break }; b++ }; b`, RunOutput: int64(1), Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2), "y": int64(3)}, "b": int64(1)}},

		{Script: `a = ["123", "456", "789"]; b = ""; for v in a { b += v[len(v) - 2:]; b += v[:len(v) - 2] }`, RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{"123", "456", "789"}, "b": "231564897"}},
		{Script: `a = [[["123"], ["456"]], [["789"]]]; b = ""; for x in a { for y in x  {  for z in y { for i = 0; i < len(z); i++ { b += z[i] + "-" } } } }`,
			RunOutput: nil, Output: map[string]interface{}{"a": []interface{}{[]interface{}{[]interface{}{"123"}, []interface{}{"456"}}, []interface{}{[]interface{}{"789"}}}, "b": "1-2-3-4-5-6-7-8-9-"}},

		{Script: `a = {"x": 2}; b = 0; for k, v in a { b = k }; b`, RunOutput: "x", Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2)}, "b": "x"}},
		{Script: `a = {"x": 2}; b = 0; for k, v in a { b = v }; b`, RunOutput: int64(2), Output: map[string]interface{}{"a": map[interface{}]interface{}{"x": int64(2)}, "b": int64(2)}},

		{Script: `a = make(chan int64, 1); a <- 1; v = 0; for val in a { v = val; break; }; v`, RunOutput: int64(1), Output: map[string]interface{}{"v": int64(1)}},
	}
	testlib.Run(t, tests, nil)
}

func TestItemInList(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `"a" in ["a"]`, RunOutput: true},
		{Script: `"a" in ["b"]`, RunOutput: false},
		{Script: `"a" in ["c", "b", "a"]`, RunOutput: true},
		{Script: `"a" in ["a", "b", 1]`, RunOutput: true},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"a"}}, RunOutput: true},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"b"}}, RunOutput: false},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"c", "b", "a"}}, RunOutput: true},
		{Script: `"a" in l`, Input: map[string]interface{}{"l": []interface{}{"a", "b", 1}}, RunOutput: true},

		{Script: `1 in [1]`, RunOutput: true},
		{Script: `1 in [2]`, RunOutput: false},
		{Script: `1 in [3, 2, 1]`, RunOutput: true},
		{Script: `1 in ["1"]`, RunOutput: true},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `"1" in l`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{2}}, RunOutput: false},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{3, 2, 1}}, RunOutput: true},
		{Script: `1 in l`, Input: map[string]interface{}{"l": []interface{}{"1"}}, RunOutput: true},

		{Script: `0.9 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1.0 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1.1 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1 in [0.9]`, Input: map[string]interface{}{"l": []interface{}{0.9}}, RunOutput: false},
		{Script: `1 in [1.0]`, Input: map[string]interface{}{"l": []interface{}{1.0}}, RunOutput: true},
		{Script: `1 in [1.1]`, Input: map[string]interface{}{"l": []interface{}{1.1}}, RunOutput: false},
		{Script: `0.9 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1.0 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1.1 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1 in [0.9]`, Input: map[string]interface{}{"l": []interface{}{0.9}}, RunOutput: false},
		{Script: `1 in [1.0]`, Input: map[string]interface{}{"l": []interface{}{1.0}}, RunOutput: true},
		{Script: `1 in [1.1]`, Input: map[string]interface{}{"l": []interface{}{1.1}}, RunOutput: false},

		{Script: `true in ["true"]`, RunOutput: true},
		{Script: `true in [true]`, RunOutput: true},
		{Script: `true in [true, false]`, RunOutput: true},
		{Script: `false in [false, true]`, RunOutput: true},
		{Script: `true in l`, Input: map[string]interface{}{"l": []interface{}{"true"}}, RunOutput: true},
		{Script: `true in l`, Input: map[string]interface{}{"l": []interface{}{true}}, RunOutput: true},
		{Script: `true in l`, Input: map[string]interface{}{"l": []interface{}{true, false}}, RunOutput: true},
		{Script: `false in l`, Input: map[string]interface{}{"l": []interface{}{false, true}}, RunOutput: true},

		{Script: `"a" in ["b", "a", "c"][1:]`, RunOutput: true},
		{Script: `"a" in ["b", "a", "c"][:1]`, RunOutput: false},
		{Script: `"a" in ["b", "a", "c"][1:2]`, RunOutput: true},
		{Script: `l = ["b", "a", "c"];"a" in l[1:]`, RunOutput: true},
		{Script: `l = ["b", "a", "c"];"a" in l[:1]`, RunOutput: false},
		{Script: `l = ["b", "a", "c"];"a" in l[1:2]`, RunOutput: true},
		{Script: `"a" in l[1:]`, Input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, RunOutput: true},
		{Script: `"a" in l[:1]`, Input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, RunOutput: false},
		{Script: `"a" in l[1:2]`, Input: map[string]interface{}{"l": []interface{}{"b", "a", "c"}}, RunOutput: true},

		// for i in list && item in list
		{Script: `list_of_list = [["a"]];for l in list_of_list { return "a" in l }`, RunOutput: true},
		{Script: `for l in list_of_list { return "a" in l }`, Input: map[string]interface{}{"list_of_list": []interface{}{[]interface{}{"a"}}}, RunOutput: true},

		// not slice or array
		// todo: support `"a" in "aaa"` ?
		{Script: `"a" in "aaa"`, RunError: fmt.Errorf("second argument must be slice or array; but have string")},
		{Script: `1 in 12345`, RunError: fmt.Errorf("type int64 does not support slice operation")},

		// a in item in list
		{Script: `"a" in 5 in [1, 2, 3]`, RunError: fmt.Errorf("type bool does not support slice operation")},

		// applying a in b in several part of expresstion/statement
		{Script: `switch 1 in [1] {case true: return true;default: return false}`, RunOutput: true},
		{Script: `switch 1 in [2,3] {case true: return true;default: return false}`, RunOutput: false},
		{Script: `switch true {case 1 in [1]: return true;default: return false}`, RunOutput: true},
		{Script: `switch false {case 1 in [1]: return true;default: return false}`, RunOutput: false},
		{Script: `if 1 in [1] {return true} else {return false}`, RunOutput: true},
		{Script: `if 1 in [2,3] {return true} else {return false}`, RunOutput: false},
		{Script: `for i in [1,2,3] { i++ }`},
		{Script: `a=1; a=a in [1]`, RunOutput: true},
		{Script: `a=1; a=a in [2,3]`, RunOutput: false},
		{Script: `1 in [1] && true`, RunOutput: true},
		{Script: `1 in [1] && false`, RunOutput: false},
		{Script: `1 in [1] || true`, RunOutput: true},
		{Script: `1 in [1] || false`, RunOutput: true},
		{Script: `1 in [2,3] && true`, RunOutput: false},
		{Script: `1 in [2,3] && false`, RunOutput: false},
		{Script: `1 in [2,3] || true`, RunOutput: true},
		{Script: `1 in [2,3] || false`, RunOutput: false},
		{Script: `1++ in [1, 2, 3]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `3++ in [1, 2, 3]`, RunError: fmt.Errorf("invalid operation")},
		{Script: `1 in 1++`, RunError: fmt.Errorf("invalid operation")},
		{Script: `a=1;a++ in [1, 2, 3]`, RunOutput: true},
		{Script: `a=3;a++ in [1, 2, 3]`, RunOutput: false},
		{Script: `switch 1 in l {case true: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `switch 1 in l {case true: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `switch true {case 1 in l: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `switch false {case 1 in l: return true;default: return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `if 1 in l {return true} else {return false}`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `if 1 in l {return true} else {return false}`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `for i in l { i++ }`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}},
		{Script: `a=1; a=a in l`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `a=1; a=a in l`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1 in l && true`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l && false`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `1 in l || true`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l || false`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in l && true`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1 in l && false`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1 in l || true`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: true},
		{Script: `1 in l || false`, Input: map[string]interface{}{"l": []interface{}{2, 3}}, RunOutput: false},
		{Script: `1++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunError: fmt.Errorf("invalid operation")},
		{Script: `3++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunError: fmt.Errorf("invalid operation")},
		{Script: `a=1;a++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunOutput: true},
		{Script: `a=3;a++ in l`, Input: map[string]interface{}{"l": []interface{}{1, 2, 3}}, RunOutput: false},

		// multidimensional slice
		{Script: `1 in [1]`, RunOutput: true},
		{Script: `1 in [[1]]`, RunOutput: false},
		{Script: `1 in [[[1]]]`, RunOutput: false},
		{Script: `1 in [[1],[[1]],1]`, RunOutput: true},
		{Script: `[1] in [1]`, RunOutput: false},
		{Script: `[1] in [[1]]`, RunOutput: true},
		{Script: `[1] in [[[1]]]`, RunOutput: false},
		{Script: `[[1]] in [1]`, RunOutput: false},
		{Script: `[[1]] in [[1]]`, RunOutput: false},
		{Script: `[[1]] in [[[1]]]`, RunOutput: true},
		{Script: `1 in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: true},
		{Script: `1 in [[1]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, RunOutput: false},
		{Script: `1 in [[[1]]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, RunOutput: false},
		{Script: `1 in [[1],[[1]],1]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}, []interface{}{[]interface{}{1}}, 1}}, RunOutput: true},
		{Script: `[1] in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `[1] in [[1]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, RunOutput: true},
		{Script: `[1] in [[[1]]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, RunOutput: false},
		{Script: `[[1]] in [1]`, Input: map[string]interface{}{"l": []interface{}{1}}, RunOutput: false},
		{Script: `[[1]] in [[1]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{1}}}, RunOutput: false},
		{Script: `[[1]] in [[[1]]]`, Input: map[string]interface{}{"l": []interface{}{[]interface{}{[]interface{}{1}}}}, RunOutput: true},
	}
	testlib.Run(t, tests, nil)
}
