package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")

	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		sum += FuelCounterUpper(int(x))
	}

	fmt.Printf("Answer for Part 1: %d\n", sum)
}

func FuelCounterUpper(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
