package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "example.com/docs/fibersimple"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server
// @termsOfService http://test.com

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Fiber
	app := fiber.New()
	// Middleware
	app.Use(recover.New())
	app.Use(cors.New())
	// Routes
	app.Get("/", HealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start Server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server
// @Description get the status of server
// @Tags root
// @Accpet */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "server is up and running",
	}

	if err := c.Status(200).JSON(res); err != nil {
		return err
	}

	return nil
}
