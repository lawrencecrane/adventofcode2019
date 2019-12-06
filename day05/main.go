package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Opcodes
const (
	ADD           = 1
	MULTIPLY      = 2
	INPUT         = 3
	OUTPUT        = 4
	JUMP_IF_TRUE  = 5
	JUMP_IF_FALSE = 6
	LESS_THAN     = 7
	EQUALS        = 8
	HALT          = 99
)

// Modes
const (
	POSITION  = 0
	IMMEDIATE = 1
)

func main() {
	stack := parse(split(loadInput()))

	code, output, _ := exec(stack, 1)

	fmt.Printf("Output is %v\n", output)
	fmt.Printf("Diagnostic code is %v\n", code)

	code, output, _ = exec(stack, 5)

	fmt.Printf("Output is %v\n", output)
	fmt.Printf("Diagnostic code is %v\n", code)
}

func exec(stack []int, input int) (int, []int, error) {
	copied := make([]int, len(stack))
	copy(copied, stack)

	var output []int
	started, _ := execStart(copied, input)
	_, _, output, err := execHelper(started, 2, output)

	if err != nil {
		return 1, nil, err
	}

	return output[len(output)-1], output[:len(output)-1], nil
}

func execStart(stack []int, input int) ([]int, error) {
	switch stack[0] {
	case INPUT:
		positionModeWrite(stack, 1, input)
		return stack, nil
	default:
		return nil, errors.New("Cannot start execution")
	}
}

func execHelper(stack []int, pos int, output []int) ([]int, int, []int, error) {
	opcode, modes := parseInstruction(stack[pos])

	switch opcode {
	case ADD:
		stack := calc(stack, addTrailingZeros(modes, 3-len(modes)), pos, add)
		return execHelper(stack, pos+4, output)
	case MULTIPLY:
		stack := calc(stack, addTrailingZeros(modes, 3-len(modes)), pos, mult)
		return execHelper(stack, pos+4, output)
	case OUTPUT:
		output := out(stack, addTrailingZeros(modes, 1-len(modes)), output, pos)
		return execHelper(stack, pos+2, output)
	case HALT:
		return stack, pos, output, nil
	default:
		return nil, 0, output, errors.New("Cannot execute given stack")
	}
}

func calc(stack []int, modes []int, pos int, f func(int, int) int) []int {
	f_fst, _ := modeToFunc(modes[0])
	f_snd, _ := modeToFunc(modes[1])
	res := f(f_fst(stack, pos+1), f_snd(stack, pos+2))

	positionModeWrite(stack, pos+3, res)

	return stack
}

func out(stack []int, modes []int, output []int, pos int) []int {
	fun, _ := modeToFunc(modes[0])

	output = append(output, fun(stack, pos+1))
	return output
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func modeToFunc(mode int) (func([]int, int) int, error) {
	switch mode {
	case POSITION:
		return positionModeRead, nil
	case IMMEDIATE:
		return immediateModeRead, nil
	default:
		return nil, errors.New("Unknown mode")
	}
}

func parseInstruction(x int) (int, []int) {
	modes := parseModes(x)
	modes = addTrailingZeros(modes, 3-len(modes))

	opcode := modes[0] + 10*modes[1]

	return opcode, modes[2:]
}

func parseModes(x int) []int {
	var modes []int

	for x >= 1 {
		modes = append(modes, x%10)

		x = x / 10
	}

	return modes
}

func addTrailingZeros(modes []int, n int) []int {
	for i := 0; i < n; i++ {
		modes = append(modes, 0)
	}

	return modes
}

func positionModeWrite(stack []int, pos, value int) {
	stack[immediateModeRead(stack, pos)] = value
}

func positionModeRead(stack []int, pos int) int {
	return stack[immediateModeRead(stack, pos)]
}

func immediateModeRead(stack []int, pos int) int {
	return stack[pos]
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

func loadInput() string {
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
