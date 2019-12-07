package main

import (
	"testing"
)

func TestFindMostAmplfied(t *testing.T) {
	FindMostAmplifiedTestHelper(t,
		[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		[]int{4, 3, 2, 1, 0},
		43210)

	FindMostAmplifiedTestHelper(t,
		[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		[]int{0, 1, 2, 3, 4},
		54321)

	FindMostAmplifiedTestHelper(t,
		[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		[]int{1, 0, 4, 3, 2},
		65210)
}

func TestExecCircuit(t *testing.T) {
	ExecCircuitTestHelper(t,
		[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		[]int{4, 3, 2, 1, 0},
		43210)

	ExecCircuitTestHelper(t,
		[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		[]int{0, 1, 2, 3, 4},
		54321)

	ExecCircuitTestHelper(t,
		[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		[]int{1, 0, 4, 3, 2},
		65210)
}

func TestParseInstruction(t *testing.T) {
	ParseInstructionTestHelper(t, 1234, 34, []int{2, 1})
	ParseInstructionTestHelper(t, 1203004, 4, []int{0, 3, 0, 2, 1})

}

func FindMostAmplifiedTestHelper(t *testing.T, stack, expectedPhase []int, expectedSignal int) {
	signal, phase := findMostAmplified(stack)

	if signal != expectedSignal || !Equal(phase, expectedPhase) {
		t.Errorf("Expected: %v, Got %v -- Expected: %v, Got: %v",
			expectedSignal,
			signal,
			expectedPhase,
			phase)
	}
}

func ExecCircuitTestHelper(t *testing.T, stack, phases []int, expected int) {
	stacks := nCopyStack(len(phases), stack)
	signal, _, _ := execCircuit(stacks, phases, 0)

	if signal != expected {
		t.Errorf("Expected: %v, Got %v", expected, signal)
	}
}

func ParseInstructionTestHelper(t *testing.T, x, expectedOpcode int, expectedModes []int) {
	opcode, modes := parseInstruction(x)

	if !Equal(modes, expectedModes) {
		t.Errorf("Expected %v, Got %v", expectedModes, modes)
	}

	if opcode != expectedOpcode {
		t.Errorf("Expected %v, Got %v", expectedOpcode, opcode)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
