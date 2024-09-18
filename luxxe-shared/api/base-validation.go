package api

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
	Param       string
}

type XValidator struct {
	validator *validator.Validate
}

var validate = validator.New()
var BaseValidator = &XValidator{
	validator: validate,
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true
			elem.Param = err.Param()

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func ValidateAPIData(data interface{}) (bool, *fiber.Error) {
	if errs := BaseValidator.Validate(data); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s' as %s",
				err.FailedField,
				err.Value,
				err.Tag,
				err.Param,
			))
		}

		return false, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}
	return true, nil
}

