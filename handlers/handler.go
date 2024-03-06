package handlers

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	// ** Auth Handlers
	r.HandleFunc("/auth/login", AuthLogin).Methods("POST")
	r.HandleFunc("/auth/register", AuthRegister).Methods("POST")

	// ** Add middleware to next routes
	// r.Use(middlewares.JWTAuthenticator)

	// ** Get authed user
	r.HandleFunc("/auth/me", AuthGetAuthedUser).Methods("GET")

	// ** Users Handlers
	r.HandleFunc("/users", UsersList).Methods("GET")
	r.HandleFunc("/users", UsersCreate).Methods("POST")
	r.HandleFunc("/users/{userId}", UsersFind).Methods("GET")
	r.HandleFunc("/users/{userId}", UsersUpdate).Methods("PATCH")
	r.HandleFunc("/users/{userId}", UsersDelete).Methods("DELETE")
}
