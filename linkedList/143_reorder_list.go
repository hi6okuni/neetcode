package linkedlist

// You are given the head of a singly linked-list. The list can be represented as:

// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.

// Example 1:

// Input: head = [1,2,3,4]
// Output: [1,4,2,3]
// Example 2:

// Input: head = [1,2,3,4,5]
// Output: [1,5,2,4,3]

// Constraints:

// The number of nodes in the list is in the range [1, 5 * 104].
// 1 <= Node.val <= 1000

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find the middle of the linked list
	middle := midNode(head)
	newHead := middle.Next
	middle.Next = nil

	newHead = reverseList(newHead)

	c1 := head
	c2 := newHead
	var tmp1, tmp2 *ListNode
	for c1 != nil && c2 != nil {
		tmp1 = c1.Next
		tmp2 = c2.Next

		c1.Next = c2
		c2.Next = tmp1

		c1 = tmp1
		c2 = tmp2
	}
}

func midNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
