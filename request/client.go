package request

import (
	"net/http"
	"time"
)

var sharedClient *http.Client = nil

func init() {
	tr := &http.Transport{DisableKeepAlives: false}
	sharedClient = &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
}
