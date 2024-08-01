package controllers

import (
	"log"

	"github.com/akumakumu/suiren/databases"
	"github.com/akumakumu/suiren/models"
	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	db := databases.SharedConnection()

	if db == nil {
		log.Println("Database connection is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	var users []models.User

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

	var user models.User

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

	var request models.CreateUserRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Password: request.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	db := databases.SharedConnection()

	var user models.User

	db.First(&user, id)

	if user.Username == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User w/ that id not found",
		})
	}

	db.Delete(&user)

	return c.JSON(user)
}
