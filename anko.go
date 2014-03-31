package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	anko_core "github.com/mattn/anko/builtins/core"
	anko_http "github.com/mattn/anko/builtins/http"
	anko_json "github.com/mattn/anko/builtins/json"
	anko_os "github.com/mattn/anko/builtins/os"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

var e = flag.String("e", "", "One line of program")

func main() {
	flag.Parse()

	env := vm.NewEnv()

	anko_core.Import(env)
	anko_http.Import(env)
	anko_json.Import(env)
	anko_os.Import(env)

	if flag.NArg() > 0 || *e != "" {
		var code string
		if *e != "" {
			code = *e
			env.Set("args", reflect.ValueOf(flag.Args()))
		} else {
			body, err := ioutil.ReadFile(flag.Arg(0))
			if err != nil {
				ct.ChangeColor(ct.Red, false, ct.None, false)
				fmt.Fprintln(os.Stderr, err)
				ct.ResetColor()
				os.Exit(1)
			}
			env.Set("args", reflect.ValueOf(flag.Args()[1:]))
			code = string(body)
		}

		scanner := new(parser.Scanner)
		scanner.Init(code)
		stmts, err := parser.Parse(scanner)
		if err != nil {
			ct.ChangeColor(ct.Red, false, ct.None, false)
			fmt.Fprintln(os.Stderr, err)
			ct.ResetColor()
		} else {
			_, err := vm.RunStmts(stmts, env)
			if err != nil {
				ct.ChangeColor(ct.Red, false, ct.None, false)
				if e, ok := err.(*vm.Error); ok {
					fmt.Fprintf(os.Stderr, "%s:%d: %s\n", flag.Arg(0), e.Pos().Line, err)
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
				ct.ResetColor()
				os.Exit(1)
			}
		}
	} else {
		env.Define("args", reflect.ValueOf([]string{}))
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
				s = strings.TrimSpace(string(b) + ";")
				scanner = new(parser.Scanner)
				scanner.Init(s)
				stmts, err = parser.Parse(scanner)
				if err != nil {
					ct.ChangeColor(ct.Red, false, ct.None, false)
					if e, ok := err.(*vm.Error); ok {
						fmt.Fprintf(os.Stderr, "typein:%d: %s\n", e.Pos().Line, err)
					} else {
						fmt.Fprintln(os.Stderr, err)
					}
					ct.ResetColor()
				}
			}

			if err == nil {
				v, err := vm.RunStmts(stmts, env)
				if err != nil {
					ct.ChangeColor(ct.Red, false, ct.None, false)
					if e, ok := err.(*vm.Error); ok {
						fmt.Fprintf(os.Stderr, "typein:%d: %s\n", e.Pos().Line, err)
					} else {
						fmt.Fprintln(os.Stderr, err)
					}
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
}
