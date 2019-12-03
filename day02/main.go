package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Answer to Part 1: %d\n", input(solveInput))
	fmt.Printf("Answer to Part 2: %v\n", input(solveOutput))
}

func solveOutput(in string) int {
	stack := parse(split(in))
	noun, verb, _ := findInputPair(stack, 19690720)

	return 100*noun + verb
}

func solveInput(in string) int {
	stack := parse(split(in))
	res, _ := execWithNoMutation(stack, 12, 2)

	return res[0]
}

func findInputPair(stack []int, output int) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			res, err := execWithNoMutation(stack, noun, verb)

			if err == nil && output == res[0] {
				return noun, verb, nil
			}
		}
	}

	return 0, 0, errors.New("No input pair produces given output")
}

func execWithNoMutation(stack []int, noun, verb int) ([]int, error) {
	copied := make([]int, len(stack))
	copy(copied, stack)

	copied[1] = noun
	copied[2] = verb

	return exec(copied)
}

func exec(stack []int) ([]int, error) {
	stack, _, err := execHelper(stack, 0)
	return stack, err
}

func execHelper(stack []int, pos int) ([]int, int, error) {
	switch stack[pos] {
	case 1:
		return execHelper(execOpcode(pos, add, stack), pos+4)
	case 2:
		return execHelper(execOpcode(pos, multiply, stack), pos+4)
	case 99:
		return stack, pos, nil
	default:
		return nil, 0, errors.New("Cannot execute given stack")
	}
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func execOpcode(pos int, fun func(int, int) int, stack []int) []int {
	stack[stack[pos+3]] = fun(stack[stack[pos+1]], stack[stack[pos+2]])
	return stack
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

func input(solver func(string) int) int {
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
