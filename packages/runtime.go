// +build go1.6

package packages

import (
	"reflect"
	"runtime"
)

func init() {
	if _, ok := Packages["runtime"]; !ok {
		Packages["runtime"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["runtime"]; !ok {
		PackageTypes["runtime"] = make(map[string]interface{})
	}
	var blockprofilerecord runtime.BlockProfileRecord
	var error_ runtime.Error
	var func_ runtime.Func
	var memprofilerecord runtime.MemProfileRecord
	var memstats runtime.MemStats
	var stackrecord runtime.StackRecord
	var typeassertionerror runtime.TypeAssertionError
	Packages["runtime"]["BlockProfile"] = runtime.BlockProfile
	Packages["runtime"]["Breakpoint"] = runtime.Breakpoint
	Packages["runtime"]["CPUProfile"] = runtime.CPUProfile
	Packages["runtime"]["Caller"] = runtime.Caller
	Packages["runtime"]["Callers"] = runtime.Callers
	Packages["runtime"]["Compiler"] = runtime.Compiler
	Packages["runtime"]["FuncForPC"] = runtime.FuncForPC
	Packages["runtime"]["GC"] = runtime.GC
	Packages["runtime"]["GOARCH"] = runtime.GOARCH
	Packages["runtime"]["GOMAXPROCS"] = runtime.GOMAXPROCS
	Packages["runtime"]["GOOS"] = runtime.GOOS
	Packages["runtime"]["GOROOT"] = runtime.GOROOT
	Packages["runtime"]["Goexit"] = runtime.Goexit
	Packages["runtime"]["GoroutineProfile"] = runtime.GoroutineProfile
	Packages["runtime"]["Gosched"] = runtime.Gosched
	Packages["runtime"]["LockOSThread"] = runtime.LockOSThread
	Packages["runtime"]["MemProfile"] = runtime.MemProfile
	Packages["runtime"]["MemProfileRate"] = runtime.MemProfileRate
	Packages["runtime"]["NumCPU"] = runtime.NumCPU
	Packages["runtime"]["NumCgoCall"] = runtime.NumCgoCall
	Packages["runtime"]["NumGoroutine"] = runtime.NumGoroutine
	Packages["runtime"]["ReadMemStats"] = runtime.ReadMemStats
	Packages["runtime"]["ReadTrace"] = runtime.ReadTrace
	Packages["runtime"]["SetBlockProfileRate"] = runtime.SetBlockProfileRate
	Packages["runtime"]["SetCPUProfileRate"] = runtime.SetCPUProfileRate
	Packages["runtime"]["SetFinalizer"] = runtime.SetFinalizer
	Packages["runtime"]["Stack"] = runtime.Stack
	Packages["runtime"]["StartTrace"] = runtime.StartTrace
	Packages["runtime"]["StopTrace"] = runtime.StopTrace
	Packages["runtime"]["ThreadCreateProfile"] = runtime.ThreadCreateProfile
	Packages["runtime"]["UnlockOSThread"] = runtime.UnlockOSThread
	Packages["runtime"]["Version"] = runtime.Version
	PackageTypes["runtime"]["BlockProfileRecord"] = reflect.TypeOf(&blockprofilerecord).Elem()
	PackageTypes["runtime"]["Error"] = reflect.TypeOf(&error_).Elem()
	PackageTypes["runtime"]["Func"] = reflect.TypeOf(&func_).Elem()
	PackageTypes["runtime"]["MemProfileRecord"] = reflect.TypeOf(&memprofilerecord).Elem()
	PackageTypes["runtime"]["MemStats"] = reflect.TypeOf(&memstats).Elem()
	PackageTypes["runtime"]["StackRecord"] = reflect.TypeOf(&stackrecord).Elem()
	PackageTypes["runtime"]["TypeAssertionError"] = reflect.TypeOf(&typeassertionerror).Elem()
}
