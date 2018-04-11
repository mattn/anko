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
	flagExecute *string
	env         *vm.Env
)

func main() {
	parseFlags()
	setupEnv()
	runNonInteractive()
	runInteractive()
}

func parseFlags() {
	flagVersion := flag.Bool("v", false, "prints out the version and then exits")
	flagExecute = flag.String("e", "", "execute the Anko code")
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		os.Exit(0)
	}
}

func setupEnv() {
	env = vm.NewEnv()
	if len(*flagExecute) > 0 || flag.NArg() < 1 {
		env.Define("args", flag.Args())
	} else {
		env.Define("args", flag.Args()[1:])
	}
	core.Import(env)
	AddPackageColortext()
	packages.DefineImport(env)
}

func runNonInteractive() {
	if len(*flagExecute) < 1 && flag.NArg() < 1 {
		return
	}

	var source string
	if len(*flagExecute) > 1 {
		source = *flagExecute
	} else {
		sourceBytes, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			fmt.Println("ReadFile error:", err)
			os.Exit(2)
		}
		source = string(sourceBytes)
	}

	_, err := env.Execute(source)
	if err != nil {
		fmt.Println("Execute error:", err)
		os.Exit(4)
	}

	os.Exit(0)
}

func runInteractive() {
	var reader *bufio.Reader
	var source string
	var b []byte
	var code string
	var following bool

	interactive := true
	reader = bufio.NewReader(os.Stdin)
	source = "typein"

	for {
		if interactive {
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
			if code != "" {
				code += "\n"
			}
			code += string(b)
		} else {
			code = string(b)
		}

		parser.EnableErrorVerbose()

		stmts, err := parser.ParseSrc(code)

		if interactive {
			if e, ok := err.(*parser.Error); ok {
				es := e.Error()
				if strings.HasPrefix(es, "syntax error: unexpected") {
					if strings.HasPrefix(es, "syntax error: unexpected $end,") {
						following = true
						continue
					}
				} else {
					if e.Pos.Column == len(b) && !e.Fatal {
						println(e.Error())
						following = true
						continue
					}
					if e.Error() == "unexpected EOF" {
						following = true
						continue
					}
				}
			}
		}

		following = false
		code = ""
		var v interface{}

		if err == nil {
			v, err = vm.Run(stmts, env)
		}
		if err != nil {
			colortext(ct.Red, false, func() {
				if e, ok := err.(*vm.Error); ok {
					fmt.Fprintf(os.Stderr, "%s:%d:%d %s\n", source, e.Pos.Line, e.Pos.Column, err)
				} else if e, ok := err.(*parser.Error); ok {
					if e.Filename != "" {
						source = e.Filename
					}
					fmt.Fprintf(os.Stderr, "%s:%d:%d %s\n", source, e.Pos.Line, e.Pos.Column, err)
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			})

			if interactive {
				continue
			} else {
				os.Exit(1)
			}
		} else {
			if interactive {
				colortext(ct.Black, true, func() {
					fmt.Printf("%#v\n", v)
				})
			} else {
				break
			}
		}
	}
}
