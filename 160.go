package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	movea := false
	pb := headB
	moveb := false

	for pa != nil && pb != nil && pa != pb {
		pa = pa.Next
		pb = pb.Next

		if pa == nil && !movea {
			pa = headB
			movea = true
		}
		if pb == nil && !moveb {
			pb = headA
			moveb = true
		}
	}

	if pa == nil || pb == nil {
		return nil
	}

	return pa
}
