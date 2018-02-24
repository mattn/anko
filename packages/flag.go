package packages

import (
	"flag"
)

func init() {
	Packages["flag"] = map[string]interface{}{
		"Arg":             flag.Arg,
		"Args":            flag.Args,
		"Bool":            flag.Bool,
		"BoolVar":         flag.BoolVar,
		"CommandLine":     flag.CommandLine,
		"ContinueOnError": flag.ContinueOnError,
		"Duration":        flag.Duration,
		"DurationVar":     flag.DurationVar,
		"ErrHelp":         flag.ErrHelp,
		"ExitOnError":     flag.ExitOnError,
		"Float64":         flag.Float64,
		"Float64Var":      flag.Float64Var,
		"Int":             flag.Int,
		"Int64":           flag.Int64,
		"Int64Var":        flag.Int64Var,
		"IntVar":          flag.IntVar,
		"Lookup":          flag.Lookup,
		"NArg":            flag.NArg,
		"NFlag":           flag.NFlag,
		"NewFlagSet":      flag.NewFlagSet,
		"PanicOnError":    flag.PanicOnError,
		"Parse":           flag.Parse,
		"Parsed":          flag.Parsed,
		"PrintDefaults":   flag.PrintDefaults,
		"Set":             flag.Set,
		"String":          flag.String,
		"StringVar":       flag.StringVar,
		"Uint":            flag.Uint,
		"Uint64":          flag.Uint64,
		"Uint64Var":       flag.Uint64Var,
		"UintVar":         flag.UintVar,
		"Usage":           flag.Usage,
		"Var":             flag.Var,
		"Visit":           flag.Visit,
		"VisitAll":        flag.VisitAll,
	}
}
