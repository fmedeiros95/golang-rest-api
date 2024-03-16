package handlers

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/services"
	"fmedeiros95/golang-rest-api/app/validations"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(db *core.Database) *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(db),
	}
}

func (ah *AuthHandler) AuthLogin(c *fiber.Ctx) error {
	var payload validations.Login
	if err := c.BodyParser(&payload); err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, "Invalid request")
	}

	data, err := ah.authService.Login(payload)
	if err != nil {
		return core.RespondWithError(c, fiber.StatusUnauthorized, err.Error())
	}

	return core.RespondWithJSON(c, fiber.StatusOK, data, "Login has been successful")
}

func (ah *AuthHandler) AuthRegister(c *fiber.Ctx) error {
	var payload validations.Register
	if err := c.BodyParser(&payload); err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, "Invalid request")
	}

	_, err := ah.authService.Register(payload)
	if err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, err.Error())
	}

	return core.RespondWithJSON(c, fiber.StatusCreated, nil, "User has been created")
}
