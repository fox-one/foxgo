package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoRequest(t *testing.T) {
	param := NewQueryParam()
	param.SetValue("type", 2)

	req, err := New("GET", "/index/gbi?type=1", param)
	if assert.Nil(t, err) {
		r, err := req.request()
		if assert.Nil(t, err) {
			assert.Empty(t, r.URL.String())
		}
	}
}
