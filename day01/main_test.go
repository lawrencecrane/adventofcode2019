package main

import (
	"math/rand"
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

var DATA = get_random_dataset()

func benchmark(fun func(int) int) {
	for _, value := range DATA {
		fun(value)
	}
}

func BenchmarkSolveIterative(b *testing.B) {
	benchmark(AdditionalFuelCounterUpperIterative)
}

func BenchmarkSolveRecursive(b *testing.B) {
	benchmark(AdditionalFuelCounterUpperRecursive)
}

func BenchmarkSolveTailRecursive(b *testing.B) {
	benchmark(AdditionalFuelCounterUpperTailRecursive)
}

func BenchmarkSolveChannel(b *testing.B) {
	for _, value := range DATA {
		res := make(chan int)
		AdditionalFuelCounterUpperTailRecursiveChannel(value, 0, res)
	}
}

func get_random_dataset() [10000]int {
	var data [10000]int

	for key, _ := range data {
		data[key] = random(1000000, 10000000)
	}

	return data
}

func random(min, max int) int {
	return min + rand.Intn(max-min)
}
