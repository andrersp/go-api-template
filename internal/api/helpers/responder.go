package helpers

import (
	"encoding/json"
	"net/http"
)

func Responder(status int, w http.ResponseWriter, payload interface{}) {

	if status == 204 || payload == nil {
		w.WriteHeader(status)
		return
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
