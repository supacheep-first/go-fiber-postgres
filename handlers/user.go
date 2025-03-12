package handlers

import (
	"fiber-postgre/database"
	"fiber-postgre/models"
	"fiber-postgre/validators"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := validators.Validate.Struct(user); err != nil {
		return c.Status(422).JSON(fiber.Map{"validation_error": err.Error()})
	}
	database.DB.Create(user)
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	updateData := new(models.User)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := validators.Validate.Struct(updateData); err != nil {
		return c.Status(422).JSON(fiber.Map{"validation_error": err.Error()})
	}

	user.Name = updateData.Name
	user.Email = updateData.Email
	database.DB.Save(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	database.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted"})
}
