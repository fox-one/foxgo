package foxerr

import (
	"fmt"
)

type baseError struct {
	code int
	msg  string

	// error hint, used to debug
	hint string
}

func (e baseError) Code() int {
	return e.code
}

func (e baseError) Message() string {
	return e.msg
}

func (e baseError) Error() string {
	msg := fmt.Sprintf("%d: %s", e.code, e.msg)

	if len(e.hint) > 0 {
		msg = fmt.Sprintf("%s (%s)", msg, e.hint)
	}

	return msg
}

func (e baseError) String() string {
	return e.Error()
}

func newBaseError(code int, msgs ...string) *baseError {
	b := &baseError{
		code: code,
	}

	switch len(msgs) {
	default:
		b.hint = msgs[1]
		fallthrough
	case 1:
		b.msg = msgs[0]
	case 0:
	}

	return b
}

var _ Error = &baseError{}

type foxError Error

type requestError struct {
	foxError

	statusCode int
}

func (r requestError) Error() string {
	return fmt.Sprintf("status code : %d, %s", r.statusCode, r.foxError.Error())
}

func (r requestError) StatusCode() int {
	return r.statusCode
}

func (r requestError) String() string {
	return r.Error()
}

func newRequestError(err Error, statusCode int) *requestError {
	return &requestError{
		foxError:   err,
		statusCode: statusCode,
	}
}

var _ RequestFailure = &requestError{}
