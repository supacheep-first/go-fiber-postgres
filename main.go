package main

import (
	"fiber-postgre/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterUserRoutes(app)

	app.Listen(":3000")
}
