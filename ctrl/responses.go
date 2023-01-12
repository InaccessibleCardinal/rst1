package ctrl

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, response any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
