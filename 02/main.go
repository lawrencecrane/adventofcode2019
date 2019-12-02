package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, value := range input(solve) {
		fmt.Println(value)
	}
}

func solve(in string) []int {
	return exec(parse(split(in)))
}

func exec(in []int) []int {
	return in
}

func parse(in []string) []int {
	out := make([]int, len(in))

	for key, value := range in {
		parsed, _ := strconv.ParseInt(value, 10, 32)
		out[key] = int(parsed)
	}

	return out
}

func split(in string) []string {
	return strings.Split(in, ",")
}

func input(solver func(string) []int) []int {
	f, err := os.Open("input")

	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return solver(scanner.Text())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
