// +build !appengine

package main

import (
	"github.com/mattn/anko/packages"

	"github.com/daviddengcn/go-colortext"
)

var ntoc = map[string]ct.Color{
	"none":    ct.None,
	"black":   ct.Black,
	"red":     ct.Red,
	"green":   ct.Green,
	"yellow":  ct.Yellow,
	"blue":    ct.Blue,
	"mazenta": ct.Magenta,
	"cyan":    ct.Cyan,
	"white":   ct.White,
}

func colorOf(name string) ct.Color {
	if c, ok := ntoc[name]; ok {
		return c
	}
	return ct.None
}

// AddPackageColortext adds the go-colortext package and the ChangeColor function
func AddPackageColortext() {
	packages.Packages["github.com/daviddengcn/go-colortext"] = map[string]interface{}{
		"ChangeColor": func(fg string, fa bool, rest ...interface{}) {
			if len(rest) == 2 {
				bg, ok := rest[0].(string)
				if !ok {
					panic("Argument #3 should be string")
				}
				ba, ok := rest[1].(bool)
				if !ok {
					panic("Argument #4 should be string")
				}
				ct.ChangeColor(colorOf(fg), fa, colorOf(bg), ba)
			} else {
				ct.ChangeColor(colorOf(fg), fa, ct.None, false)
			}
		},
		"ResetColor": func() {
			ct.ResetColor()
		},
	}
}
