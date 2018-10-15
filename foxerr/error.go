package foxerr

type Error interface {
	// Satisfy the generic error interface.
	error

	// return the error defined by fox.ONE
	Code() int

	// return the error detail msg
	Message() string
}

func New(code int, message ...string) Error {
	return newBaseError(code, message...)
}

type RequestFailure interface {
	Error

	// The status code of the HTTP response.
	StatusCode() int
}

func NewRequestFailure(err Error, statusCode int) RequestFailure {
	return newRequestError(err, statusCode)
}

const NotFoxError = 0

func Code(err error) int {
	if e, ok := err.(Error); ok {
		return e.Code()
	}

	return NotFoxError
}

func MatchCode(err error, code int) bool {
	return Code(err) == code
}

func IsFoxError(err error) bool {
	return MatchCode(err, NotFoxError)
}

func StatusCode(err error) int {
	if e, ok := err.(RequestFailure); ok {
		return e.StatusCode()
	}

	return 0
}
