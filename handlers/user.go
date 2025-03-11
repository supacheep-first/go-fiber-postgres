package handlers

import (
	"fiber-postgre/models"

	"github.com/gofiber/fiber/v2"
)

var users = []models.User{
	{ID: "1", Name: "Alice", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", Email: "bob@example.com"},
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, u := range users {
		if u.ID == id {
			return c.JSON(u)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "User not found"})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	users = append(users, *user)
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	for i, u := range users {
		if u.ID == id {
			users[i].Name = user.Name
			users[i].Email = user.Email
			return c.JSON(users[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "User not found"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(fiber.Map{"message": "User deleted"})
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "User not found"})
}
