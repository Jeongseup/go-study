package routes

import (
	"context"
	"fmt"

	"example.com/database"
	"example.com/entities"
	"example.com/models"
	"example.com/utils"
	"github.com/gofiber/fiber/v2"
)

// =====

//Function making for create response message for user
func CreateResponseDevice(deviceModel models.Device) entities.Device {
	return entities.Device{ID: deviceModel.ID, Token: deviceModel.Token, SubscribeTx: deviceModel.SubscribeTx, SubscribeAnnouncement: deviceModel.SubscribeAnnouncement}
}

//Function making for build the response message infomation to caller RestAPI
func ReturnResponseMessageDevice(status string, message string, data entities.Device) map[string]interface{} {
	if data.ID == 0 {
		return fiber.Map{"status": status, "message": message, "data": nil}
	} else {
		return fiber.Map{"status": status, "message": message, "data": data}
	}
}

// =====

//@Description Create new infomation of device
//@Tags device
//@Accept application/json
//@Product application/json
//@Success 200 {array} models.Device
//@router /v1/api/device [post]
func CreateDevice(c *fiber.Ctx) error {
	ctx := context.Background()

	var device models.Device
	if err := c.BodyParser(&device); err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))
	}

	res, err := database.Database.DB.NewInsert().
		Model(&device).
		Exec(ctx)

	fmt.Println(res)

	if err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))
	}

	responseDevice := CreateResponseDevice(device)
	return c.Status(utils.CODE_SUCCESS).JSON(ReturnResponseMessageDevice("OK", "Create device successfully", responseDevice))
}

//@Description Get all existing of devices
//@Tags devices
//@Accept application/json
//@Product application/json
//@Success 200 {array} models.Device
//@router /v1/api/devices [get]
func GetDevices(c *fiber.Ctx) error {
	ctx := context.Background()

	var devices []models.Device
	err := database.Database.DB.NewSelect().Model(&devices).Scan(ctx)
	if err != nil {
		// log.Fatalln(err)
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))
	}

	responseDevices := []entities.Device{}
	for _, device := range devices {
		tmpResponseDevice := CreateResponseDevice(device)
		responseDevices = append(responseDevices, tmpResponseDevice)
	}
	return c.Status(utils.CODE_SUCCESS).JSON(fiber.Map{"status": "OK", "message": "OK", "data": responseDevices})
}

//@Description Get for device details
//@Tags device
//@Accept application/json
//@Product application/json
//@Success 200 {object} entities.Device
//@router /v1/api/devices/{id} [get]
func GetDevice(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := c.ParamsInt("id")
	// log.Println(id)
	var device models.Device

	if err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageDevice("FAILED", "wrong para", entities.Device{}))
	}

	err = database.Database.DB.NewSelect().
		Model(&device).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return c.Status(utils.CODE_NOT_FOUND).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))
	}

	responseDevice := CreateResponseDevice(device)
	return c.Status(utils.CODE_SUCCESS).JSON(ReturnResponseMessageDevice("OK", "OK", responseDevice))
}

//Function making for update the device infomation
//@Description Update device information
//@Tags device
//@Accept application/json
//@Product application/json
//@Param subscribe_tx body string true "SubscribeTx"
//@Param subscribe_announcement body string true "SubscribeAnnouncement"
//@Success 200 {object} entities.Device
//@router /v1/api/devices/{id} [put]
func UpdateDevice(c *fiber.Ctx) error {
	ctx := context.Background()
	id, err := c.ParamsInt("id")

	// Checking for id as parameter is valid
	if err != nil {
		return c.Status(utils.CODE_FAILED).
			JSON(ReturnResponseMessageDevice("FAILED", "wrong param", entities.Device{}))
	}

	//Define data type for update user info
	type UpdateDevice struct {
		SubscribeTx           bool `json:"subscribe_tx"`
		SubscribeAnnouncement bool `json:"subscribe_announcement"`
	}

	var updateDeviceData UpdateDevice
	if err := c.BodyParser(&updateDeviceData); err != nil {
		return c.Status(utils.CODE_FAILED).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))
	}

	value := map[string]interface{}{
		"subscribe_tx":           updateDeviceData.SubscribeTx,
		"subscribe_announcement": updateDeviceData.SubscribeAnnouncement,
	}

	res, err := database.Database.DB.NewUpdate().
		Model(&value).
		TableExpr("push.device").
		Where("id = ?", id).
		Returning("*").
		Exec(ctx)

	fmt.Println(res)

	if err != nil {
		return c.Status(utils.CODE_NOT_FOUND).JSON(ReturnResponseMessageDevice("FAILED", err.Error(), entities.Device{}))

	}

	return c.Status(utils.CODE_SUCCESS).JSON(fiber.Map{"status": "OK", "message": "OK", "data": entities.Device{}})
}
