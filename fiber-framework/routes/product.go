package routes

import (
	"example.com/database"
	"example.com/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProdcut(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}
