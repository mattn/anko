# anko

[![Build Status](https://travis-ci.org/mattn/anko.png?branch=master)](https://travis-ci.org/mattn/anko)

![](https://raw.githubusercontent.com/mattn/anko/master/anko.png)

(Picture licensed under CC BY-SA 3.0 by wikipedia)

Scriptable interpreter written in golang

## Usage

```
# declare function
func foo(x){
  return x + 1
}

# declare variables
x = 1
y = x + 1

# print values 
println(x * (y + 2 * x + foo(x) / 2))

# if/else condition
if foo(y) > 1 {
  println("こんにちわ世界")
} else {
  println("Hello, World")
}

# array type
a = [1,2,3]
println(a[2])
println(len(a))

# map type
m = {"foo": "bar", "bar": "baz"}
for k in keys(m) {
  println(m[k])
}
```

## Requirements

* golang

## Installation

```
$ go get github.com/mattn/anko
```

# License

MIT

# Author

Yasuhiro Matsumoto (a.k.a mattn)

