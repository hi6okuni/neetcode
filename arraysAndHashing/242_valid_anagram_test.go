package arraysandhashing

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"hello", "billion", false},
		{"a", "a", true},
		{"", "", true},
	}

	for _, test := range tests {
		t.Run("IsAnagram", func(t *testing.T) {
			t.Logf("Running test case with s: %v, t: %v", test.s, test.t)
			result := IsAnagram(test.s, test.t)
			if result != test.expected {
				t.Errorf("IsAnagram(%v, %v) = %v; expected %v", test.s, test.t, result, test.expected)
			}
		})
	}
}
