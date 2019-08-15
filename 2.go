package leetcode

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	current := res

	addon := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		current.Next = &ListNode{}
		current = current.Next

		sum := l1.Val + l2.Val + addon
		current.Val = sum % 10
		addon = sum / 10
	}

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

	if addon != 0 {
		current.Next = &ListNode{}
		current = current.Next

		current.Val = addon
	}

	return res.Next
}
