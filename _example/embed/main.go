package main

import (
	"fmt"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"log"
	"reflect"
)

func main() {
	env := vm.NewEnv()
	env.Define("foo", reflect.ValueOf(1))

	scanner := new(parser.Scanner)
	scanner.Init(`foo + 1`)
	stmts, err := parser.Parse(scanner)
	if err != nil {
		log.Fatal(err)
	}
	v, err := vm.Run(stmts, env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v.Interface())
}
