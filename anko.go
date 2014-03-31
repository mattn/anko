package main

import (
	"bufio"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/anko/builtins"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func main() {
	env := vm.NewEnv()

	builtins.SetupBuiltins(env)

	if len(os.Args) > 1 {
		scanner := new(parser.Scanner)
		body, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		args := []reflect.Value{}
		for _, arg := range os.Args[2:] {
			args = append(args, reflect.ValueOf(arg))
		}

		scanner.Init(string(body))
		stmts, err := parser.Parse(scanner)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		for _, stmt := range stmts {
			_, err := vm.Run(stmt, env)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			b, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			if len(b) == 0 {
				continue
			}
			s := strings.TrimSpace(string(b))
			scanner := new(parser.Scanner)
			scanner.Init(s)
			stmts, err := parser.Parse(scanner)
			if err != nil {
				fmt.Println(err)
			}
			v, err := vm.RunStmts(stmts, env)
			if err != nil {
				ct.ChangeColor(ct.Red, false, ct.None, false)
				fmt.Fprintln(os.Stderr, err)
				ct.ResetColor()
			} else {
				ct.ChangeColor(ct.Black, true, ct.None, false)
				if v == vm.NilValue {
					fmt.Println("nil")
				} else {
					fmt.Println(v.Interface())
				}
				ct.ResetColor()
			}
		}
	}
}
