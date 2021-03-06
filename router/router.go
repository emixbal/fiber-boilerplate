package router

import (
	"fiber-boilerplate/handler"
	"fiber-boilerplate/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	api.Get("/users", handler.ListUser)

	api.Get("/test_middleware", middleware.Protected(), handler.Hello)
}
