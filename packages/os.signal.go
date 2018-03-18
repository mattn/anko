package packages

import (
	"os/signal"
)

func init() {
	Packages["os/signal"] = map[string]interface{}{
		"Notify": signal.Notify,
		"Stop":   signal.Stop,
	}
}
