package main

import (
	"fmedeiros95/golang-rest-api/app"
	"fmedeiros95/golang-rest-api/app/core"
	"log"
)

func main() {
	log.Print("⌛ Loading enviroment variables...")
	core.LoadEnv(".env")

	// ** Initialize app
	app := &app.App{}
	app.Initialize(&core.Config)
	app.Run()
}
