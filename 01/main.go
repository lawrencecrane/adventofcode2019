package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Answer for Part 1: %d\n", solve(FuelCounterUpper))

	fmt.Printf("Answer for Part 2 (via iteration): %d\n",
		solve(AdditionalFuelCounterUpperIterative))

	fmt.Printf("Answer for Part 2 (via recursion): %d\n",
		solve(AdditionalFuelCounterUpperRecursive))

	fmt.Printf("Answer for Part 2 (via tail recursion): %d\n",
		solve(AdditionalFuelCounterUpperTailRecursive))

	fmt.Printf("Answer for Part 2 (via tail recursive channel): %d\n",
		solveViaChannel(AdditionalFuelCounterUpperTailRecursiveChannel))
}

func solveViaChannel(fun func(int, int, chan int)) int {
	f, err := os.Open("input")

	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		x, _ := strconv.ParseInt(scanner.Text(), 10, 64)

		res := make(chan int)
		fun(int(x), 0, res)

		sum += <-res
	}

	return sum
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

func AdditionalFuelCounterUpperIterative(mass int) int {
	total := 0

	for fuel := FuelCounterUpper(mass); fuel > 0; {
		total += fuel
		fuel = FuelCounterUpper(fuel)
	}

	return total
}

func AdditionalFuelCounterUpperTailRecursiveChannel(mass int, total int, result chan int) {
	fuel := FuelCounterUpper(mass)

	if fuel <= 0 {
		result <- total
		return
	}

	go AdditionalFuelCounterUpperTailRecursiveChannel(fuel, total+fuel, result)
}

func AdditionalFuelCounterUpperTailRecursive(mass int) int {
	return AdditionalFuelCounterUpperTailRecursiveHelper(mass, 0)
}

func AdditionalFuelCounterUpperTailRecursiveHelper(mass int, total int) int {
	fuel := FuelCounterUpper(mass)

	if fuel <= 0 {
		return total
	}

	return AdditionalFuelCounterUpperTailRecursiveHelper(fuel, total+fuel)
}

func AdditionalFuelCounterUpperRecursive(mass int) int {
	fuel := FuelCounterUpper(mass)

	if fuel <= 0 {
		return 0
	}

	return fuel + AdditionalFuelCounterUpperRecursive(fuel)
}

func FuelCounterUpper(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
