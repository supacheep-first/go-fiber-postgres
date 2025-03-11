package routes

import (
	"fiber-postgre/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	route := app.Group("/users")

	route.Get("/", handlers.GetUsers)
	route.Get("/:id", handlers.GetUserByID)
	route.Post("/", handlers.CreateUser)
}
