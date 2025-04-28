package validator

import (
	"errors"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
	serviceError "github.com/ryanadiputraa/tetra/internal/errors"
)

const DateStringFormat = "2006-01-02"

type Validator interface {
	Validate(val any) (map[string]string, error)
}

type validation struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	v := validator.New()
	v.RegisterValidation("isodate", isISODate)
	v.RegisterValidation("date", isValidDate)
	return &validation{
		validator: v,
	}
}

func (v *validation) Validate(val any) (errDetails map[string]string, err error) {
	err = v.validator.Struct(val)
	errDetails = make(map[string]string)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			field := fieldToSnakeCase(fieldErr.Field())
			errDetails[field] = FieldErrMsg(fieldErr)
		}
		err = errors.New("bad_request")
		return
	}
	return
}

func FieldErrMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return serviceError.RequiredField
	case "max":
		return serviceError.MaxLengthField
	case "min":
		return serviceError.MinLengthField
	case "gt":
		return serviceError.MaxLengthField
	case "gte":
		return serviceError.MaxLengthField
	case "email":
		return serviceError.EmailField
	case "http_url":
		return serviceError.URLField
	case "isodate":
		return serviceError.DateField
	case "date":
		return serviceError.DateField
	default:
		return err.Error()
	}
}

func fieldToSnakeCase(input string) string {
	prev := rune(0)
	var result []rune
	for i, char := range input {
		if i > 0 && unicode.IsUpper(char) && unicode.IsLower(prev) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
		prev = char
	}
	return string(result)
}

func isISODate(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339Nano, fl.Field().String())
	return err == nil
}

func isValidDate(fl validator.FieldLevel) bool {
	_, err := time.Parse(DateStringFormat, fl.Field().String())
	return err == nil
}
