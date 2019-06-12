package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	comp := twoSum([]int{3, 2, 4}, 6)
	if comp[0] != 1 || comp[1] != 2 {
		t.Error("Failed to get expected compliments")
	}
}
