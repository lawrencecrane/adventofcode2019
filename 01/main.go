package main

import (
	"math"
)

func main() {
}

func FuelCounterUpper(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}
