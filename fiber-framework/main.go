package main

import (
	"log"

	"example.com/database"
	"example.com/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API!!!")
}

func setupRoutes(app *fiber.App) {
	// welcome endpoint
	app.Get("/api", welcome)
	// User endpoint
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Product endpoints
	app.Post("/api/products", routes.CreateProdcut)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
