package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var cachedHeight = make(map[*TreeNode]int)

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

	leftHeight := 0
	if height, ok := cachedHeight[root.Left]; ok {
		leftHeight = height
	} else {
		leftHeight = getHeight(root.Left)
		cachedHeight[root.Left] = leftHeight
	}
	rightHeight := 0
	if height, ok := cachedHeight[root.Right]; ok {
		rightHeight = height
	} else {
		rightHeight = getHeight(root.Right)
		cachedHeight[root.Right] = rightHeight
	}

	if leftHeight >= rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}
