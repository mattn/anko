// +build go1.6

package packages

import (
	"os/signal"
)

func init() {
	if _, ok := Packages["os/signal"]; !ok {
		Packages["os/signal"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["os/signal"]; !ok {
		PackageTypes["os/signal"] = make(map[string]interface{})
	}
	Packages["os/signal"]["Ignore"] = signal.Ignore
	Packages["os/signal"]["Notify"] = signal.Notify
	Packages["os/signal"]["Reset"] = signal.Reset
	Packages["os/signal"]["Stop"] = signal.Stop
}
