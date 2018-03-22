# anko

[![GoDoc Reference](https://godoc.org/github.com/mattn/anko/vm?status.svg)](http://godoc.org/github.com/mattn/anko/vm)
[![Build Status](https://travis-ci.org/mattn/anko.png?branch=master)](https://travis-ci.org/mattn/anko)
[![Coverage](https://codecov.io/gh/mattn/anko/branch/master/graph/badge.svg)](https://codecov.io/gh/mattn/anko)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattn/anko)](https://goreportcard.com/report/github.com/mattn/anko)

Anko is a scriptable interpreter written in Go.

![](https://raw.githubusercontent.com/mattn/anko/master/anko.png)

(Picture licensed under CC BY-SA 3.0 by wikipedia)

## Installation
Requires Go.
```
$ go get -u github.com/mattn/anko
```

## Examples

```bash
# declare function
func plus(n){
  return n + 1
}

# declare variables
x = 1
y = x + 1

# print values
println(x * (y + 2 * x + plus(x) / 2))

# if/else condition
if plus(y) > 1 {
  println("こんにちわ世界")
} else {
  println("Hello, World")
}

# array type
a = [1,2,3]
println(a[2])
println(len(a))

# map type
m = {"foo": "bar", "far": "boo"}
m.foo = "baz"
for k in keys(m) {
  println(m[k])
}
```

See `_examples/scripts` for more examples.



## Usage

Embedding the interpreter into your own program:

```Go
var env = vm.NewEnv()

env.Define("foo", 1)
env.Define("bar", func() int {
	return 2
})

val, err := env.Execute(`foo + bar()`)
if err != nil {
	panic(err)
}

fmt.Println(val)
// output:
// 3
```

To import all core language builtins, allowing the example scripts to work:

```Go
import "github.com/mattn/anko/core"

var env = vm.NewEnv()
core.Import(env)

_, err := env.Execute(`println("test")`)
if err != nil {
	panic(err)
}
// output:
// "test"
```

Running scripts using anko command-line tool:

```
$ anko script.ank
```

# Author

Yasuhiro Matsumoto (a.k.a mattn)

# Note

Please note that the master branch is not stable,  the language and API may change at any time.

To mitigate breaking changes, please use tagged branches. New tagged branches will be created for breaking changes.
