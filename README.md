# Anko

[![GoDoc Reference](https://godoc.org/github.com/mattn/anko/vm?status.svg)](http://godoc.org/github.com/mattn/anko/vm)
[![Build Status](https://travis-ci.org/mattn/anko.svg?branch=master)](https://travis-ci.org/mattn/anko)
[![Financial Contributors on Open Collective](https://opencollective.com/mattn-anko/all/badge.svg?label=financial+contributors)](https://opencollective.com/mattn-anko) [![Coverage](https://codecov.io/gh/mattn/anko/branch/master/graph/badge.svg)](https://codecov.io/gh/mattn/anko)
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

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func main() {
	e := env.NewEnv()

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
println("Hello World :)")
`

	_, err = vm.Execute(e, nil, script)
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

// slice
a = []interface{1, 2, 3}
println(a) // [1 2 3]
println(a[0]) // 1

// map
a = map[interface]interface{"x": 1}
println(a) // map[x:1]
a.b = 2
a["c"] = 3
println(a["b"]) // 2
println(a.c) // 3

// struct
a = make(struct {
	A int64,
	B float64
})
a.A = 4
a.B = 5.5
println(a.A) // 4
println(a.B) // 5.5

// function
func a (x) {
	println(x + 1)
}
a(5) // 6
```


## Please note that the master branch is not stable

The master branch language and API may change at any time.

To mitigate breaking changes, please use tagged branches. New tagged branches will be created for breaking changes.


## Author

Yasuhiro Matsumoto (a.k.a mattn)

## Contributors

### Code Contributors

This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="https://github.com/mattn/anko/graphs/contributors"><img src="https://opencollective.com/mattn-anko/contributors.svg?width=890&button=false" /></a>

### Financial Contributors

Become a financial contributor and help us sustain our community. [[Contribute](https://opencollective.com/mattn-anko/contribute)]

#### Individuals

<a href="https://opencollective.com/mattn-anko"><img src="https://opencollective.com/mattn-anko/individuals.svg?width=890"></a>

#### Organizations

Support this project with your organization. Your logo will show up here with a link to your website. [[Contribute](https://opencollective.com/mattn-anko/contribute)]

<a href="https://opencollective.com/mattn-anko/organization/0/website"><img src="https://opencollective.com/mattn-anko/organization/0/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/1/website"><img src="https://opencollective.com/mattn-anko/organization/1/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/2/website"><img src="https://opencollective.com/mattn-anko/organization/2/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/3/website"><img src="https://opencollective.com/mattn-anko/organization/3/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/4/website"><img src="https://opencollective.com/mattn-anko/organization/4/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/5/website"><img src="https://opencollective.com/mattn-anko/organization/5/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/6/website"><img src="https://opencollective.com/mattn-anko/organization/6/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/7/website"><img src="https://opencollective.com/mattn-anko/organization/7/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/8/website"><img src="https://opencollective.com/mattn-anko/organization/8/avatar.svg"></a>
<a href="https://opencollective.com/mattn-anko/organization/9/website"><img src="https://opencollective.com/mattn-anko/organization/9/avatar.svg"></a>
