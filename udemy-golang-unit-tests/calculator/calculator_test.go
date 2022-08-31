package calculator

import (
	"testing"
)

func TestDiscoountCalulator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		discout               int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 150, expectedAmount: 130},
		{minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 200, expectedAmount: 160},
		{minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 350, expectedAmount: 290},
		{minimumPurchaseAmount: 100, discout: 20, purchaseAmount: 50, expectedAmount: 50},
	}

	for _, tc := range testCases {
		// t.Log(tc)
		tempCalculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discout)
		tempAmount := tempCalculator.Calculator(tc.purchaseAmount)

		if tempAmount != tc.expectedAmount {
			t.Errorf("expected %v, got %v", tc.expectedAmount, tempAmount)
		}
	}
}

// func Test_DiscountApplied(t *testing.T) {
// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculator(150)

// 	if amount != 130 {
// 		t.FailNow()
// 	}
// }

// func Test_DiscountMultipliedByTwoApplied(t *testing.T) {
// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculator(200)

// 	if amount != 160 {
// 		t.Errorf("expected 160, got %v", amount) // Error = Log + Fail
// 	}
// }

// func Test_DiscountNotApplied(t *testing.T) {
// 	calculator := NewDiscountCalculator(100, 20)
// 	amount := calculator.Calculator(50)

// 	if amount != 50 {
// 		// t.Logf("expected 50, got %v", amount)
// 		// t.Fail()

// 		// same as above
// 		t.Errorf("expected 50, got %v. failed because the discount was not expected", amount) // Error = Log + Fail
// 	}
// }
