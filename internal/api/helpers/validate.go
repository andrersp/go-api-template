package helpers

import (
	"sync"

	erroresponse "github.com/andrersp/go-api-template/internal/pkg/error-response"
	"github.com/go-playground/validator/v10"
)

var (
	customValidator *CustomValidator
	once            sync.Once
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidator {

	once.Do(func() {
		newValidate := validator.New()

		customValidator = &CustomValidator{
			Validator: newValidate,
		}

	})

	return customValidator

}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {

		errorDetail := NormalizeError(err)
		err := erroresponse.NewErrorResponse("PAYLOAD_ERROR", errorDetail)
		return err
	}
	return nil
}

func NormalizeError(requestErr error) string {

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
		default:
			resultError += "error on field " + err.Tag() + " " + err.Field()
			return
		}
	}

	return
}
