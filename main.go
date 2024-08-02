package main

import (
	"log"

	"github.com/akumakumu/suiren/databases"
	"github.com/akumakumu/suiren/models"
	"github.com/akumakumu/suiren/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Database Connection
	err := databases.InitDatabase()

	if err != nil {
		log.Fatalf("Failed Connect Database: %v", err)
	}

	databases.DB.AutoMigrate(&models.User{})
	log.Printf("Database Migrated")

	pgDB, err := databases.DB.DB()

	if err != nil {
		log.Fatalf("Failed Getting Database Connection: %v", err)
	}

	defer func() {
		if err := pgDB.Close(); err != nil {
			log.Printf("Closing database Connection: %v", err)
		}
	}()

	app := fiber.New()

	routes.Router(app)

	app.Listen(":3000")
}
