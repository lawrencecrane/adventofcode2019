package main

import (
	"testing"
)

func TestExec(t *testing.T) {
	ExecTestHelper(t, []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	ExecTestHelper(t, []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	ExecTestHelper(t, []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	ExecTestHelper(t, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99})
	ExecTestHelper(t, []int{1, 0, 0, 4, 99, 5, 6, 0, 99}, []int{30, 0, 0, 4, 2, 5, 6, 0, 99})
}

func TestExecWithNoMutation(t *testing.T) {
	stack := []int{1, 0, 0, 0, 99}
	execWithNoMutation(stack, 0, 0)

	if !Equal(stack, []int{1, 0, 0, 0, 99}) {
		t.Error("Exec mutates stack")
	}
}

func ExecTestHelper(t *testing.T, x, expected []int) {
	ans, _ := exec(x)

	if !Equal(ans, expected) {
		t.Errorf("Expected %v, Got %v", expected, ans)
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
