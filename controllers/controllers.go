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
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUser(c fiber.Ctx) error {
	db := databases.SharedConnection()

	if db == nil {
		log.Println("Database connection is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	var users []User

	result := db.Find(&users)

	if result.Error != nil {
		log.Printf("Failed to fetch User: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch user",
		})
	}

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No user found",
		})
	}

	return c.JSON(users)
}
