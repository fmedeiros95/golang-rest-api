package routes

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/handlers"
	"fmedeiros95/golang-rest-api/app/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *core.Database) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// ** Handle App info
	r.HandleFunc("/", handlers.Index).Methods("GET")

	// ** Setup Auth routes
	authHandler := handlers.NewAuthHandler(db)
	authRoutes := r.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/login", authHandler.AuthLogin).Methods("POST")
	authRoutes.HandleFunc("/register", authHandler.AuthRegister).Methods("POST")

	// ** Setup Users routes
	userHandler := handlers.NewUserHandler(db)
	userRoutes := r.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/", userHandler.CreateUser).Methods("POST")
	userRoutes.HandleFunc("/", userHandler.ListUsers).Methods("GET")
	userRoutes.HandleFunc("/me", userHandler.AuthedUser).Methods("GET")
	userRoutes.HandleFunc("/{id}", userHandler.FindUser).Methods("GET")
	userRoutes.HandleFunc("/{id}", userHandler.UpdateUser).Methods("PATCH")
	userRoutes.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")

	// ** Apply auth middleware to user routes
	authMiddleware := middlewares.NewAuthMiddleware(db)
	userRoutes.Use(authMiddleware.JWTAuthenticator)

	return r
}
