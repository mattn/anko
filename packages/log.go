// +build go1.6

package packages

import (
	"log"
	"reflect"
)

func init() {
	if _, ok := Packages["log"]; !ok {
		Packages["log"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["log"]; !ok {
		PackageTypes["log"] = make(map[string]interface{})
	}
	var logger log.Logger
	Packages["log"]["Fatal"] = log.Fatal
	Packages["log"]["Fatalf"] = log.Fatalf
	Packages["log"]["Fatalln"] = log.Fatalln
	Packages["log"]["Flags"] = log.Flags
	Packages["log"]["LUTC"] = log.LUTC
	Packages["log"]["Ldate"] = log.Ldate
	Packages["log"]["Llongfile"] = log.Llongfile
	Packages["log"]["Lmicroseconds"] = log.Lmicroseconds
	Packages["log"]["Lshortfile"] = log.Lshortfile
	Packages["log"]["LstdFlags"] = log.LstdFlags
	Packages["log"]["Ltime"] = log.Ltime
	Packages["log"]["New"] = log.New
	Packages["log"]["Output"] = log.Output
	Packages["log"]["Panic"] = log.Panic
	Packages["log"]["Panicf"] = log.Panicf
	Packages["log"]["Panicln"] = log.Panicln
	Packages["log"]["Prefix"] = log.Prefix
	Packages["log"]["Print"] = log.Print
	Packages["log"]["Printf"] = log.Printf
	Packages["log"]["Println"] = log.Println
	Packages["log"]["SetFlags"] = log.SetFlags
	Packages["log"]["SetOutput"] = log.SetOutput
	Packages["log"]["SetPrefix"] = log.SetPrefix
	PackageTypes["log"]["Logger"] = reflect.TypeOf(&logger).Elem()
}
