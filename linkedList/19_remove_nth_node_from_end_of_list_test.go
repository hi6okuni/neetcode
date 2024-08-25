package linkedlist

import (
	"reflect"
	"testing"
)

type RemoveNthFromEndParam struct {
	head []int
	n    int
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input    RemoveNthFromEndParam
		expected []int
	}{
		{
			RemoveNthFromEndParam{[]int{1, 2, 3, 4, 5}, 2},
			[]int{1, 2, 3, 5},
		},
		{
			RemoveNthFromEndParam{[]int{1}, 1},
			[]int{},
		},
		{
			RemoveNthFromEndParam{[]int{1, 2}, 1},
			[]int{1},
		},
		{
			RemoveNthFromEndParam{[]int{1, 2, 3}, 3},
			[]int{2, 3},
		},
	}

	for _, test := range tests {
		t.Run("RemoveNthFromEnd", func(t *testing.T) {
			head := convertSliceToList(test.input.head)
			expectedList := convertSliceToList(test.expected)
			removedList := removeNthFromEnd(head, test.input.n)

			if !reflect.DeepEqual(expectedList, removedList) {
				t.Errorf("removeNthFromEnd(%v, %v) = %v; expected %v", test.input.head, test.input.n, convertListToSlice(removedList), test.expected)
			}
		})
	}
}
