package helpers

import (
	"encoding/json"
	"errors"
	"net/http"

	erroresponse "github.com/andrersp/go-api-template/internal/pkg/error-response"
)

func SuccessResponder(status int, w http.ResponseWriter, payload interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func ErrorResponder(status int, w http.ResponseWriter, err error) {

	var genericError *erroresponse.ErrorResponse

	if errors.As(err, &genericError) {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(err)
		return
	}

	genericError = erroresponse.NewErrorResponse("UNPROCESSABLE_ENTITY", "")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(genericError)

}
