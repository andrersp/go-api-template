package customvalidator

import (
	"errors"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	v    *customValidator
	once sync.Once
)

func Get() *customValidator {

	once.Do(func() {
		newValidate := validator.New()

		v = &customValidator{
			Validator: newValidate,
		}

	})

	return v

}

type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {

		errorDetail := normalizeError(err)

		return errors.New(errorDetail)
	}
	return nil
}

func normalizeError(requestErr error) string {

	return validate(requestErr.(validator.ValidationErrors))

}

func validate(errors validator.ValidationErrors) (resultError string) {

	for _, err := range errors {

		switch err.Tag() {
		case "required":
			resultError = err.Field() + " required!"
			return
		case "email":
			resultError = err.Field() + " invalid!"
			return
		case "min":
			resultError = err.Field() + " must have " + err.Param() + " characters at least!"
			return
		case "containsany":
			resultError = err.Field() + " must contain at least one of these characters " + err.Param()
			return
		case "eqfield":
			resultError = err.Field() + " does not match " + err.Param()
			return
		default:
			resultError += "error on field " + err.Tag() + " " + err.Field()
			return
		}
	}

	return
}
