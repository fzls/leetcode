package leetcode

// 2019/09/20 23:22 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	pre := dummy
	first := pre.Next
	for first != nil && first.Next != nil {
		second := first.Next
		third := second.Next

		// swap first and second
		pre.Next = second
		second.Next = first
		first.Next = third

		// move forward
		pre = pre.Next.Next
		first = pre.Next
	}

	return dummy.Next
}
