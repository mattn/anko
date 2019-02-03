# Anko

[![GoDoc Reference](https://godoc.org/github.com/mattn/anko/vm?status.svg)](http://godoc.org/github.com/mattn/anko/vm)
[![Build Status](https://travis-ci.org/mattn/anko.svg?branch=master)](https://travis-ci.org/mattn/anko)
[![Coverage](https://codecov.io/gh/mattn/anko/branch/master/graph/badge.svg)](https://codecov.io/gh/mattn/anko)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattn/anko)](https://goreportcard.com/report/github.com/mattn/anko)

Anko is a scriptable interpreter written in Go.

![](https://raw.githubusercontent.com/mattn/anko/master/anko.png)

(Picture licensed under CC BY-SA 3.0, photo by Ocdp)


## Usage Example - Embedded

```go
package main

import (
	"fmt"
	"log"

	"github.com/mattn/anko/vm"
)

func main() {
	env := vm.NewEnv()

	err := env.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
println("Hello World :)")
`

	_, err = env.Execute(script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}

	// output: Hello World :)
}
```

More examples are located in the GoDoc:

https://godoc.org/github.com/mattn/anko/vm


## Usage Example - Command Line

### Building
```
go get github.com/mattn/anko
go install github.com/mattn/anko
```

### Running an Anko script file named script.ank
```
./anko script.ank
```

## Anko Script Quick Start
```
// declare variables
x = 1
y = x + 1

// print using outside the script defined println function
println(x + y) // 3

// if else statement
if x < 1 || y < 1 {
	println(x)
} else if x < 1 && y < 1 {
	println(y)
} else {
	println(x + y)
}

// array
a = [1, 2, 3]
println(a) // [1 2 3]
println(a[0]) // 1

// map
a = {"x": 1}
println(a) // map[x:1]
a.b = 2
a["c"] = 3
println(a["b"]) // 2
println(a.c) // 3

// function
func a (x) {
	println(x + 1)
}
a(3) // 4
```


## Please note that the master branch is not stable

The master branch language and API may change at any time.

To mitigate breaking changes, please use tagged branches. New tagged branches will be created for breaking changes.


## Author

Yasuhiro Matsumoto (a.k.a mattn)
