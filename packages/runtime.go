package packages

import (
	"runtime"
)

func init() {
	Packages["runtime"] = map[string]interface{}{
		"GC":         runtime.GC,
		"GOARCH":     runtime.GOARCH,
		"GOMAXPROCS": runtime.GOMAXPROCS,
		"GOOS":       runtime.GOOS,
		"GOROOT":     runtime.GOROOT,
		"Version":    runtime.Version,
	}
}
