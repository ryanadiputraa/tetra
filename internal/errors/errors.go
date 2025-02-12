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
	RecordNotFound            = "record_not_found"
	EmailTaken                = "email_taken"
	MissingAuthHeader         = "missing_auth_header"
	InvalidAuthHeader         = "invalid_auth_header"
	OrganizationAlreadyExists = "organization_already_exists"
)

var HttpErrMap = map[string]int{
	BadRequest:   http.StatusBadRequest,
	Unauthorized: http.StatusUnauthorized,
	Forbidden:    http.StatusForbidden,
	NotFound:     http.StatusNotFound,
	ServerError:  http.StatusInternalServerError,
}

type Error struct {
	ErrCode string
	msg     string
}

func NewServiceErr(errCode, msg string) error {
	return &Error{
		ErrCode: errCode,
		msg:     msg,
	}
}

func (e *Error) Error() string {
	return e.msg
}
