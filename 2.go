package leetcode

//
////Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	carry := 0
	// 当两个链表都没有到底的时候，同时推进，并依次相加，计算溢出值
	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{}
		current = current.Next
		current.Val = sum % 10
		carry = sum / 10
	}

	// 两条链表均处理完后若仍有溢出，则额外增加一个节点
	if carry != 0 {
		current.Next = &ListNode{}
		current = current.Next

		current.Val = carry
	}

	return dummyHead.Next
}
