package calculator

import (
	"errors"
	"golang-unit-tests/database"
)

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountRepository    database.Repository
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountRepository database.Repository) (*DiscountCalculator, error) {

	if minimumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not be zero")
	}

	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}

func (c *DiscountCalculator) Calculator(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {

		multiplier := purchaseAmount / c.minimumPurchaseAmount

		discount := c.discountRepository.FindCurrentDiscount()
		return purchaseAmount - (discount * multiplier)
	}

	return purchaseAmount
}
