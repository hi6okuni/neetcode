package trees

import (
	"reflect"
	"testing"
)

// テストヘルパー関数：スライスからバイナリツリーを構築
func BuildTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(nodes) {
		node := queue[0]
		queue = queue[1:]
		if i < len(nodes) && nodes[i] != nil {
			node.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(nodes) && nodes[i] != nil {
			node.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// テストヘルパー関数：バイナリツリーをスライスに変換
func TreeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}
	result := []interface{}{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	// 末尾のnilを削除
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected []interface{}
	}{
		{
			name:     "Example 1",
			input:    []interface{}{4, 2, 7, 1, 3, 6, 9},
			expected: []interface{}{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Example 2",
			input:    []interface{}{2, 1, 3},
			expected: []interface{}{2, 3, 1},
		},
		{
			name:     "Empty Tree",
			input:    []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "Single Node",
			input:    []interface{}{1},
			expected: []interface{}{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := BuildTree(tt.input)
			inverted := invertTree(root)
			result := TreeToSlice(inverted)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("invertTree() = %v, want %v", result, tt.expected)
			}
		})
	}
}
