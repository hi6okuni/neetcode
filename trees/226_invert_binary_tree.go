package trees

// Given the root of a binary tree, invert the tree, and return its root.

// Example 1:

// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
// Example 2:

// Input: root = [2,1,3]
// Output: [2,3,1]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// func invertTree(root *TreeNode) *TreeNode {
// 	if root != nil {
// 		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
// 	}
// 	return root
// }

// BFS
type Queue []*TreeNode

func (q *Queue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Pop() *TreeNode {
	node := (*q)[0]
	*q = (*q)[1:]
	return node
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := Queue{root}

	for !queue.IsEmpty() {
		node := queue.Pop()

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue.Push(node.Left)
		}

		if node.Right != nil {
			queue.Push(node.Right)
		}
	}

	return root
}
