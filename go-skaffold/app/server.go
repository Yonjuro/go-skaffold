package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yonjuro/go-skaffold/handlers"
)

type server struct {
	router *fiber.App

	// Handlers
	pingHandler *handlers.PingHandler
}

func NewServer() *server {
	router := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	return &server{
		router: router,
	}
}

// StartApplication ...
func (s *server) Run() error {
	s.pingHandler = handlers.NewPingHandler()

	s.router.Use(logger.New())

	mapRoutes(s)

	log.Fatal(s.router.Listen(":5000"))

	return nil
}

// Custom error handler
func errorHandler(ctx *fiber.Ctx, err error) error {
	log.Println(ctx.OriginalURL(), err)

	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom statuscode if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error page
	err = ctx.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})

	if err != nil {
		ctx.Status(500).SendString("Internal Server Error")
	}

	return nil
}
