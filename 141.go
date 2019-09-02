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

	// 最直观的算法：使用hashmap记录出现过的节点，若重复，则说明有环
	nodeSet := make(map[*ListNode]struct{})
	for head != nil {
		if _, exists := nodeSet[head]; exists {
			return true
		}

		nodeSet[head] = struct{}{}
		head = head.Next
	}

	return false
}
