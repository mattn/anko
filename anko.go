package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	anko_core "github.com/mattn/anko/builtins"
	anko_encoding "github.com/mattn/anko/builtins/encoding"
	anko_flag "github.com/mattn/anko/builtins/flag"
	anko_io "github.com/mattn/anko/builtins/io"
	anko_math "github.com/mattn/anko/builtins/math"
	anko_net "github.com/mattn/anko/builtins/net"
	anko_os "github.com/mattn/anko/builtins/os"
	anko_path "github.com/mattn/anko/builtins/path"
	anko_regexp "github.com/mattn/anko/builtins/regexp"
	anko_sort "github.com/mattn/anko/builtins/sort"
	anko_strings "github.com/mattn/anko/builtins/strings"
	anko_term "github.com/mattn/anko/builtins/term"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"github.com/mattn/go-isatty"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

const version = "0.0.1"

var fs = flag.NewFlagSet(os.Args[0], 1)
var e = fs.String("e", "", "One line of program")
var v = fs.Bool("v", false, "Display version")

var istty = isatty.IsTerminal(os.Stdout.Fd())

func colortext(color ct.Color, bright bool, f func()) {
	if istty {
		ct.ChangeColor(color, bright, ct.None, false)
		f()
		ct.ResetColor()
	} else {
		f()
	}
}

func main() {
	fs.Parse(os.Args[1:])
	if *v {
		fmt.Println(version)
		os.Exit(0)
	}

	env := vm.NewEnv()

	var code string
	var b []byte
	var reader *bufio.Reader
	var following bool
	var source string

	repl := fs.NArg() == 0 && *e == ""

	env.Define("args", reflect.ValueOf(fs.Args()))

	if repl {
		reader = bufio.NewReader(os.Stdin)
		source = "typein"
		os.Args = append([]string{os.Args[0]}, fs.Args()...)
	} else {
		if *e != "" {
			b = []byte(*e)
			source = "argument"
		} else {
			var err error
			b, err = ioutil.ReadFile(fs.Arg(0))
			if err != nil {
				colortext(ct.Red, false, func() {
					fmt.Fprintln(os.Stderr, err)
				})
				os.Exit(1)
			}
			env.Define("args", reflect.ValueOf(fs.Args()[1:]))
			source = filepath.Clean(fs.Arg(0))
		}
		os.Args = fs.Args()
	}

	anko_core.Import(env)
	anko_flag.Import(env)
	anko_net.Import(env)
	anko_encoding.Import(env)
	anko_os.Import(env)
	anko_io.Import(env)
	anko_math.Import(env)
	anko_path.Import(env)
	anko_regexp.Import(env)
	anko_sort.Import(env)
	anko_strings.Import(env)
	anko_term.Import(env)

	for {
		if repl {
			colortext(ct.Green, true, func() {
				if following {
					fmt.Print("  ")
				} else {
					fmt.Print("> ")
				}
			})
			b, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			if len(b) == 0 {
				continue
			}
			code += string(b)
		} else {
			code = string(b)
		}

		scanner := new(parser.Scanner)
		scanner.Init(code)
		stmts, err := parser.Parse(scanner)

		v := vm.NilValue

		if repl {
			if following {
				continue
			}
			if e, ok := err.(*parser.Error); ok && e.Pos().Column == len(b) {
				if !e.Fatal() {
					following = true
					continue
				}
			}
		}
		following = false
		code = ""
		if err == nil {
			v, err = vm.Run(stmts, env)
		}

		if err != nil {
			colortext(ct.Red, false, func() {
				if e, ok := err.(*vm.Error); ok {
					fmt.Fprintf(os.Stderr, "%s:%d: %s\n", source, e.Pos().Line, err)
				} else if e, ok := err.(*parser.Error); ok {
					fmt.Fprintf(os.Stderr, "%s:%d: %s\n", source, e.Pos().Line, err)
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			})

			if repl {
				continue
			} else {
				os.Exit(1)
			}
		} else {
			if repl {
				colortext(ct.Black, true, func() {
					if v == vm.NilValue || !v.IsValid() {
						fmt.Println("nil")
					} else {
						s, ok := v.Interface().(fmt.Stringer)
						if v.Kind() != reflect.String && ok {
							fmt.Println(s)
						} else {
							fmt.Printf("%#v\n", v.Interface())
						}
					}
				})
			} else {
				break
			}
		}
	}
}
