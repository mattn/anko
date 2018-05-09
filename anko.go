// +build !appengine

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/anko/core"
	"github.com/mattn/anko/packages"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
)

const version = "0.0.3"

var (
	flagExecute string
	file        string
	args        []string
	env         *vm.Env
)

func main() {
	var exitCode int

	parseFlags()
	setupEnv()
	if flagExecute != "" || flag.NArg() > 0 {
		exitCode = runNonInteractive()
	} else {
		exitCode = runInteractive()
	}

	os.Exit(exitCode)
}

func parseFlags() {
	flagVersion := flag.Bool("v", false, "prints out the version and then exits")
	flag.StringVar(&flagExecute, "e", "", "execute the Anko code")
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if flagExecute != "" || flag.NArg() < 1 {
		args = flag.Args()
		return
	}

	file = flag.Arg(0)
	args = flag.Args()[1:]
}

func setupEnv() {
	env = vm.NewEnv()
	env.Define("args", args)
	core.Import(env)
	packages.DefineImport(env)
}

func runNonInteractive() int {
	var source string
	if flagExecute != "" {
		source = flagExecute
	} else {
		sourceBytes, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("ReadFile error:", err)
			return 2
		}
		source = string(sourceBytes)
	}

	_, err := env.Execute(source)
	if err != nil {
		fmt.Println("Execute error:", err)
		return 4
	}

	return 0
}

func runInteractive() int {
	var following bool
	var source string
	scanner := bufio.NewScanner(os.Stdin)

	parser.EnableErrorVerbose()

	for {
		if following {
			fmt.Print("  ")
		} else {
			fmt.Print("> ")
		}

		if !scanner.Scan() {
			break
		}
		source += scanner.Text()
		if source == "" {
			continue
		}
		if source == "quit()" {
			break
		}

		stmts, err := parser.ParseSrc(source)

		if e, ok := err.(*parser.Error); ok {
			es := e.Error()
			if strings.HasPrefix(es, "syntax error: unexpected") {
				if strings.HasPrefix(es, "syntax error: unexpected $end,") {
					following = true
					continue
				}
			} else {
				if e.Pos.Column == len(source) && !e.Fatal {
					fmt.Fprintln(os.Stderr, e)
					following = true
					continue
				}
				if e.Error() == "unexpected EOF" {
					following = true
					continue
				}
			}
		}

		following = false
		source = ""
		var v interface{}

		if err == nil {
			v, err = vm.Run(stmts, env)
		}
		if err != nil {
			if e, ok := err.(*vm.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
			} else if e, ok := err.(*parser.Error); ok {
				fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
			} else {
				fmt.Fprintln(os.Stderr, err)
			}
			continue
		}

		fmt.Printf("%#v\n", v)
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "ReadString error:", err)
			return 12
		}
	}

	return 0
}
