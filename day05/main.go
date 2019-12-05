package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	// res := solveInput(input())
}

func solveInput(in string) int {
	stack := parse(split(in))
	res, _ := execWithNoMutation(stack, 12, 2)

	return res[0]
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

func input() string {
	f, err := os.Open("input")

	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return scanner.Text()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
