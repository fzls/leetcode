package leetcode

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%v", root.Val)}
	}

	// 先获取左右子节点的结果
	var childPaths []string
	childPaths = append(childPaths, binaryTreePaths(root.Left)...)
	childPaths = append(childPaths, binaryTreePaths(root.Right)...)

	// 然后最终结果是在所有结果前面加上root的值
	var res []string
	for _, path := range childPaths {
		res = append(res, fmt.Sprintf("%v->%v", root.Val, path))
	}

	return res
}
