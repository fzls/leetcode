package leetcode

// 2019/09/20 23:40 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}

	pre := dummy
	start := pre.Next
OUTER_LOOP:
	for start != nil {
		// 找到本轮的k个节点区间
		end := start
		for i := 1; i < k; i++ {
			if end == nil {
				// 剩余节点数不足k个，则不做额外处理
				break OUTER_LOOP
			}
			end = end.Next
		}

		// 从start到end进行翻转
		__reverse(pre, start, end)

		// 迭代
		pre = start
		start = pre.Next
	}

	return dummy.Next
}

// 反转列表，使得pre start ... end post ...变为pre end ... start post ...
func __reverse(pre, start, end *ListNode) {
	// pre start ... end post ...
	if end == nil || start == nil {
		return
	}

	post := end.Next

	dummy := &ListNode{Next: start}

	_pre := dummy
	current := start
	for current != post {
		next := current.Next

		current.Next = _pre

		_pre = current
		current = next
	}

	// now ->  pre|end ... start|post ...
	pre.Next = end
	start.Next = post
}
