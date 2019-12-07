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
	ADD      = 1
	MULTIPLY = 2
	INPUT    = 3
	OUTPUT   = 4
	JUMPT    = 5
	JUMPF    = 6
	LESS     = 7
	EQUALS   = 8
	HALT     = 99
)

// Modes
const (
	POSITION  = 0
	IMMEDIATE = 1
)

func main() {
	stack := parse(split(loadInput()))

	code, output, _ := exec(stack, 1, 0)

	fmt.Printf("Output is %v\n", output)
	fmt.Printf("Diagnostic code is %v\n", code)

	code, output, _ = exec(stack, 5, 0)

	fmt.Printf("Output is %v\n", output)
	fmt.Printf("Diagnostic code is %v\n", code)
}

func exec(stack []int, input, phase int) (int, []int, error) {
	copied := make([]int, len(stack))
	copy(copied, stack)

	var output []int
	_, _, output, err := execHelper(copied, 0, []int{input, phase}, output)

	if err != nil || len(output) < 1 {
		return 1, nil, err
	}

	return output[len(output)-1], output[:len(output)-1], nil
}

func execHelper(stack []int, pos int, input, output []int) ([]int, int, []int, error) {
	opcode, modes := parseInstruction(stack[pos])

	switch opcode {
	case ADD:
		stack := calc(stack, addTrailingZeros(modes, 3-len(modes)), pos, add)
		return execHelper(stack, pos+4, input, output)
	case MULTIPLY:
		stack := calc(stack, addTrailingZeros(modes, 3-len(modes)), pos, mult)
		return execHelper(stack, pos+4, input, output)
	case INPUT:
		positionModeWrite(stack, pos+1, input[0])
		return execHelper(stack, pos+2, input[1:], output)
	case OUTPUT:
		output := out(stack, addTrailingZeros(modes, 1-len(modes)), output, pos)
		return execHelper(stack, pos+2, input, output)
	case JUMPT:
		pos := jump(stack, addTrailingZeros(modes, 2-len(modes)), pos, true)
		return execHelper(stack, pos, input, output)
	case JUMPF:
		pos := jump(stack, addTrailingZeros(modes, 2-len(modes)), pos, false)
		return execHelper(stack, pos, input, output)
	case LESS:
		stack := comparison(stack, addTrailingZeros(modes, 3-len(modes)), pos, isLess)
		return execHelper(stack, pos+4, input, output)
	case EQUALS:
		stack := comparison(stack, addTrailingZeros(modes, 3-len(modes)), pos, isEqual)
		return execHelper(stack, pos+4, input, output)
	case HALT:
		return stack, pos, output, nil
	default:
		return nil, 0, output, errors.New("Cannot execute given stack")
	}
}

func comparison(stack []int, modes []int, pos int, pred func(int, int) bool) []int {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	if pred(fFst(stack, pos+1), fSnd(stack, pos+2)) {
		positionModeWrite(stack, pos+3, 1)
	} else {
		positionModeWrite(stack, pos+3, 0)
	}

	return stack
}

func jump(stack []int, modes []int, pos int, cmp bool) int {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	if (fFst(stack, pos+1) != 0) == cmp {
		return fSnd(stack, pos+2)
	}

	return pos + 3
}

func calc(stack []int, modes []int, pos int, f func(int, int) int) []int {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	res := f(fFst(stack, pos+1), fSnd(stack, pos+2))

	positionModeWrite(stack, pos+3, res)
	return stack
}

func out(stack []int, modes []int, output []int, pos int) []int {
	fun, _ := modeToFunc(modes[0])

	output = append(output, fun(stack, pos+1))
	return output
}

func isLess(a, b int) bool {
	return a < b
}

func isEqual(a, b int) bool {
	return a == b
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func pairModeToFuncs(a, b int) (func([]int, int) int, func([]int, int) int) {
	af, _ := modeToFunc(a)
	bf, _ := modeToFunc(b)

	return af, bf
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
