package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for headA != nil {
		// 判断A是否在B后面
		for p := headB; p != nil; p = p.Next {
			if p == headA {
				return headA
			}
		}

		// A前进一格
		headA = headA.Next
	}

	return nil
}
