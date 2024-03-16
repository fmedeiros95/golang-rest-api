package core

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func RespondWithError(ctx *fiber.Ctx, statusCode int, message string) error {
	response := Response{Message: message, StatusCode: statusCode}
	return ctx.Status(statusCode).JSON(response)
}

func RespondWithJSON(ctx *fiber.Ctx, statusCode int, data interface{}, message string) error {
	response := Response{Message: message, StatusCode: statusCode, Data: &data}
	return ctx.Status(statusCode).JSON(response)
}
