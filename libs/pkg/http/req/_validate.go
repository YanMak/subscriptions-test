package req

import (
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validate := validator.New()
	validate.RegisterValidation("date_format", validateDateFormat)
	err := validate.Struct(payload)
	return err
}
