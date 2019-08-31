package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for nextDiffent := head.Next; nextDiffent != nil; nextDiffent = nextDiffent.Next {
		if nextDiffent.Val != current.Val {
			current.Next = nextDiffent
			current = nextDiffent
		}
	}

	current.Next = nil

	return head
}
