// +build go1.7

package packages

import (
	"net/http"
)

func init() {
	if _, ok := Packages["net/http"]; !ok {
		Packages["net/http"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/http"]; !ok {
		PackageTypes["net/http"] = make(map[string]interface{})
	}
	Packages["net/http"]["ErrUseLastResponse"] = http.ErrUseLastResponse
	Packages["net/http"]["LocalAddrContextKey"] = http.LocalAddrContextKey
	Packages["net/http"]["ServerContextKey"] = http.ServerContextKey
	Packages["net/http"]["StatusAlreadyReported"] = http.StatusAlreadyReported
	Packages["net/http"]["StatusFailedDependency"] = http.StatusFailedDependency
	Packages["net/http"]["StatusIMUsed"] = http.StatusIMUsed
	Packages["net/http"]["StatusInsufficientStorage"] = http.StatusInsufficientStorage
	Packages["net/http"]["StatusLocked"] = http.StatusLocked
	Packages["net/http"]["StatusLoopDetected"] = http.StatusLoopDetected
	Packages["net/http"]["StatusMultiStatus"] = http.StatusMultiStatus
	Packages["net/http"]["StatusNotExtended"] = http.StatusNotExtended
	Packages["net/http"]["StatusPermanentRedirect"] = http.StatusPermanentRedirect
	Packages["net/http"]["StatusProcessing"] = http.StatusProcessing
	Packages["net/http"]["StatusUnprocessableEntity"] = http.StatusUnprocessableEntity
	Packages["net/http"]["StatusUpgradeRequired"] = http.StatusUpgradeRequired
	Packages["net/http"]["StatusVariantAlsoNegotiates"] = http.StatusVariantAlsoNegotiates
}
