package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	erroresponse "github.com/andrersp/go-api-template/internal/pkg/error-response"
)

func DecodeJsonBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		fmt.Println(contentType)

	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)

		default:
			return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", "")

		}
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", msg)
	}

	return nil

}
