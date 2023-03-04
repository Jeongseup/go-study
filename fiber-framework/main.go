package main

import (
	"fmt"
	"log"

	// internal
	"example.com/config"
	"example.com/database"
	"example.com/routes"

	// swagger
	_ "example.com/docs"

	// fiber
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func setupRoutes(app *fiber.App) {
	apiVersion := "/v1"
	apiURL := "/api"

	// Test endpoint
	app.Get(apiVersion+apiURL, routes.Welcome)

	// TODO : public static

	// Swagger endpoint making for API document
	app.Get("/v1/api/swagger/*", swagger.HandlerDefault)

	//==============================================================================
	//Endpoint Group Summary
	//==============================================================================
	account := app.Group(apiVersion + apiURL + "/accounts")
	device := app.Group(apiVersion + apiURL + "/devices")

	//==============================================================================
	//Device endpoints
	//==============================================================================
	device.Get("/", routes.GetDevices)
	device.Get("/:id", routes.GetDevice)
	device.Post("/", routes.CreateDevice)
	device.Post("/:id", routes.UpdateDevice)

	//==============================================================================
	//Account endpoints
	//==============================================================================
	account.Post("/", routes.CreateAccount)
	account.Get("/", routes.GetAccounts)
	account.Get("/:address", routes.GetAccountByAddress)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /v1/api
func main() {
	// Connect to the database
	database.ConnectDatabase()
	// skip auth

	appPort := config.Config("APP_PORT")
	// Create Fiber new instance
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	setupRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", appPort)))
}
