package arraysandhashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 2, 3, 4, 5, 5}, true},
	}

	for _, test := range tests {
		t.Run("ContainsDuplicate", func(t *testing.T) {
			t.Logf("Running test case with input: %v", test.input)
			result := ContainsDuplicate(test.input)
			if result != test.expected {
				t.Errorf("ContainsDuplicate(%v) = %v; expected %v", test.input, result, test.expected)
			}
		})
	}
}
