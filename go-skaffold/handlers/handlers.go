package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HandlerFunc = func(c *fiber.Ctx) error

func ToFiberError(err error) *fiber.Error {
	code := fiber.StatusInternalServerError
	message := "server/error"

	return &fiber.Error{
		Code:    code,
		Message: message,
	}
}
