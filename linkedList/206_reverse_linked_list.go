package linkedlist

// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var revHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
	}

	return revHead
}
