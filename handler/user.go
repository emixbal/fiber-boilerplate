package handler

import (
	"fiber-boilerplate/database"
	"fiber-boilerplate/model"

	"github.com/gofiber/fiber"
)

// ListUser query all users
func ListUser(c *fiber.Ctx) {
	db := database.DB
	var users []model.User
	db.Find(&users)
	c.JSON(fiber.Map{"status": "success", "message": "All products", "data": users})
}
