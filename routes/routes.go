package routes

import (
	"github.com/akumakumu/suiren/controllers"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	// Login
	app.Post("/login", controllers.Login)

	// Public Routes
	app.Get("/user", controllers.GetUser)
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:id", controllers.GetUserById)
	app.Delete("/user/:id", controllers.DeleteUser)
}
