package account

import (
	"github.com/fox-one/foxgo/foxerr"
)

const (
	ErrorCodeLoginRequired = 1537
	ErrorCodePinInvalid    = 1554
)

func IsLoginRequired(err error) bool {
	return foxerr.MatchCode(err, ErrorCodeLoginRequired)
}

func IsPinInvalid(err error) bool {
	return foxerr.MatchCode(err, ErrorCodePinInvalid)
}
