package handlers

import (
	"fiber-postgre/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{"id": id, "name": "Mock User"})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
