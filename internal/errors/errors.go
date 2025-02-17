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
	SubscriptionEnd           = "subscription_end"
	RecordNotFound            = "record_not_found"
	EmailTaken                = "email_taken"
	MissingAuthHeader         = "missing_auth_header"
	InvalidAuthHeader         = "invalid_auth_header"
	OrganizationAlreadyExists = "organization_already_exists"
	UserHasJoinedOrg          = "user_has_joined_org"

	// Validation Err Code
	RequiredField  = "required_field"
	EmailField     = "email_field"
	MinLengthField = "min_length_field"
	MaxLengthField = "max_length_field"
	URLField       = "url_field"
	DateField      = "date_field"
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
