package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (p *PingHandler) Get() HandlerFunc {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"result": "ok",
		})
	}
}
