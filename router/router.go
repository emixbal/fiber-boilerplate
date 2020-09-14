package router

import (
	"ev-fiber/handler"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
}
