package app

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(config *core.EnvConfig) {
	log.Print("🔥 Starting application...")

	log.Print("🎲 Connecting to database...")
	db, err := core.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("🚩 Setup routes...")
	a.Router = routes.SetupRoutes(db)
}

func (a *App) Run() {
	server := &http.Server{
		Addr:         "0.0.0.0:" + core.Config.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router,
	}
	log.Print("🚀 Application started on: http://" + server.Addr)
	log.Fatal(server.ListenAndServe())
}
