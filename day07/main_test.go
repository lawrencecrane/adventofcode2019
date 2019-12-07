package main

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	ParseInstructionTestHelper(t, 1234, 34, []int{2, 1})
	ParseInstructionTestHelper(t, 1203004, 4, []int{0, 3, 0, 2, 1})

}

func ParseInstructionTestHelper(t *testing.T, x, expected_opcode int, expected_modes []int) {
	opcode, modes := parseInstruction(x)

	if !Equal(modes, expected_modes) {
		t.Errorf("Expected %v, Got %v", expected_modes, modes)
	}

	if opcode != expected_opcode {
		t.Errorf("Expected %v, Got %v", expected_opcode, opcode)
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
