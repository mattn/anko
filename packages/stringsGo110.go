// +build go1.10

package packages

import (
	"strings"
)

func stringsGo110() {
	PackageTypes["strings"] = map[string]interface{}{
		"Builder": strings.Builder{},
	}
}
