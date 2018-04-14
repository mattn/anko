// +build !appengine

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/anko/core"
	"github.com/mattn/anko/packages"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"

	"github.com/daviddengcn/go-colortext"
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
	AddPackageColortext()
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
	var source string
	var b []byte
	var following bool
	reader := bufio.NewReader(os.Stdin)

	parser.EnableErrorVerbose()

	for {
		colortext(ct.Green, true, func() {
			if following {
				fmt.Print("  ")
			} else {
				fmt.Print("> ")
			}
		})
		var err error
		b, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		if len(b) == 0 {
			continue
		}
		if source != "" {
			source += "\n"
		}
		source += string(b)

		if source == "quit()" {
			return 0
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
				if e.Pos.Column == len(b) && !e.Fatal {
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
			colortext(ct.Red, false, func() {
				if e, ok := err.(*vm.Error); ok {
					fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
				} else if e, ok := err.(*parser.Error); ok {
					fmt.Fprintf(os.Stderr, "%d:%d %s\n", e.Pos.Line, e.Pos.Column, err)
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			})

			continue
		} else {
			colortext(ct.Black, true, func() {
				fmt.Printf("%#v\n", v)
			})
		}
	}

	return 0
}
