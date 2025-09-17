package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrorType struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func ValidateStruct(s any) []*ErrorType {
	var errors []*ErrorType

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var message string

			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", err.Field())
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
			case "max":
				message = fmt.Sprintf("%s cannot be longer than %s characters", err.Field(), err.Param())
			case "email":
				message = fmt.Sprintf("%s must be a valid email address", err.Field())
			default:
				message = fmt.Sprintf("%s is not valid", err.Field())
			}

			errors = append(errors, &ErrorType{
				Field:   err.Field(),
				Message: message,
			})
		}
	}

	return errors
}
