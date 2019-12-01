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

	sum_part_two := solve(AdditionalFuelCounterUpperRecursive)
	fmt.Printf("Answer for Part 2: %d\n", sum_part_two)

	sum_part_two_b := solve_via_channel(AdditionalFuelCounterUpperChannel)
	fmt.Printf("Answer for Part 2 (via channel): %d\n", sum_part_two_b)
}

func solve_via_channel(fun func(int, int, chan int)) int {
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

func AdditionalFuelCounterUpperChannel(mass int, total int, result chan int) {
	fuel := FuelCounterUpper(mass)

	if fuel <= 0 {
		result <- total
		return
	}

	go AdditionalFuelCounterUpperChannel(fuel, total+fuel, result)
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
