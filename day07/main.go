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

	signal, phase := findMostAmplified(stack)

	fmt.Printf("Most Amplifed signal is %v\n", signal)
	fmt.Printf("Most Amplifed phase is %v\n", phase)
}

func findMostAmplified(stack []int) (int, []int) {
	max := 0
	imax := 0

	stacks := nCopyStack(5, stack)
	permutations := permutationsHeap([]int{0, 1, 2, 3, 4})

	for i, phases := range permutations {
		output, _, _ := execCircuit(stacks, phases, 0)

		if output > max {
			max = output
			imax = i
		}
	}

	return max, permutations[imax]
}

// func execFeedbackLoop(stack []int, phases []int) (int, error) {
// 	input := 0
// 	halted := false
// 	stacks := nCopyStack(len(phases), stack)

// 	for !halted {
// 		execCircuit(stacks, phases, input)
// 	}

// 	return 1, nil
// }

func nCopyStack(n int, stack []int) [][]int {
	stacks := make([][]int, n)

	for i := range stacks {
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		stacks[i] = tmp
	}

	return stacks
}

func execCircuit(stacks [][]int, phases []int, input int) (int, bool, error) {
	for i, phase := range phases {
		output, halted, err := exec(stacks[i], []int{phase, input})

		if err != nil {
			return 1, false, err
		}

		if halted {
			return output, halted, nil
		}

		input = output
	}

	return input, false, nil
}

func exec(stack []int, input []int) (int, bool, error) {
	_, _, output, halted, err := execHelper(stack, 0, input)

	if err != nil {
		return 1, false, errors.New("Execution failed")
	}

	return output, halted, nil
}

func execHelper(stack []int, pos int, input []int) ([]int, int, int, bool, error) {
	opcode, modes := parseInstruction(stack[pos])
	modes = addPaddingToModes(opcode, modes)

	switch opcode {
	case ADD:
		stack, pos := calc(stack, modes, pos, add)
		return execHelper(stack, pos, input)
	case MULTIPLY:
		stack, pos := calc(stack, modes, pos, mult)
		return execHelper(stack, pos, input)
	case INPUT:
		positionModeWrite(stack, pos+1, input[0])
		return execHelper(stack, pos+2, input[1:])
	case OUTPUT:
		output, pos := out(stack, modes, pos)
		return stack, pos, output, false, nil
	case JUMPT:
		pos := jump(stack, modes, pos, true)
		return execHelper(stack, pos, input)
	case JUMPF:
		pos := jump(stack, modes, pos, false)
		return execHelper(stack, pos, input)
	case LESS:
		stack, pos := comparison(stack, modes, pos, isLess)
		return execHelper(stack, pos, input)
	case EQUALS:
		stack, pos := comparison(stack, modes, pos, isEqual)
		return execHelper(stack, pos, input)
	case HALT:
		return stack, pos, 0, true, nil
	default:
		return nil, 0, 1, false, errors.New("Cannot execute given stack")
	}
}

func addPaddingToModes(opcode int, modes []int) []int {
	switch opcode {
	case
		ADD,
		MULTIPLY,
		LESS,
		EQUALS:
		return addTrailingZeros(modes, 3-len(modes))
	case
		JUMPT,
		JUMPF:
		return addTrailingZeros(modes, 2-len(modes))
	case OUTPUT:
		return addTrailingZeros(modes, 1-len(modes))
	default:
		return modes
	}
}

func comparison(stack []int, modes []int, pos int, pred func(int, int) bool) ([]int, int) {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	if pred(fFst(stack, pos+1), fSnd(stack, pos+2)) {
		positionModeWrite(stack, pos+3, 1)
	} else {
		positionModeWrite(stack, pos+3, 0)
	}

	return stack, pos + 4
}

func jump(stack []int, modes []int, pos int, cmp bool) int {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	if (fFst(stack, pos+1) != 0) == cmp {
		return fSnd(stack, pos+2)
	}

	return pos + 3
}

func calc(stack []int, modes []int, pos int, f func(int, int) int) ([]int, int) {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	res := f(fFst(stack, pos+1), fSnd(stack, pos+2))

	positionModeWrite(stack, pos+3, res)
	return stack, pos + 4
}

func out(stack []int, modes []int, pos int) (int, int) {
	fun, _ := modeToFunc(modes[0])
	return fun(stack, pos+1), pos + 2
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

func permutationsHeap(xs []int) [][]int {
	return permutationsHeapHelper(len(xs), xs, [][]int{})
}

func permutationsHeapHelper(n int, xs []int, permutations [][]int) [][]int {
	if n == 1 {
		tmp := make([]int, len(xs))
		copy(tmp, xs)
		return append(permutations, tmp)
	}

	var res [][]int

	for i := 0; i < n; i++ {
		res = append(res, permutationsHeapHelper(n-1, xs, permutations)...)

		if n%2 == 0 {
			swap(xs, i, n-1)
		} else {
			swap(xs, 0, n-1)
		}
	}

	return res
}

func swap(xs []int, i, j int) {
	tmp := xs[i]
	xs[i] = xs[j]
	xs[j] = tmp
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
