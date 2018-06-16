// +build go1.9

package packages

import (
	"sync"
)

func syncGo19() {
	PackageTypes["sync"]["Map"] = sync.Map{}
}
