package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	env := vm.NewEnv()

	setupBuiltins(env)

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
			scanner := new(parser.Scanner)
			fmt.Print("> ")
			b, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			scanner.Init(string(b))
			stmts, err := parser.Parse(scanner)
			if err != nil {
				fmt.Println(err)
			}
			for _, stmt := range stmts {
				_, err := vm.Run(stmt, env)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
