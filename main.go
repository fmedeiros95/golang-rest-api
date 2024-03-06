package main

import (
	"fmedeiros95/golang-rest-api/core"
	"fmedeiros95/golang-rest-api/routes"
	"log"
	"net/http"
)

func main() {
	log.Print("🔥 Starting application...")

	log.Print("⌛ Loading enviroment variables...")
	core.LoadEnv()

	log.Print("🚩 Setup routes...")
	r := routes.SetupRoutes()

	log.Print("🎲 Connecting to database...")
	_, err := core.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// ** Start listen to app
	log.Print("🚀 Application started on: http://127.0.0.1:" + core.Config.Port)
	log.Fatal(http.ListenAndServe(":"+core.Config.Port, r))
}
