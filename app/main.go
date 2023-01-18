package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"rst1/mw"
	"rst1/route"
	"rst1/xl"
)

func main() {
	// xl.ReadXL()
	xl.UnzipXL()
	router := mux.NewRouter()
	router.Use(mw.LoggingMiddleware)
	route.InitRoutes(router)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
