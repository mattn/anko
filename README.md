# anko

![](https://raw.githubusercontent.com/mattn/anko/master/anko.png)

Scriptable interpreter written in golang

## Usage

```
# declare function
func foo(x){
  return x + 1;
}

# declare variables
var x = 1;
var y = x + 1;

# print values 
println(x * (y + 2 * x + foo(x) / 2));

# if/else condition
if (foo(y) > 1) {
  println("こんにちわ世界");
} else {
  println("Hello, World");
}

# array type
var a = [1,2,3];
println(a[2]);
println(len(a));

# map type
var m = {"foo": "bar", "bar": "baz"};
for k in keys(m) {
  println(m[k]);
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

