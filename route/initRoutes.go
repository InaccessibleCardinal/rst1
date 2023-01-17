package route

import (
	"rst1/ctrl"
	"rst1/db"
	"rst1/ht"
	"rst1/svc"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/patrickmn/go-cache"
)

var (
	conn              = db.GetDb()
	userRepo          = db.NewUserRepository(conn)
	userService       = svc.NewUserService(userRepo)
	remoteUserService = svc.NewRemoteUserService(
		ht.MakeHttpRequest,
		userService,
		cache.New(2*time.Minute, 5*time.Minute),
	)
)

func createUserController() *ctrl.UserController {
	return ctrl.NewUserController(userService)
}

func createRemoteUserController() *ctrl.RemoteUserController {
	return ctrl.NewRemoteUserController(remoteUserService)
}

func InitRoutes(router *mux.Router) {
	userController := createUserController()
	ruController := createRemoteUserController()
	router.HandleFunc("/users", userController.HandleGetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.HandleGetUser).Methods("GET")
	router.HandleFunc("/users/{id}", userController.HandleUpdateUser).Methods("PUT")
	router.HandleFunc("/address", ruController.HandleGetAddress).Methods("GET")
}
