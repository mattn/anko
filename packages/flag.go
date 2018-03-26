// +build go1.6

package packages

import (
	"flag"
	"reflect"
)

func init() {
	if _, ok := Packages["flag"]; !ok {
		Packages["flag"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["flag"]; !ok {
		PackageTypes["flag"] = make(map[string]interface{})
	}
	var errorhandling flag.ErrorHandling
	var flag_ flag.Flag
	var flagset flag.FlagSet
	var getter flag.Getter
	var value flag.Value
	Packages["flag"]["Arg"] = flag.Arg
	Packages["flag"]["Args"] = flag.Args
	Packages["flag"]["Bool"] = flag.Bool
	Packages["flag"]["BoolVar"] = flag.BoolVar
	Packages["flag"]["CommandLine"] = flag.CommandLine
	Packages["flag"]["ContinueOnError"] = flag.ContinueOnError
	Packages["flag"]["Duration"] = flag.Duration
	Packages["flag"]["DurationVar"] = flag.DurationVar
	Packages["flag"]["ErrHelp"] = flag.ErrHelp
	Packages["flag"]["ExitOnError"] = flag.ExitOnError
	Packages["flag"]["Float64"] = flag.Float64
	Packages["flag"]["Float64Var"] = flag.Float64Var
	Packages["flag"]["Int"] = flag.Int
	Packages["flag"]["Int64"] = flag.Int64
	Packages["flag"]["Int64Var"] = flag.Int64Var
	Packages["flag"]["IntVar"] = flag.IntVar
	Packages["flag"]["Lookup"] = flag.Lookup
	Packages["flag"]["NArg"] = flag.NArg
	Packages["flag"]["NFlag"] = flag.NFlag
	Packages["flag"]["NewFlagSet"] = flag.NewFlagSet
	Packages["flag"]["PanicOnError"] = flag.PanicOnError
	Packages["flag"]["Parse"] = flag.Parse
	Packages["flag"]["Parsed"] = flag.Parsed
	Packages["flag"]["PrintDefaults"] = flag.PrintDefaults
	Packages["flag"]["Set"] = flag.Set
	Packages["flag"]["String"] = flag.String
	Packages["flag"]["StringVar"] = flag.StringVar
	Packages["flag"]["Uint"] = flag.Uint
	Packages["flag"]["Uint64"] = flag.Uint64
	Packages["flag"]["Uint64Var"] = flag.Uint64Var
	Packages["flag"]["UintVar"] = flag.UintVar
	Packages["flag"]["UnquoteUsage"] = flag.UnquoteUsage
	Packages["flag"]["Usage"] = flag.Usage
	Packages["flag"]["Var"] = flag.Var
	Packages["flag"]["Visit"] = flag.Visit
	Packages["flag"]["VisitAll"] = flag.VisitAll
	PackageTypes["flag"]["ErrorHandling"] = reflect.TypeOf(&errorhandling).Elem()
	PackageTypes["flag"]["Flag"] = reflect.TypeOf(&flag_).Elem()
	PackageTypes["flag"]["FlagSet"] = reflect.TypeOf(&flagset).Elem()
	PackageTypes["flag"]["Getter"] = reflect.TypeOf(&getter).Elem()
	PackageTypes["flag"]["Value"] = reflect.TypeOf(&value).Elem()
}
