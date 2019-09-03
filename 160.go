package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cache := make(map[*ListNode]struct{})
	for headA != nil {
		cache[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		if _, exists := cache[headB]; exists {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
