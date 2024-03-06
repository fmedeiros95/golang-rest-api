package routes

import (
	"fmedeiros95/golang-rest-api/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// ** Handle App info
	r.HandleFunc("/", handlers.Index).Methods("GET")

	// ** Setup modules routes
	SetupAuthRoutes(r)
	SetupUserRoutes(r)
	return r
}
