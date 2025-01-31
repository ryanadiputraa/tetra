package validator

import (
	"errors"
	"fmt"
	"time"
	"unicode"

	"github.com/go-playground/validator/v10"
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
		return fmt.Sprintf("is required")
	case "max":
		return fmt.Sprintf("should have a maximum length of %s", err.Param())
	case "min":
		return fmt.Sprintf("should have a minimum length of %s", err.Param())
	case "gt":
		return fmt.Sprintf("should be greater than %s", err.Param())
	case "gte":
		return fmt.Sprintf("should be greater or equal than %s", err.Param())
	case "email":
		return fmt.Sprintf("should be a valid email address")
	case "http_url":
		return fmt.Sprintf("should be a valid http url")
	case "isodate":
		return fmt.Sprintf("should be a valid ISO date format")
	case "date":
		return fmt.Sprintf("should be a valid date format 'YYYY-MM-DD'")
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
