package models

import (
	"github.com/go-playground/validator/v10"
)

type ItemRequest struct {
	Name    string `json:"name" validate:"required"`
	Price   int    `json:"price" validate:"required,gt=0"`
	Quality int    `json:"quality" validate:"gtd=0"`
}

func (itemInput ItemRequest) ValidateStruct() []*ErrorResponse {

	var errors []*ErrorResponse

	validate := validator.New()
	err := validate.Struct(itemInput)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Filed = err.Field()
			errors = append(errors, &element)
		}
	}

	return errors
}
