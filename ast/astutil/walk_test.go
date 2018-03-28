package astutil

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mattn/anko/ast"
	"github.com/mattn/anko/parser"
)

const (
	goodSrc string = `
var fmt = import("fmt")

a = "1"
b = 2
func testA(arg1, arg2, arg3) {
	return "A" + arg1 + arg2 + arg3
}

func Main(arg1) {
	fmt.Println("enter Main")
	b = testA(1, 2, 3) + Tester()

	if b == 0 {
		fmt.Println("b is 0")
	} else if b == 1 {
		fmt.Println("b is 1")
	} else {
		fmt.Println("b is other")
	}

	switch arg1 {
	case 0:
		fmt.Println("arg0 is 0")
	case 1:
		fmt.Println("arg0 is 1")
	default:
		fmt.Println("arg0 is other")
	}

	try {
		throw "WTF!"
	} catch e {
		fmt.Println(e)
	}

	for n = 0; n < 3; n++ {
		if n < 2 {
			continue
		}
		fmt.Println(n)
	}

	for n in [1, 2, 3, 4, 5] {
		fmt.Println(n)
		if n > 3 {
			break
		}
	}

	n = 0
	for n < 3 {
		n++
	}

	a = {"foo": "bar"}
	a.foo = "baz"
	if a["foo"] == "zoo" {
		fmt.Println("foo is zoo")
	}
	fmt.Println(a["foo"] == "zoo" ? "zoo" : "baz")

	c = make(chan int64)

	go func() {
	  c <- 1
	  c <- 2
	  c <- 3
	}()

	println(<-c)
	println(<-c)
	println(<-c)

	v = make([]int, 3)
	fmt.Println("sizeof v is ", len(v))

	x = 1
	y = (&x)
	*y = 2
	fmt.Println(x)

	x, y = !x, 2
	fmt.Println(x, y)

	x = new(string)
	fmt.Println(x)

	var f = func() {
		return "foo"
	}
	x = f()[0:1]
	y = f()[0]
	fmt.Println(x == y ? true : false)
}

func Tester() {
	return "YES"
}

func testLen() {
	return len("test")
}

fmt.Println(Main(1))
`
)

func TestWalk(t *testing.T) {
	stmts, err := parser.ParseSrc(goodSrc)
	if err != nil {
		t.Fatal(err)
	}
	var mainFound bool
	var lenFound bool
	err = Walk(stmts, func(e interface{}) error {
		switch exp := e.(type) {
		case *ast.CallExpr:
			switch exp.Name {
			case `testA`:
				if len(exp.SubExprs) != 3 {
					return errors.New("Invalid parameter count")
				}
			case `Main`:
				if len(exp.SubExprs) != 1 {
					return errors.New("Invalid parameter count")
				}
			case `Tester`:
				if len(exp.SubExprs) != 0 {
					return errors.New("Invalid parameter count")
				}
			}
		case *ast.FuncExpr:
			if !mainFound && exp.Name == `Main` {
				mainFound = true
			} else if mainFound && exp.Name == `Main` {
				return errors.New("Main redefined")
			}
		case *ast.LenExpr:
			lenFound = true
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !mainFound {
		t.Fatal("Main not found")
	}
	if !lenFound {
		t.Fatal("len not found")
	}
}

func Example_astWalk() {
	src := `
var fmt = import("fmt")

func TestFunc(arg1, arg2, arg3) {
	return (arg1 + arg2) * arg3
}

func Main() {
	return TestFunc(1, 2, 3) + BuiltinFuncX(1, 2, 3)
}

fmt.Println(Main())
	`
	stmts, err := parser.ParseSrc(src)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	var mainFound bool
	err = Walk(stmts, func(e interface{}) error {
		switch e := e.(type) {
		case *ast.CallExpr:
			//check if the BuiltinFuncX is getting the right number of args
			if e.Name == `BuiltinFuncX` && len(e.SubExprs) != 3 {
				return errors.New("Invalid number of arguments to BuiltinFuncX")
			}
		case *ast.FuncExpr:
			if !mainFound && e.Name == `Main` {
				if len(e.Params) != 0 {
					return errors.New("Too many args to main")
				}
				mainFound = true
			} else if mainFound && e.Name == `Main` {
				return errors.New("Main redefined")
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	if !mainFound {
		fmt.Println("ERROR: Main not found")
	}
}

func TestBadCode(t *testing.T) {
	var codes = []string{
		`const 1 = 2`,
		`a["foo"] = 2, 3`,
		`if a, 2 {}`,
		`if break {}`,
		`if a { else }`,
		`if a { } else foo { }`,
		`try a { } else { }`,
		`try { } catch 1, 2 { }`,
		`throw 1, 2`,
		`for.true`,
		`switch { var }`,
		`case { var }`,
	}
	for _, code := range codes {
		_, err := parser.ParseSrc(code)
		if err == nil {
			t.Fatalf("code %q should fail", code)
		}
	}
}
