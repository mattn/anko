package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func main() {
	env := vm.Env{}

	env["println"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		for i, arg := range args {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Print(arg.Interface())
		}
		fmt.Println()
		return vm.NilValue, nil
	})

	env["len"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.Array && args[0].Kind() != reflect.Slice {
			return vm.NilValue, errors.New("Argument should be array")
		}
		return reflect.ValueOf(args[0].Len()), nil
	})

	env["keys"] = vm.ToFunc(func(args ...reflect.Value) (reflect.Value, error) {
		if len(args) < 1 {
			return vm.NilValue, errors.New("Missing arguments")
		}
		if len(args) > 1 {
			return vm.NilValue, errors.New("Too many arguments")
		}
		if args[0].Kind() != reflect.Map {
			return vm.NilValue, errors.New("Argument should be map")
		}
		keys := []string{}
		mk := args[0].MapKeys()
		for _, key := range mk {
			keys = append(keys, key.String())
		}
		return reflect.ValueOf(keys), nil
	})

	if len(os.Args) > 1 {
		scanner := new(parser.Scanner)
		body, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		args := []reflect.Value{}
		for _, arg := range os.Args[2:] {
			args = append(args, reflect.ValueOf(arg))
		}

		scanner.Init(string(body))
		stmts, err := parser.Parse(scanner)
		if err != nil {
			log.Fatal(err)
		}
		for _, stmt := range stmts {
			_, err := vm.Run(stmt, env)
			if err != nil {
				log.Fatal(err)
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
