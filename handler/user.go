package handler

import (
	"github.com/gofiber/fiber"
)

// ListUser hanlde api status
func ListUser(c *fiber.Ctx) {
	c.JSON(fiber.Map{"status": "success", "message": "Im user handler", "data": nil})
}
