package arraysandhashing

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, tt := range tests {
		t.Run("LongestConsecutive", func(t *testing.T) {
			t.Logf("Running test case with input: %v", tt.input)
			result := LongestConsecutive(tt.input)
			if result != tt.expected {
				t.Errorf("LongestConsecutive(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
