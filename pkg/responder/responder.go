package responder

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(status int, w http.ResponseWriter, payload interface{}) {

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
