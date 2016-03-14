// +build !appengine

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/daviddengcn/go-colortext"
	"github.com/mattn/anko/parser"
	"github.com/mattn/anko/vm"
	"github.com/mattn/go-isatty"

	anko_core "github.com/mattn/anko/builtins"
	anko_encoding_json "github.com/mattn/anko/builtins/encoding/json"
	anko_flag "github.com/mattn/anko/builtins/flag"
	anko_fmt "github.com/mattn/anko/builtins/fmt"
	anko_io "github.com/mattn/anko/builtins/io"
	anko_io_ioutil "github.com/mattn/anko/builtins/io/ioutil"
	anko_math "github.com/mattn/anko/builtins/math"
	anko_math_rand "github.com/mattn/anko/builtins/math/rand"
	anko_net "github.com/mattn/anko/builtins/net"
	anko_net_http "github.com/mattn/anko/builtins/net/http"
	anko_net_url "github.com/mattn/anko/builtins/net/url"
	anko_os "github.com/mattn/anko/builtins/os"
	anko_os_exec "github.com/mattn/anko/builtins/os/exec"
	anko_path "github.com/mattn/anko/builtins/path"
	anko_path_filepath "github.com/mattn/anko/builtins/path/filepath"
	anko_regexp "github.com/mattn/anko/builtins/regexp"
	anko_sort "github.com/mattn/anko/builtins/sort"
	anko_strings "github.com/mattn/anko/builtins/strings"

	anko_colortext "github.com/mattn/anko/builtins/github.com/daviddengcn/go-colortext"
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

	tbl := map[string]func(env *vm.Env) *vm.Env{
		"encoding/json": anko_encoding_json.Import,
		"flag":          anko_flag.Import,
		"fmt":           anko_fmt.Import,
		"io":            anko_io.Import,
		"io/ioutil":     anko_io_ioutil.Import,
		"math":          anko_math.Import,
		"math/rand":     anko_math_rand.Import,
		"net":           anko_net.Import,
		"net/http":      anko_net_http.Import,
		"net/url":       anko_net_url.Import,
		"os":            anko_os.Import,
		"os/exec":       anko_os_exec.Import,
		"path":          anko_path.Import,
		"path/filepath": anko_path_filepath.Import,
		"regexp":        anko_regexp.Import,
		"sort":          anko_sort.Import,
		"strings":       anko_strings.Import,
		"github.com/daviddengcn/go-colortext": anko_colortext.Import,
	}

	env.Define("import", reflect.ValueOf(func(s string) interface{} {
		if loader, ok := tbl[s]; ok {
			return loader(env)
		}
		panic(fmt.Sprintf("package '%s' not found", s))
	}))

	for {
		if repl {
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

		scanner := new(parser.Scanner)
		scanner.Init(code)
		stmts, err := parser.Parse(scanner)

		v := vm.NilValue

		if repl {
			if e, ok := err.(*parser.Error); ok {
				if e.Pos.Column == len(b) && !e.Fatal {
					following = true
					continue
				}
				if e.Error() == "unexpected EOF" {
					following = true
					continue
				}
				if strings.HasPrefix(e.Error(), "syntax error: unexpected $end,") {
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
					fmt.Fprintf(os.Stderr, "%s:%d: %s\n", source, e.Pos.Line, err)
				} else if e, ok := err.(*parser.Error); ok {
					if e.Filename != "" {
						source = e.Filename
					}
					fmt.Fprintf(os.Stderr, "%s:%d: %s\n", source, e.Pos.Line, err)
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
