package controllers

import (
	"log"

	"github.com/akumakumu/suiren/databases"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"-"`
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

func GetUserById(c fiber.Ctx) error {
	id := c.Params("id")
	db := databases.SharedConnection()

	if db == nil {
		log.Println("Database connection is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	var user User

	result := db.Find(&user, id)

	if result.Error != nil {
		log.Printf("Failed to fetch User: %v", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch user",
		})
	}

	return c.JSON(user)
}

func CreateUser(c fiber.Ctx) error {
	db := databases.SharedConnection()

	var user User

	user.Fullname = "Nightingale Baldea"
	user.Username = "nightingale"
	user.Password = "unencrypted"

	db.Create(&user)

	return c.JSON(user)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	db := databases.SharedConnection()

	var user User

	db.First(&user, id)

	if user.Username == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User w/ that id not found",
		})
	}

	db.Delete(&user)

	return c.JSON(user)
}
