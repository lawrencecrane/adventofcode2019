package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sum_part_one := solve(FuelCounterUpper)
	fmt.Printf("Answer for Part 1: %d\n", sum_part_one)

	sum_part_two := solve(AdditionalFuelCounterUpper)
	fmt.Printf("Answer for Part 2: %d\n", sum_part_two)
}

func solve(fun func(int) int) int {
	f, err := os.Open("input")

	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		sum += fun(int(x))
	}

	return sum
}

func AdditionalFuelCounterUpper(mass int) int {
	fuel := FuelCounterUpper(mass)

	if fuel <= 0 {
		return 0
	}

	return fuel + AdditionalFuelCounterUpper(fuel)
}

func FuelCounterUpper(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
