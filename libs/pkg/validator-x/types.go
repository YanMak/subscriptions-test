package validatorХ

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ValidatorXModuleDeps struct{}

type ValidatorXModule struct {
	*ValidatorX
}

func NewValidatorXModule(deps ValidatorXModuleDeps) *ValidatorXModule {
	vld := NewValidatorX()
	return &ValidatorXModule{ValidatorX: vld}
}

type ValidatorXModuleConstructor func(ValidatorXModuleDeps) *ValidatorXModule

type ValidatorX struct {
	*validator.Validate
}

func NewValidatorX() *ValidatorX {
	validate := validator.New()

	_ = validate.RegisterValidation("date_format_mm_yyyy", func(fl validator.FieldLevel) bool {
		return validateDate(fl, "01-2006")
	})

	_ = validate.RegisterValidation("date_format_dd_mm_yyyy", func(fl validator.FieldLevel) bool {
		return validateDate(fl, "02-01-2006")
	})

	_ = validate.RegisterValidation("date_format_iso", func(fl validator.FieldLevel) bool {
		return validateDate(fl, time.RFC3339) // e.g., "2025-07-30T15:04:05Z"
	})

	return &ValidatorX{Validate: validate}
}
