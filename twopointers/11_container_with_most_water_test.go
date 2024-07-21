package twopointers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		// {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		// {[]int{1, 1}, 1},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, tt := range tests {
		t.Run("MaxArea", func(t *testing.T) {
			t.Logf("Running test case with input: %v", tt.input)
			result := MaxArea(tt.input)
			if result != tt.expected {
				t.Errorf("MaxArea(%v) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
