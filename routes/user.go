package routes

import (
	"fiber-postgre/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	api := app.Group("/users")
	api.Get("/", handlers.GetUsers)
	api.Get("/:id", handlers.GetUser)
	api.Post("/", handlers.CreateUser)
	api.Put("/:id", handlers.UpdateUser)
	api.Delete("/:id", handlers.DeleteUser)
}
