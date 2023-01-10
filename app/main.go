package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"rst1/ctrl"
	"rst1/db"
	"rst1/mw"
	"rst1/svc"
)

func CreatController() *ctrl.UserController {
	conn := db.GetDb()
	userRepo := db.NewUserRepository(conn)
	userService := svc.NewUserService(userRepo)
	return ctrl.NewUserController(&userService)
}

func InitRoutes(router *mux.Router, userController *ctrl.UserController) {
	router.HandleFunc("/users", userController.HandleGetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.HandleGetUser).Methods("GET")
	router.HandleFunc("/users/{id}", userController.HandleUpdateUser).Methods("PUT")
}

func main() {
	router := mux.NewRouter()
	router.Use(mw.LoggingMiddleware)
	InitRoutes(router, CreatController())

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
