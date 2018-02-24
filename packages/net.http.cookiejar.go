package packages

import (
	"net/http/cookiejar"
)

func init() {
	Packages["net/http/cookiejar"] = map[string]interface{}{
		"New": cookiejar.New,
	}
	PackageTypes["net/http/cookiejar"] = map[string]interface{}{
		"Options": cookiejar.Options{},
	}
}
