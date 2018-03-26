// +build go1.6

package packages

import (
	"fmt"
	"reflect"
)

func init() {
	if _, ok := Packages["fmt"]; !ok {
		Packages["fmt"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["fmt"]; !ok {
		PackageTypes["fmt"] = make(map[string]interface{})
	}
	var formatter fmt.Formatter
	var gostringer fmt.GoStringer
	var scanstate fmt.ScanState
	var scanner fmt.Scanner
	var state fmt.State
	var stringer fmt.Stringer
	Packages["fmt"]["Errorf"] = fmt.Errorf
	Packages["fmt"]["Fprint"] = fmt.Fprint
	Packages["fmt"]["Fprintf"] = fmt.Fprintf
	Packages["fmt"]["Fprintln"] = fmt.Fprintln
	Packages["fmt"]["Fscan"] = fmt.Fscan
	Packages["fmt"]["Fscanf"] = fmt.Fscanf
	Packages["fmt"]["Fscanln"] = fmt.Fscanln
	Packages["fmt"]["Print"] = fmt.Print
	Packages["fmt"]["Printf"] = fmt.Printf
	Packages["fmt"]["Println"] = fmt.Println
	Packages["fmt"]["Scan"] = fmt.Scan
	Packages["fmt"]["Scanf"] = fmt.Scanf
	Packages["fmt"]["Scanln"] = fmt.Scanln
	Packages["fmt"]["Sprint"] = fmt.Sprint
	Packages["fmt"]["Sprintf"] = fmt.Sprintf
	Packages["fmt"]["Sprintln"] = fmt.Sprintln
	Packages["fmt"]["Sscan"] = fmt.Sscan
	Packages["fmt"]["Sscanf"] = fmt.Sscanf
	Packages["fmt"]["Sscanln"] = fmt.Sscanln
	PackageTypes["fmt"]["Formatter"] = reflect.TypeOf(&formatter).Elem()
	PackageTypes["fmt"]["GoStringer"] = reflect.TypeOf(&gostringer).Elem()
	PackageTypes["fmt"]["ScanState"] = reflect.TypeOf(&scanstate).Elem()
	PackageTypes["fmt"]["Scanner"] = reflect.TypeOf(&scanner).Elem()
	PackageTypes["fmt"]["State"] = reflect.TypeOf(&state).Elem()
	PackageTypes["fmt"]["Stringer"] = reflect.TypeOf(&stringer).Elem()
}
