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

	addon := 0
	// 当两个链表都没有到底的时候，同时推进，并依次相加，计算溢出值
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		current.Next = &ListNode{}
		current = current.Next

		sum := l1.Val + l2.Val + addon
		current.Val = sum % 10
		addon = sum / 10
	}

	// 推进还有元素未处理的列表
	for ; l1 != nil; l1 = l1.Next {
		current.Next = &ListNode{}
		current = current.Next

		sum := l1.Val + addon
		current.Val = sum % 10
		addon = sum / 10
	}

	for ; l2 != nil; l2 = l2.Next {
		current.Next = &ListNode{}
		current = current.Next

		sum := l2.Val + addon
		current.Val = sum % 10
		addon = sum / 10
	}

	// 两条链表均处理完后若仍有溢出，则额外增加一个节点
	if addon != 0 {
		current.Next = &ListNode{}
		current = current.Next

		current.Val = addon
	}

	return dummyHead.Next
}
