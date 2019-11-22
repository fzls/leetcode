package leetcode

// 2019/11/22 22:50 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummyLess := &ListNode{}
	dummyGreateOrEqual := &ListNode{}

	lessCurrent, geCurrent := dummyLess, dummyGreateOrEqual
	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			lessCurrent.Next = node
			lessCurrent = node
		} else {
			geCurrent.Next = node
			geCurrent = node
		}
	}

	lessCurrent.Next = dummyGreateOrEqual.Next
	geCurrent.Next = nil
	return dummyLess.Next
}
