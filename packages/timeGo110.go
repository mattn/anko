// +build go1.10

package packages

import (
	"time"
)

func timeGo110() {
	Packages["time"]["LoadLocationFromTZData"] = time.LoadLocationFromTZData
}
