package main

import (
	"log"

	"github.com/akumakumu/suiren/controllers"
	"github.com/akumakumu/suiren/databases"
	"github.com/akumakumu/suiren/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Database Connection
	db, err := databases.InitDatabase()

	if err != nil {
		log.Fatalf("Failed Connect Database: %v", err)
	}

	db.AutoMigrate(&controllers.User{})
	log.Printf("Database Migrated")

	pgDB, err := db.DB()

	if err != nil {
		log.Fatalf("Failed Getting Database Connection: %v", err)
	}

	defer func() {
		if err := pgDB.Close(); err != nil {
			log.Printf("Closing database Connection: %v", err)
		}
	}()

	app := fiber.New()

	// Routes
	routes.Router(app)

	app.Listen(":3000")
}
