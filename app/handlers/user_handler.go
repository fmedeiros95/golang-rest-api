package handlers

import (
	"fmedeiros95/golang-rest-api/app/core"
	"fmedeiros95/golang-rest-api/app/models"
	"fmedeiros95/golang-rest-api/app/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(db *core.Database) *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(db),
	}
}

func (uh *UserHandler) AuthedUser(c *fiber.Ctx) error {
	// user := utils.GetUserFromContext(r.Context())
	return core.RespondWithJSON(c, fiber.StatusOK, fiber.Map{"a": "b"}, "User found successfully")
}

func (uh *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := uh.userService.ListUsers()
	if err != nil {
		return core.RespondWithError(c, fiber.StatusInternalServerError, "Error while trying to list users")
	}
	return core.RespondWithJSON(c, fiber.StatusOK, users, "Users found successfully")
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var payload models.User
	if err := c.BodyParser(&payload); err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, "Invalid request")
	}

	if err := uh.userService.CreateUser(&payload); err != nil {
		return core.RespondWithError(c, fiber.StatusInternalServerError, "Error while trying to create user")
	}

	return core.RespondWithJSON(c, fiber.StatusCreated, nil, "User created successfully")
}

func (uh *UserHandler) FindUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, "Invalid user id")
	}

	user, err := uh.userService.FindUser(uint(id))
	if err != nil {
		return core.RespondWithError(c, fiber.StatusNotFound, "User not found")
	}

	return core.RespondWithJSON(c, fiber.StatusOK, user, "User found successfully")
}

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var payload models.User
	if err := c.BodyParser(&payload); err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, "Invalid request")
	}

	if err := uh.userService.UpdateUser(payload); err != nil {
		return core.RespondWithError(c, fiber.StatusInternalServerError, "Error while trying to update user")
	}

	return core.RespondWithJSON(c, fiber.StatusOK, nil, "User updated successfully")
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return core.RespondWithError(c, fiber.StatusBadRequest, "Invalid user id")
	}

	user, err := uh.userService.FindUser(uint(id))
	if err != nil {
		return core.RespondWithError(c, fiber.StatusNotFound, "User not found")
	}

	if err := uh.userService.DeleteUser(*user); err != nil {
		return core.RespondWithError(c, fiber.StatusInternalServerError, "Error while trying to delete user")
	}

	return core.RespondWithJSON(c, fiber.StatusOK, nil, "User deleted successfully")
}
