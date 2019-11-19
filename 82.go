package leetcode

// 2019/11/19 22:11 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func _deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}

	current := dummy
	next := head
	for next != nil {
		nn := next.Next
		if nn != nil && nn.Val == next.Val {
			for nn != nil && nn.Val == next.Val {
				nn = nn.Next
			}
		} else {
			current.Next = &ListNode{Val: next.Val}
			current = current.Next
		}
		next = nn
	}

	return dummy.Next
}
