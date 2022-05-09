package model

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

func GetValidationErrors(srcErr error) []*ValidationError {
	var errors []*ValidationError

	for _, err := range srcErr.(validator.ValidationErrors) {
		var element ValidationError
		element.FailedField = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors, &element)
	}

	return errors
}

func (e *ValidationError) Error() string {
	return "" // todo
}
