package vm

import (
	"fmt"
	"os"
	"testing"
)

func TestIf(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
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
	}
	RunTests(t, tests, nil)
}

func TestSwitch(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []Test{
		{Script: `switch {}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a=1; switch a {case: return a}`, ParseError: fmt.Errorf("syntax error")},
		{Script: `a=1; switch a; {}`, ParseError: fmt.Errorf("syntax error")},

		{Script: `a=1; switch a {}`, RunOutput: int64(1)},

		{Script: `a=1; switch a {case 1: return a}`, RunOutput: int64(1)},
		{Script: `a=2; switch a {case 1,2: return a}`, RunOutput: int64(2)},
		{Script: `a=3; switch a {case 1,2,3: return a}`, RunOutput: int64(3)},

		{Script: `a=2; switch a {case 1: return a; case 2: return a}`, RunOutput: int64(2)},
		{Script: `a=3; switch a {case 1,3: return a; case 2: return a}`, RunOutput: int64(3)},
		{Script: `a=3; switch a {case 1: return a; case 2,3: return a}`, RunOutput: int64(3)},

		{Script: `a=99; switch a {default: return 'default'}`, RunOutput: "default"},
		{Script: `a=99; switch a {case 1: return a; default: return 'default'}`, RunOutput: "default"},
		{Script: `a=99; switch a {case 1,2: return a; default: return 'default'}`, RunOutput: "default"},
		{Script: `a=99; switch a {case 1,2,3: return a; default: return 'default'}`, RunOutput: "default"},
		{Script: `a=99; switch a {case 1: return a; case 2: return a; default: return 'default'}`, RunOutput: "default"},
		{Script: `a=99; switch a {case 1,3: return a; case 2: return a; default: return 'default'}`, RunOutput: "default"},
		{Script: `a=99; switch a {case 1: return a; case 2,3: return a; default: return 'default'}`, RunOutput: "default"},

		// TODO: After a parse error, the second default statement still runs and returns. Is this something to fix?
		{Script: `a=99; switch a {case 1: return a; default: return 'default'; default: return 'default2'}`, ParseError: fmt.Errorf("multiple default statement"), RunOutput: "default2"},
		{Script: `a=99; switch a {case 1,2: return a; default: return 'default'; default: return 'default2'}`, ParseError: fmt.Errorf("multiple default statement"), RunOutput: "default2"},
		{Script: `a=99; switch a {case 1: return a; case 2: return a; default: return 'default'; default: return 'default2'}`, ParseError: fmt.Errorf("multiple default statement"), RunOutput: "default2"},
	}
	RunTests(t, tests, nil)
}
