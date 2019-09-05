package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	next := cur.Next
	for next != nil {
		// 保存下下个节点
		nextNext := next.Next

		// 将下个节点指向本节点
		next.Next = cur

		// 迭代
		cur = next
		next = nextNext
	}
	// 将原头结点的Next指向nil
	head.Next = nil

	return cur
}
