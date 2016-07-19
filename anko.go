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
	anko_errors "github.com/mattn/anko/builtins/errors"
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
	anko_os_signal "github.com/mattn/anko/builtins/os/signal"
	anko_path "github.com/mattn/anko/builtins/path"
	anko_path_filepath "github.com/mattn/anko/builtins/path/filepath"
	anko_regexp "github.com/mattn/anko/builtins/regexp"
	anko_runtime "github.com/mattn/anko/builtins/runtime"
	anko_sort "github.com/mattn/anko/builtins/sort"
	anko_strings "github.com/mattn/anko/builtins/strings"
	anko_time "github.com/mattn/anko/builtins/time"

	anko_colortext "github.com/mattn/anko/builtins/github.com/daviddengcn/go-colortext"
)

const version = "0.0.1"

var (
	fs   = flag.NewFlagSet(os.Args[0], 1)
	line = fs.String("e", "", "One line of program")
	v    = fs.Bool("v", false, "Display version")

	istty = isatty.IsTerminal(os.Stdout.Fd())
)

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

	var (
		code      string
		b         []byte
		reader    *bufio.Reader
		following bool
		source    string
	)

	env := vm.NewEnv()
	interactive := fs.NArg() == 0 && *line == ""

	env.Define("args", fs.Args())

	if interactive {
		reader = bufio.NewReader(os.Stdin)
		source = "typein"
		os.Args = append([]string{os.Args[0]}, fs.Args()...)
	} else {
		if *line != "" {
			b = []byte(*line)
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
			env.Define("args", fs.Args()[1:])
			source = filepath.Clean(fs.Arg(0))
		}
		os.Args = fs.Args()
	}

	anko_core.Import(env)

	pkgs := map[string]func(env *vm.Env) *vm.Env{
		"encoding/json": anko_encoding_json.Import,
		"errors":        anko_errors.Import,
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
		"os/signal":     anko_os_signal.Import,
		"path":          anko_path.Import,
		"path/filepath": anko_path_filepath.Import,
		"regexp":        anko_regexp.Import,
		"runtime":       anko_runtime.Import,
		"sort":          anko_sort.Import,
		"strings":       anko_strings.Import,
		"time":          anko_time.Import,
		"github.com/daviddengcn/go-colortext": anko_colortext.Import,
	}

	env.Define("import", func(s string) interface{} {
		if loader, ok := pkgs[s]; ok {
			m := loader(env)
			return m
		}
		panic(fmt.Sprintf("package '%s' not found", s))
	})

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
		v := vm.NilValue

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

			if interactive {
				continue
			} else {
				os.Exit(1)
			}
		} else {
			if interactive {
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
