package main

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	ParseInstructionTestHelper(t, 1234, 34, []int{2, 1})
	ParseInstructionTestHelper(t, 1203004, 4, []int{0, 3, 0, 2, 1})

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
