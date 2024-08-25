package linkedlist

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			[][]int{{1, 2, 4}, {1, 3, 4}},
			[]int{1, 1, 2, 3, 4, 4},
		},
		{
			[][]int{{}, {}},
			[]int{},
		},
		{
			[][]int{{}, {0}},
			[]int{0},
		},
	}

	for _, test := range tests {
		t.Run("MergeTwoLists", func(t *testing.T) {
			list1 := convertSliceToList(test.input[0])
			list2 := convertSliceToList(test.input[1])
			expectedList := convertSliceToList(test.expected)
			mergedList := mergeTwoLists(list1, list2)

			if !reflect.DeepEqual(expectedList, mergedList) {
				t.Errorf("mergeTwoLists(%v, %v) = %v; expected %v", test.input[0], test.input[1], convertListToSlice(mergedList), test.expected)
			}
		})
	}
}
