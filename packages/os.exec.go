package packages

import (
	"os/exec"
)

func init() {
	Packages["os/exec"] = map[string]interface{}{
		"ErrNotFound": exec.ErrNotFound,
		"LookPath":    exec.LookPath,
		"Command":     exec.Command,
	}
}
