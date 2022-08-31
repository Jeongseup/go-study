package calculator

import (
	"testing"
)

type DiscountRepositoryMock struct{}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func TestDiscoountCalulator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discout               int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 150, expectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 200, expectedAmount: 160},
		{name: "should apply 60", minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 350, expectedAmount: 290},
		{name: "should not apply", minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositryMock := DiscountRepositoryMock{}
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositryMock)
			if err != nil {
				t.Fatalf("could not instantiate the calculator %s", err.Error())
			}

			amount := calculator.Calculator(tc.purchaseAmount)

			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			}
		})
	}
}

func TestDiscountCalculatorShoudFailWithZeroMinimumAmount(t *testing.T) {
	discountRepositryMock := DiscountRepositoryMock{}
	_, err := NewDiscountCalculator(0, discountRepositryMock)
	if err == nil {
		t.Fatalf("should not create the calculator with zero value")
	}
}
