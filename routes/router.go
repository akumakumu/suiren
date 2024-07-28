package routes

import "github.com/gofiber/fiber/v3"

func Router(r *fiber.App) {
	r.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello Mom!")
	})
}
