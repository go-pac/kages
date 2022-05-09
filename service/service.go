package service

import (
	"github.com/go-playground/validator/v10"

	"github.com/go-pac/kages/model"
)

const (
	ErrNotAllowed = "not allowed"
	ErrNotFound   = "not found"
)

var (
	validate = validator.New()
)

func ValidateStruct(iStruct interface{}) []*model.ValidationError {
	if err := validate.Struct(iStruct); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err) // todo: improve
		}

		return model.GetValidationErrors(err)
	}

	return nil
}
