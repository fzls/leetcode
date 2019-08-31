package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	if !(leftHeight-rightHeight >= -1 && leftHeight-rightHeight <= 1) {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(getHeight(root.Left)), float64(getHeight(root.Right))))
}
