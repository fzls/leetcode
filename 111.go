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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	md := math.MaxInt32
	if root.Left != nil {
		md = int(math.Min(float64(minDepth(root.Left)), float64(md)))
	}
	if root.Right != nil {
		md = int(math.Min(float64(minDepth(root.Right)), float64(md)))
	}

	return 1 + md
}
