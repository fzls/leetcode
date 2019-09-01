package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	remain := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return remain == 0
	}

	return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
