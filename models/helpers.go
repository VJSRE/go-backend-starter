package models

import "github.com/go-playground/validator/v10"

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "gt":
		return "The value of " + err.Field() + " must be greater than " + err.Param()

	case "gte":
		return "The value of " + err.Field() + " must be greater than or equals " + err.Param()

	case "email":
		return "the email is invalid"
		// return the error message if the field's length is not matched to the minimum value
	case "min":
		return "the minimum length of " + err.Field() + " is equals " + err.Param()

	default:
		return "Validation error in " + err.Field()
	}

}
