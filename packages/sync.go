package packages

import (
	"sync"
)

func init() {
	Packages["sync"] = map[string]interface{}{
		"NewCond": sync.NewCond,
	}
	PackageTypes["sync"] = map[string]interface{}{
		"Cond":      sync.Cond{},
		"Mutex":     sync.Mutex{},
		"Once":      sync.Once{},
		"Pool":      sync.Pool{},
		"RWMutex":   sync.RWMutex{},
		"WaitGroup": sync.WaitGroup{},
	}
	syncGo19()
}
