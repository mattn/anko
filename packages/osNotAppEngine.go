// +build !appengine

package packages

import (
	"os"
)

func osNotAppEngine() {
	Packages["os"]["Getppid"] = os.Getppid
}
