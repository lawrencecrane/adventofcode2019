package main

import (
	"testing"
)

func Helper(t *testing.T, f func(int) int, x, expected int) {
	ans := f(x)

	if ans != expected {
		t.Errorf("Expected %d, Got %d", expected, ans)
	}
}

func FuelCounterUpperTestHelper(t *testing.T, x, expected int) {
	Helper(t, FuelCounterUpper, x, expected)
}

func AdditionalFuelCounterUpperTestHelper(t *testing.T, x, expected int) {
	Helper(t, AdditionalFuelCounterUpperRecursive, x, expected)
}

func TestFuelCounterUpper(t *testing.T) {
	FuelCounterUpperTestHelper(t, 12, 2)
	FuelCounterUpperTestHelper(t, 14, 2)
	FuelCounterUpperTestHelper(t, 1969, 654)
	FuelCounterUpperTestHelper(t, 100756, 33583)
}

func TestAdditionalFuelCounterUpper(t *testing.T) {
	AdditionalFuelCounterUpperTestHelper(t, 12, 2)
	AdditionalFuelCounterUpperTestHelper(t, 1969, 966)
	AdditionalFuelCounterUpperTestHelper(t, 100756, 50346)
}

func BenchmarkSolveIterative(b *testing.B) {
	solve(AdditionalFuelCounterUpperIterative)
}

func BenchmarkSolveRecursive(b *testing.B) {
	solve(AdditionalFuelCounterUpperRecursive)
}

func BenchmarkSolveTailRecursive(b *testing.B) {
	solve(AdditionalFuelCounterUpperTailRecursive)
}

func BenchmarkSolveChannel(b *testing.B) {
	solveViaChannel(AdditionalFuelCounterUpperTailRecursiveChannel)
}
