package app

import (
	"encoding/json"
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Fiber *fiber.App
}

func (a *App) Initialize(config *core.EnvConfig) {
	log.Print("🔥 Starting application...")

	log.Print("🎲 Connecting to database...")
	db, err := core.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("🔧 Setting up Fiber...")
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "GoLang Rest API",
		AppName:       "GoLang Rest API v0.0.1",
	})

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	type AppMetadata struct {
	// 		Name    string `json:"name"`
	// 		Version string `json:"version"`
	// 		Author  string `json:"author"`
	// 	}
	// 	metadata := AppMetadata{
	// 		Name:    "GoLang Rest API",
	// 		Version: "0.0.1",
	// 		Author:  "Felipe Medeiros <medeiros.dev@gmail.com>",
	// 	}
	// 	return core.RespondWithJSON(c, fiber.StatusOK, metadata, "")
	// }).Name("index")
	// app.Get("/routes", func(c *fiber.Ctx) error {
	// 	return core.RespondWithJSON(c, fiber.StatusOK, app.Stack(), "")
	// }).Name("routes")

	log.Print("🚩 Setup routes...")
	routes.SetupRoutes(app, db)

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Print(string(data))

	a.Fiber = app
}

func (a *App) Run() {
	log.Print("🚀 Application started!")
	log.Fatal(a.Fiber.Listen(":" + core.Config.Port))
}
