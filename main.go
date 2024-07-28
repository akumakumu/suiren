package main

import (
	"github.com/akumakumu/suiren/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Routes
	routes.Router(app)

	app.Listen(":3000")
}
