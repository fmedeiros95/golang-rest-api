package routes

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *core.Database) {
	// ** Setup Auth routes
	authHandler := handlers.NewAuthHandler(db)
	app.Route("/auth", func(auth fiber.Router) {
		auth.Post("/login", authHandler.AuthLogin).Name("login")
		auth.Post("/register", authHandler.AuthRegister).Name("register")
	}, "auth.")

	// ** Setup Users routes
	userHandler := handlers.NewUserHandler(db)
	app.Route("/users", func(users fiber.Router) {
		users.Get("/", userHandler.ListUsers).Name("list")
		users.Post("/", userHandler.CreateUser).Name("create")
		users.Get("/:id", userHandler.FindUser).Name("find")
		users.Patch("/:id", userHandler.UpdateUser).Name("update")
		users.Delete("/:id", userHandler.DeleteUser).Name("delete")
	}, "users.")

	// ** Apply auth middleware to user routes
	// authMiddleware := middlewares.NewAuthMiddleware(db)
	// userRoutes.Use(authMiddleware.JWTAuthenticator)
}
