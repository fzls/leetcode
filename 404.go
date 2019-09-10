package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func _sumOfLeftLeaves(root *TreeNode, isLeftChild bool) int {
	if root == nil {
		return 0
	}

	if isLeftChild && root.Left == nil && root.Right == nil {
		return root.Val
	}

	return _sumOfLeftLeaves(root.Left, true) +
		_sumOfLeftLeaves(root.Right, false)
}

func sumOfLeftLeaves(root *TreeNode) int {
	return _sumOfLeftLeaves(root, false)
}
