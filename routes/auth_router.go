package routes

import (
	"fmedeiros95/golang-rest-api/handlers"

	"github.com/gorilla/mux"
)

func SetupAuthRoutes(r *mux.Router) {
	usersRouter := r.PathPrefix("/auth").Subrouter()
	usersRouter.HandleFunc("/login", handlers.AuthLogin).Methods("POST")
	usersRouter.HandleFunc("/register", handlers.AuthRegister).Methods("POST")
}
