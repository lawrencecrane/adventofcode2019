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
	RELATIVE  = 2
)

type amplifier struct {
	addr   int
	base   int
	halted bool
	input  []int
	output int
	stack  []int
}

func main() {
	stack := parse(split(loadInput()))

	signal, phase := findMostAmplified(stack,
		execCircuit,
		[]int{0, 1, 2, 3, 4})

	fmt.Printf("Most Amplifed signal is %v\n", signal)
	fmt.Printf("Most Amplifed phase is %v\n", phase)

	signal2, phase2 := findMostAmplified(stack,
		execFeedbackLoop,
		[]int{5, 6, 7, 8, 9})

	fmt.Printf("Most Amplifed feedbacked signal is %v\n", signal2)
	fmt.Printf("Most Amplifed feedbacked phase is %v\n", phase2)
}

func findMostAmplified(stack []int,
	fun func([]amplifier, int) ([]amplifier, bool),
	phaseValues []int) (int, []int) {
	max := 0
	imax := 0

	permutations := permutationsHeap(phaseValues)

	for i, phases := range permutations {
		amps := createAmplifiers(stack, phases)
		executed, _ := fun(amps, 0)

		if executed[len(executed)-1].output > max {
			max = executed[len(executed)-1].output
			imax = i
		}
	}

	return max, permutations[imax]
}

func createAmplifiers(stack []int, phases []int) []amplifier {
	amps := make([]amplifier, len(phases))

	for i, phase := range phases {
		tmp := make([]int, len(stack))
		copy(tmp, stack)

		amps[i] = amplifier{
			addr:   0,
			base:   0,
			halted: false,
			input:  []int{phase},
			output: 0,
			stack:  tmp,
		}
	}

	return amps
}

func execFeedbackLoop(amps []amplifier, input int) ([]amplifier, bool) {
	for {
		executed, halted := execCircuit(amps, input)

		if halted {
			return executed, true
		}

		amps = executed
		input = executed[len(executed)-1].output
	}

	return amps, false
}

func execCircuit(amps []amplifier, input int) ([]amplifier, bool) {
	for i, amp := range amps {
		amp.input = append(amp.input, input)
		executed, err := exec(amp)

		amps[i] = executed

		if err != nil {
			return nil, false
		}

		if executed.halted {
			return amps, true
		}

		input = executed.output
	}

	return amps, false
}

func exec(amp amplifier) (amplifier, error) {
	opcode, modes := parseInstruction(amp.stack[amp.addr])

	switch opcode {
	case ADD:
		amp = calc(amp, modes, add)
	case MULTIPLY:
		amp = calc(amp, modes, mult)
	case LESS:
		amp = calc(amp, modes, ifLess)
	case EQUALS:
		amp = calc(amp, modes, ifEqual)
	case JUMPT:
		amp = jump(amp, modes, true)
	case JUMPF:
		amp = jump(amp, modes, false)
	case INPUT:
		amp = read(amp, modes)
	case OUTPUT:
		amp = out(amp, modes)
		return amp, nil
	case HALT:
		amp.halted = true
		return amp, nil
	default:
		return amp, errors.New("Unknown opcode")
	}

	return exec(amp)
}

func jump(amp amplifier, modes []int, cmp bool) amplifier {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])

	if (fFst(amp, 1) != 0) == cmp {
		amp.addr = fSnd(amp, 2)
	} else {
		amp.addr += 3
	}

	return amp
}

func calc(amp amplifier, modes []int, f func(int, int) int) amplifier {
	fFst, fSnd := pairModeToFuncs(modes[0], modes[1])
	fWrite, _ := modeToWriteFunc(modes[2])

	amp = fWrite(amp, 3, f(fFst(amp, 1), fSnd(amp, 2)))
	amp.addr += 4

	return amp
}

func out(amp amplifier, modes []int) amplifier {
	fun, _ := modeToReadFunc(modes[0])

	amp.output = fun(amp, 1)
	amp.addr += 2

	return amp
}

func read(amp amplifier, modes []int) amplifier {
	fun, _ := modeToWriteFunc(modes[0])

	amp = fun(amp, 1, amp.input[0])
	amp.input = amp.input[1:]
	amp.addr += 2

	return amp
}

func ifLess(a, b int) int {
	if a < b {
		return 1
	}

	return 0
}

func ifEqual(a, b int) int {
	if a == b {
		return 1
	}

	return 0
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func pairModeToFuncs(a, b int) (func(amplifier, int) int, func(amplifier, int) int) {
	af, _ := modeToReadFunc(a)
	bf, _ := modeToReadFunc(b)

	return af, bf
}

func modeToWriteFunc(mode int) (func(amplifier, int, int) amplifier, error) {
	switch mode {
	case POSITION:
		return positionModeWrite, nil
	case RELATIVE:
		return relativeModeWrite, nil
	default:
		return nil, errors.New("Unknown mode")
	}
}

func modeToReadFunc(mode int) (func(amplifier, int) int, error) {
	switch mode {
	case POSITION:
		return positionModeRead, nil
	case IMMEDIATE:
		return immediateModeRead, nil
	case RELATIVE:
		return relativeModeRead, nil
	default:
		return nil, errors.New("Unknown mode")
	}
}

func parseInstruction(x int) (int, []int) {
	modes := parseModes(x)
	modes = addTrailingZeros(modes, 3-len(modes))

	opcode := modes[0] + 10*modes[1]

	return opcode, addPaddingToModes(opcode, modes[2:])
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
	case
		INPUT,
		OUTPUT:
		return addTrailingZeros(modes, 1-len(modes))
	default:
		return modes
	}
}

func addTrailingZeros(modes []int, n int) []int {
	for i := 0; i < n; i++ {
		modes = append(modes, 0)
	}

	return modes
}

func parseModes(x int) []int {
	var modes []int

	for x >= 1 {
		modes = append(modes, x%10)
		x = x / 10
	}

	return modes
}

func relativeModeWrite(amp amplifier, offset, value int) amplifier {
	amp.stack[amp.base+amp.addr] = value
	return amp
}

func positionModeWrite(amp amplifier, offset, value int) amplifier {
	amp.stack[immediateModeRead(amp, offset)] = value
	return amp
}

func relativeModeRead(amp amplifier, offset int) int {
	return amp.stack[amp.base+amp.addr+offset]
}

func positionModeRead(amp amplifier, offset int) int {
	return amp.stack[immediateModeRead(amp, offset)]
}

func immediateModeRead(amp amplifier, offset int) int {
	return amp.stack[amp.addr+offset]
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
