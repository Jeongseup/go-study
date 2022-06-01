package main

import "testing"

func TestCalculate(t *testing.T) {
	if Caculate(2) != 4 {
		t.Error("Expected 2 + 2 = 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 3},
	}

	for _, test := range tests {
		if output := Caculate(test.input); output != test.expected {
			t.Error("test failed")
		}
	}
}
