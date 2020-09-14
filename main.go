package main

import (
	"fiber-boilerplate/router"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(3000)
}
