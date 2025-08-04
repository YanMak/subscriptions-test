package req

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func validateDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	_, err := time.Parse("01-2006", dateStr)

	return err == nil
}
