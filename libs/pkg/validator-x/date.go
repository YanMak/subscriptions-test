package validatorХ

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func validateDate(fl validator.FieldLevel, layout string) bool {
	switch val := fl.Field().Interface().(type) {
	case string:
		if val == "" {
			return true
		}
		_, err := time.Parse(layout, val)
		return err == nil
	case *string:
		if val == nil || *val == "" {
			return true
		}
		_, err := time.Parse(layout, *val)
		return err == nil
	default:
		return false
	}
}
