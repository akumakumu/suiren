package routes

import (
	"log"
	"os"

	"github.com/akumakumu/suiren/handlers"
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
	app.Post("/login", handlers.Login)

	// Restricted all routes after this
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	// }))

	app.Get("/user", handlers.GetUser)
	app.Post("/user", handlers.CreateUser)
	app.Get("/user/:id", handlers.GetUserById)
	app.Delete("/user/:id", handlers.DeleteUser)

	// app.Get("/restricted", handlers.Restricted)

	// Example of only using jwt middleware on one route
	app.Get("/restricted", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	}), handlers.Restricted)

}
