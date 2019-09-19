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
	// 使用快慢指针法找到倒数第N个节点
	pNFast := head
	i := 0
	for ; pNFast != nil && i < n; i++ {
		pNFast = pNFast.Next
	}
	// 列表长度小于n
	if pNFast == nil && i < n-1 {
		panic("输出不合法")
	}

	lastN := head
	for pNFast != nil {
		pNFast = pNFast.Next
		lastN = lastN.Next
	}

	// 现在问题转化为已知链表的头结点，和其中某一个节点，删除这个节点
	dummyHead := &ListNode{}
	dummyHead.Next = head

	if lastN.Next == nil {
		// 如果是末尾
		p := dummyHead
		for p.Next != lastN {
			p = p.Next
		}
		p.Next = nil
	} else {
		// 如果是中间
		next := lastN.Next
		lastN.Val = lastN.Next.Val
		lastN.Next = lastN.Next.Next
		next.Next = nil
	}
	return dummyHead.Next
}
