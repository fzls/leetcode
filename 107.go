package leetcode

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) Push(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	elem := q.list.Front()
	q.list.Remove(elem)

	return elem.Value
}

func (q *Queue) Len() int {
	return q.list.Len()
}

type NodeWithDepth struct {
	node  *TreeNode
	depth int
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	// 广度优先遍历
	queue := NewQueue()
	queue.Push(&NodeWithDepth{node: root, depth: 1})

	for queue.Len() != 0 {
		current := queue.Pop().(*NodeWithDepth)
		if current.node == nil {
			continue
		}
		if len(res) < current.depth {
			res = append(res, []int{})
		}
		res[current.depth-1] = append(res[current.depth-1], current.node.Val)

		queue.Push(&NodeWithDepth{node: current.node.Left, depth: current.depth + 1})
		queue.Push(&NodeWithDepth{node: current.node.Right, depth: current.depth + 1})
	}

	// reverse list
	i := 0
	j := len(res) - 1
	for i < j {
		temp := res[i]
		res[i] = res[j]
		res[j] = temp

		i++
		j--
	}

	return res
}
