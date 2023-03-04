package routes

import (
	"context"
	"fmt"
	"log"

	"example.com/database"
	"example.com/entities"
	"example.com/models"
	"example.com/utils"
	"github.com/gofiber/fiber/v2"
)

//Function making for create response message for user
func CreateResponseAccount(accountModel models.Account) entities.Account {
	return entities.Account{
		ID:       accountModel.ID,
		Address:  accountModel.Address,
		DeviceID: accountModel.DeviceID,
	}
}

//Function making for build the response message infomation to caller RestAPI
func ReturnResponseMessageAccount(status string, message string, data entities.Account) map[string]interface{} {
	if data.ID == 0 {
		return fiber.Map{"status": status, "message": message, "data": nil}
	} else {
		return fiber.Map{"status": status, "message": message, "data": data}
	}
}

// =====

func CreateAccount(c *fiber.Ctx) error {
	ctx := context.Background()

	var account models.Account
	if err := c.BodyParser(&account); err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageAccount("FAILED", err.Error(), entities.Account{}))
	}

	res, err := database.Database.DB.NewInsert().
		Model(&account).
		Exec(ctx)

	fmt.Println(res)

	if err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageAccount("FAILED", err.Error(), entities.Account{}))
	}

	responseAccount := CreateResponseAccount(account)
	return c.Status(utils.CODE_SUCCESS).JSON(ReturnResponseMessageAccount("OK", "Create device successfully", responseAccount))
}

func GetAccounts(c *fiber.Ctx) error {
	ctx := context.Background()

	var accounts []models.Account
	err := database.Database.DB.NewSelect().
		Model(&accounts).
		Scan(ctx)

	if err != nil {
		// log.Fatalln(err)
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageAccount("FAILED", err.Error(), entities.Account{}))
	}

	responseAccounts := []entities.Account{}
	for _, account := range accounts {
		tempResponseAccount := CreateResponseAccount(account)
		responseAccounts = append(responseAccounts, tempResponseAccount)
	}
	return c.Status(utils.CODE_SUCCESS).JSON(fiber.Map{"status": "OK", "message": "OK", "data": responseAccounts})
}

func GetAccountById(c *fiber.Ctx) error {
	// ctx := context.Background()
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageAccount("FAILED", "wrong para", entities.Account{}))
	}
	log.Println(id)

	// var device models.Device

	// err = database.Database.DB.NewSelect().
	// 	Model(&device).
	// 	Where("id = ?", id).
	// 	Scan(ctx)

	// if err != nil {
	// 	return c.Status(utils.CODE_NOT_FOUND).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))
	// }

	// responseDevice := CreateResponseDevice(device)
	return c.Status(utils.CODE_SUCCESS).JSON(ReturnResponseMessageAccount("OK", "OK", entities.Account{}))
}

func GetAccountByUserID(c *fiber.Ctx) error {
	return nil
}

func GetAccountByAddress(c *fiber.Ctx) error {
	ctx := context.Background()
	_ = ctx
	address := c.Params("address")
	log.Printf("input params : %s", address)

	var account models.Account

	err := database.Database.DB.NewSelect().
		Model(&account).
		Where("address = ?", address).
		Scan(ctx)

	if err != nil {
		return c.Status(utils.CODE_NOT_FOUND).JSON(ReturnResponseMessageAccount("NOT FOUND", err.Error(), entities.Account{}))
	}

	responseAccount := CreateResponseAccount(account)
	return c.Status(utils.CODE_SUCCESS).JSON(ReturnResponseMessageAccount("OK", "OK", responseAccount))
}
