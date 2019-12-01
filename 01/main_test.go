package main

import (
	"testing"
)

func FuelCounterUpperTestHelper(t *testing.T, x, expected int) {
	ans := FuelCounterUpper(x)

	if ans != expected {
		t.Errorf("Expected %d, Got %d", expected, ans)
	}
}

func TestFuelCounterUpper(t *testing.T) {
	FuelCounterUpperTestHelper(t, 12, 2)
	FuelCounterUpperTestHelper(t, 14, 2)
	FuelCounterUpperTestHelper(t, 1969, 654)
	FuelCounterUpperTestHelper(t, 100756, 33583)
}
