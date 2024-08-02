package routes

import (
	"log"
	"os"

	"github.com/akumakumu/suiren/controllers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Router(app *fiber.App) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env file")
	}

	// Login
	app.Post("/login", controllers.Login)

	// Restricted
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	}))

	app.Get("/user", controllers.GetUser)
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:id", controllers.GetUserById)
	app.Delete("/user/:id", controllers.DeleteUser)

	app.Get("/restricted", controllers.Restricted)
}
