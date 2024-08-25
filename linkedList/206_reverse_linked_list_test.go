package linkedlist

import (
	"reflect"
	"testing"
)

func convertSliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func convertListToSlice(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},   // edge case: empty list
		{[]int{1}, []int{1}}, // edge case: single element
	}

	for _, test := range tests {
		t.Run("ReverseLinkedList", func(t *testing.T) {
			inputList := convertSliceToList(test.input)
			expectedList := convertSliceToList(test.expected)
			reversedList := reverseList(inputList)

			if !reflect.DeepEqual(expectedList, reversedList) {
				t.Errorf("reverseList(%v) = %v; expected %v", test.input, convertListToSlice(reversedList), test.expected)
			}
		})
	}

}
