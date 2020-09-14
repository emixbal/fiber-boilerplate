package main

import (
	"ev-fiber/router"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(3000)
}
