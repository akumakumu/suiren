package controllers

import (
	"log"
	"time"

	"github.com/akumakumu/suiren/databases"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUser(c fiber.Ctx) error {
	db, err := databases.InitDatabase()

	if err != nil {
		log.Fatalf("Failed Connect Database: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to Connect Database")
	}

	var users []User

	db.Find(&users)

	return c.JSON(users)
}
