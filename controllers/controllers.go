package controllers

import (
	"errors"
	"log"

	"github.com/akumakumu/suiren/databases"
	"github.com/akumakumu/suiren/models"
	"github.com/akumakumu/suiren/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
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

func GetUserById(c *fiber.Ctx) error {
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

func CreateUser(c *fiber.Ctx) error {
	db := databases.SharedConnection()

	var request models.CreateUserRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	hashedPassword, _ := utils.HashPassword(request.Password)

	user := models.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Password: hashedPassword,
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
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

func Login(c *fiber.Ctx) error {
	db := databases.SharedConnection()

	if db == nil {
		log.Println("Database connection is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database connection not initialized",
		})
	}

	var request models.LoginRequest
	var user models.User

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// For Debugging
	log.Printf("Parsed request: %+v", request)

	result := db.Where("username = ?", request.Username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"token":   token,
	})
}

func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {
	// Retrieve user from context
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid JWT token",
		})
	}

	// Parse claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse JWT claims",
		})
	}

	// Retrieve username from claims
	username, ok := claims["username"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Username claim not found or invalid",
		})
	}

	// Respond with a welcome message
	return c.SendString("Welcome " + username)
}
