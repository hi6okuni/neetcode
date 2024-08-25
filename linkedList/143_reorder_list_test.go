package linkedlist

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
	}

	for _, test := range tests {
		t.Run("ReorderList", func(t *testing.T) {
			inputList := convertSliceToList(test.input)
			expectedList := convertSliceToList(test.expected)
			reorderList(inputList)

			if !reflect.DeepEqual(expectedList, inputList) {
				t.Errorf("reorderList(%v) = %v; expected %v", test.input, convertListToSlice(inputList), test.expected)
			}
		})
	}
}
