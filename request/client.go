package request

import (
	"net/http"
)

var sharedClient *http.Client = nil

func init() {
	tr := &http.Transport{DisableKeepAlives: false}
	sharedClient = &http.Client{
		Transport: tr,
	}
}
