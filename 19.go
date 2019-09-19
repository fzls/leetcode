package leetcode

// 2019/09/20 0:04 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	// 使用快慢指针法找到倒数第N+1个节点
	pNFast := dummyHead
	for i := 0; i < n+1; i++ {
		pNFast = pNFast.Next
	}

	lastN1 := dummyHead
	for pNFast != nil {
		pNFast = pNFast.Next
		lastN1 = lastN1.Next
	}

	// 现在问题转化为已知链表的头结点，和其中某一个节点的前一个节点，删除这个节点
	lastN1.Next = lastN1.Next.Next
	return dummyHead.Next
}
