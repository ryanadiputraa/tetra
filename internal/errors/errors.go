package errors

import "net/http"

const (
	// Error Code
	BadRequest   = "bad_request"
	Unauthorized = "unauthorized"
	Forbidden    = "forbidden"
	NotFound     = "not_found"
	ServerError  = "internal_server_error"

	// Error Msg Code
	EmailTaken        = "email_taken"
	MissingAuthHeader = "missing_auth_header"
	InvalidAuthHeader = "invalid_auth_header"
)

var HttpErrMap = map[string]int{
	BadRequest:   http.StatusBadRequest,
	Unauthorized: http.StatusUnauthorized,
	Forbidden:    http.StatusForbidden,
	NotFound:     http.StatusNotFound,
	ServerError:  http.StatusInternalServerError,
}

type ServiceErr struct {
	ErrCode string
	msg     string
}

func NewServiceErr(errCode, msg string) error {
	return &ServiceErr{
		ErrCode: errCode,
		msg:     msg,
	}
}

func (e *ServiceErr) Error() string {
	return e.msg
}
