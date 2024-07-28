package controllers

import "github.com/gofiber/fiber/v3"

func GetUser(c fiber.Ctx) error {
	return c.SendString("Get A User")
}
