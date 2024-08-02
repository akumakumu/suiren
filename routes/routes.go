package routes

import (
	"github.com/akumakumu/suiren/controllers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// Login
	app.Post("/login", controllers.Login)

	// Restricted
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Get("/user", controllers.GetUser)
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:id", controllers.GetUserById)
	app.Delete("/user/:id", controllers.DeleteUser)

	app.Get("/restricted", controllers.Restricted)
}
