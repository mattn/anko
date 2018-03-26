// +build !appengine

package packages

import (
	"os"
)

func notAppEngine() {
	Packages["os"]["Getppid"] = os.Getppid
}
