// +build go1.8

package packages

import (
	"time"
)

func timeGo18() {
	Packages["time"]["Until"] = time.Until
}
