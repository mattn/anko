package packages

import (
	"errors"
)

func init() {
	Packages["errors"] = map[string]interface{}{
		"New": errors.New,
	}
}
