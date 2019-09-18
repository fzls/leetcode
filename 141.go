package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 使用双指针法
	p1 := head      // step 1
	p2 := head.Next // step 2
	for p2 != nil {
		// 每一轮p2会比p1多走一格，所以若有环，则必定相遇
		if p2 == p1 {
			return true
		}

		// p1 走一步
		p1 = p1.Next

		// p2走两步
		p2 = p2.Next
		// 边缘检查
		if p2 == nil {
			return false
		}
		p2 = p2.Next
	}

	return false
}
