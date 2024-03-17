package app

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		type AppMetadata struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Author  string `json:"author"`
		}
		return core.RespondWithJSON(c, fiber.StatusOK, AppMetadata{
			Name:    "GoLang Rest API",
			Version: "0.0.1",
			Author:  "Felipe Medeiros <medeiros.dev@gmail.com>",
		}, "")
	}).Name("index")

	log.Print("🚩 Setup routes...")
	routes.SetupRoutes(app, db)

	a.Fiber = app
}

func (a *App) Run() {
	log.Print("🚀 Application started!")
	log.Fatal(a.Fiber.Listen(":" + core.Config.Port))
}
