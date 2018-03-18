package packages

import (
	"encoding/json"
)

func init() {
	Packages["encoding/json"] = map[string]interface{}{
		"Marshal":   json.Marshal,
		"Unmarshal": json.Unmarshal,
	}
}
