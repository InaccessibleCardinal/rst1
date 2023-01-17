package ctrl

import (
	"net/http"
)

func HandleWindows(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusNotFound
	response := Response{Message: "hello windows"}
	SendResponse(w, response, statusCode)
}