package vm

import (
	"fmt"
	"os"
	"testing"
)

func TestIf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `if 1++ {}`, runError: fmt.Errorf("Invalid operation")},
		{script: `if false {} else if 1++ {}`, runError: fmt.Errorf("Invalid operation")},
		{script: `if false {} else if true { 1++ }`, runError: fmt.Errorf("Invalid operation")},

		{script: `if true {}`, input: map[string]interface{}{"a": nil}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: `if true {}`, input: map[string]interface{}{"a": true}, runOutput: nil, output: map[string]interface{}{"a": true}},
		{script: `if true {}`, input: map[string]interface{}{"a": int64(1)}, runOutput: nil, output: map[string]interface{}{"a": int64(1)}},
		{script: `if true {}`, input: map[string]interface{}{"a": float64(1.1)}, runOutput: nil, output: map[string]interface{}{"a": float64(1.1)}},
		{script: `if true {}`, input: map[string]interface{}{"a": "a"}, runOutput: nil, output: map[string]interface{}{"a": "a"}},

		{script: `if true {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: `if true {a = true}`, input: map[string]interface{}{"a": int64(2)}, runOutput: true, output: map[string]interface{}{"a": true}},
		{script: `if true {a = 1}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: `if true {a = 1.1}`, input: map[string]interface{}{"a": int64(2)}, runOutput: float64(1.1), output: map[string]interface{}{"a": float64(1.1)}},
		{script: `if true {a = "a"}`, input: map[string]interface{}{"a": int64(2)}, runOutput: "a", output: map[string]interface{}{"a": "a"}},

		{script: `if a == 1 {a = 1}`, input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: `if a == 2 {a = 1}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: `if a == 1 {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: `if a == 2 {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},

		{script: `if a == 1 {a = 1} else {a = 3}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), output: map[string]interface{}{"a": int64(3)}},
		{script: `if a == 2 {a = 1} else {a = 3}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(1), output: map[string]interface{}{"a": int64(1)}},
		{script: `if a == 1 {a = 1} else if a == 3 {a = 3}`, input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: `if a == 1 {a = 1} else if a == 2 {a = 3}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(3), output: map[string]interface{}{"a": int64(3)}},
		{script: `if a == 1 {a = 1} else if a == 3 {a = 3} else {a = 4}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), output: map[string]interface{}{"a": int64(4)}},

		{script: `if a == 1 {a = 1} else {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: `if a == 2 {a = nil} else {a = 3}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: `if a == 1 {a = nil} else if a == 3 {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: false, output: map[string]interface{}{"a": int64(2)}},
		{script: `if a == 1 {a = 1} else if a == 2 {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: `if a == 1 {a = 1} else if a == 3 {a = 3} else {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},

		{script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = 5}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(5), output: map[string]interface{}{"a": int64(5)}},
		{script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 4 {a = 4} else {a = nil}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
		{script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = 4} else {a = 5}`, input: map[string]interface{}{"a": int64(2)}, runOutput: int64(4), output: map[string]interface{}{"a": int64(4)}},
		{script: `if a == 1 {a = 1} else if a == 3 {a = 3} else if a == 2 {a = nil} else {a = 5}`, input: map[string]interface{}{"a": int64(2)}, runOutput: nil, output: map[string]interface{}{"a": nil}},
	}
	runTests(t, tests)
}

func TestSwitch(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		// TODO: add tests

		{script: `switch {}`, parseError: fmt.Errorf("syntax error")},
		{script: `a=1; switch a; {}`, parseError: fmt.Errorf("syntax error")},

		{script: `a=1; switch a {}`, runOutput: int64(1)},

		{script: `a=1; switch a {case 1: return a}`, runOutput: int64(1)},
		{script: `a=2; switch a {case 1,2: return a}`, runOutput: int64(2)},
		{script: `a=3; switch a {case 1,2,3: return a}`, runOutput: int64(3)},

		{script: `a=2; switch a {case 1: return a;case 2: return a}`, runOutput: int64(2)},
		{script: `a=3; switch a {case 1,3: return a;case 2: return a}`, runOutput: int64(3)},
		{script: `a=3; switch a {case 1: return a;case 2,3: return a}`, runOutput: int64(3)},

		{script: `a=99; switch a {case 1: return a;default: return 'default'}`, runOutput: "default"},
		{script: `a=99; switch a {case 1,2: return a;default: return 'default'}`, runOutput: "default"},
		{script: `a=99; switch a {case 1,2,3: return a;default: return 'default'}`, runOutput: "default"},
		{script: `a=99; switch a {case 1: return a;case 2: return a;default: return 'default'}`, runOutput: "default"},
		{script: `a=99; switch a {case 1,3: return a;case 2: return a;default: return 'default'}`, runOutput: "default"},
		{script: `a=99; switch a {case 1: return a;case 2,3: return a;default: return 'default'}`, runOutput: "default"},
		{script: `a=99; switch a {case 1: return a;default: return 'default';default: return 'default2'}`, parseError: fmt.Errorf("multiple default statement")},
		{script: `a=99; switch a {case 1,2: return a;default: return 'default';default: return 'default2'}`, parseError: fmt.Errorf("multiple default statement")},
		{script: `a=99; switch a {case 1: return a;case 2: return a;default: return 'default';default: return 'default2'}`, parseError: fmt.Errorf("multiple default statement")},
	}
	runTests(t, tests)
}
