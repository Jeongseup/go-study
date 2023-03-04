package routes

import (
	"context"
	"fmt"

	"example.com/database"
	"example.com/models"
	"example.com/utils"
	"github.com/gofiber/fiber/v2"
)

//@Description Geting for welcome endpoint
//@Tags Welcome
//@Accept json
//@Product json
//@Success 200 {string} string
//@router / [get]
func Welcome(c *fiber.Ctx) error {
	ctx := context.Background()

	var device models.Device
	err := database.Database.DB.QueryRowContext(ctx, "SELECT * FROM push.device").Scan(&device)
	fmt.Println(device)

	if err != nil {
		c.Status(utils.CODE_FAILED).JSON(err)
	}
	return c.Status(utils.CODE_NOT_FOUND).JSON(device)
}
