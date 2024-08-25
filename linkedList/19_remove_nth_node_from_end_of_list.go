package linkedlist

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []
// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pointerHistories := []*ListNode{}
	current := head
	for current != nil {
		pointerHistories = append(pointerHistories, current)
		current = current.Next
	}

	if len(pointerHistories) == n {
		if n == 1 {
			return nil
		}
		return head.Next
	}

	justBeforeTarget := pointerHistories[len(pointerHistories)-n-1]
	justBeforeTarget.Next = justBeforeTarget.Next.Next

	return head
}
