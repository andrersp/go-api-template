package responder

import (
	"encoding/json"
	"net/http"
)

func RespondWithJson(status int, w http.ResponseWriter, payload interface{}) {
	// response, _ := json.Marshal(payload)

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
