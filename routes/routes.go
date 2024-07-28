package routes

import (
	"github.com/akumakumu/suiren/controllers"
	"github.com/gofiber/fiber/v3"
)

func Router(app *fiber.App) {
	app.Get("/", controllers.GetUser)
}
