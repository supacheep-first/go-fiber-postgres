package main

import (
	"fiber-postgre/database"
	"fiber-postgre/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	routes.RegisterUserRoutes(app)
	app.Listen(":3000")
}
