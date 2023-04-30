package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func SuccessResponder(status int, w http.ResponseWriter, payload interface{}) {
	w.WriteHeader(status)

	if status == 204 {
		return
	}
	if payload == nil {
		successResponse := SuccessResponse{
			Success: true,
		}
		json.NewEncoder(w).Encode(successResponse)
		return

	}
	json.NewEncoder(w).Encode(payload)
}

func ErrorResponder(status int, w http.ResponseWriter, err error) {

	errorResponse := ErrorResponse{
		Success: false,
		Msg:     err.Error(),
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(errorResponse)

}
