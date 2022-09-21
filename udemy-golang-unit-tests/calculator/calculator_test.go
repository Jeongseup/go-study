package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	args := drm.Called()

	return args.Int(0)
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
		{
			name:                  "should apply 20",
			minimumPurchaseAmount: 100,
			discout:               20,
			purchaseAmount:        150,
			expectedAmount:        130,
		},
		{
			name:                  "should apply 40",
			minimumPurchaseAmount: 100,
			discout:               20,
			purchaseAmount:        200,
			expectedAmount:        160},
		{
			name:                  "should apply 60",
			minimumPurchaseAmount: 100,
			discout:               20,
			purchaseAmount:        350,
			expectedAmount:        290,
		},
		{
			name:                  "should not apply",
			minimumPurchaseAmount: 100,
			discout:               20,
			purchaseAmount:        50,
			expectedAmount:        50,
		},
		{
			name:                  "should not apply when discount is zero",
			minimumPurchaseAmount: 100,
			discout:               20,
			purchaseAmount:        50,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositryMock := DiscountRepositoryMock{}
			discountRepositryMock.On("FindCurrentDiscount").Return(tc.discout)

			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositryMock)
			if err != nil {
				t.Fatalf("could not instantiate the calculator %s", err.Error())
			}

			amount := calculator.Calculator(tc.purchaseAmount)

			// t, expected, actual
			assert.Equal(t, tc.expectedAmount, amount)
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
