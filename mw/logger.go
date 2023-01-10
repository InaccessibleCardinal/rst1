package mw

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v, %v\n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
