// +build go1.7

package packages

import (
	"bytes"
)

func bytesGo17() {
	Packages["bytes"]["ContainsRune"] = bytes.ContainsRune
}
