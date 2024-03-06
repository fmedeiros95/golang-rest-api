package routes

import (
	"fmedeiros95/golang-rest-api/handlers"
	"fmedeiros95/golang-rest-api/middlewares"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router) {
	usersRouter := r.PathPrefix("/users").Subrouter()

	// ** apply auth middleware
	r.Use(middlewares.JWTAuthenticator)

	// ** Authed user data
	usersRouter.HandleFunc("/me", handlers.UserAuthed).Methods("GET")

	// ** CRUD users routes
	usersRouter.HandleFunc("/", handlers.UsersCreate).Methods("POST")
	usersRouter.HandleFunc("/", handlers.UsersList).Methods("GET")
	usersRouter.HandleFunc("/{id}", handlers.UsersFind).Methods("GET")
	usersRouter.HandleFunc("/{id}", handlers.UsersUpdate).Methods("PUT")
	usersRouter.HandleFunc("/{id}", handlers.UsersDelete).Methods("DELETE")
}
