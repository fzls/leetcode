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
	for start != nil {
		// 找到本轮的k个节点区间
		end := start
		for i := 1; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			// 剩余节点数不足k个，则不做额外处理
			break
		}

		// 从start到end进行翻转
		__reverse(pre, start, end, end.Next)

		// 迭代
		pre = start
		start = pre.Next
	}

	return dummy.Next
}

// 反转列表，使得pre start ... end post ...变为pre end ... start post ...
func __reverse(pre, start, end, post *ListNode) {
	// pre start ... end post ...
	if end == nil || start == nil {
		return
	}

	dummy := &ListNode{Next: start}

	_pre := dummy
	current := start
	for current != post {
		next := current.Next

		current.Next = _pre

		_pre = current
		current = next
	}

	// now: pre->start|end ... start|end->post ..., 连接起来->pre end ... start post ...
	pre.Next = end
	start.Next = post
}
