package routes

import (
	"github.com/akumakumu/suiren/controllers"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	app.Get("/user", controllers.GetUser)
	app.Get("/user/:id", controllers.GetUserById)
	app.Post("/user", controllers.CreateUser)
}
