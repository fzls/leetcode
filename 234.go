package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome__(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 看了其他人的解法，发现可以通过把后半段给翻转来判断是否是回文
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var rightStart *ListNode
	if fast == nil {
		rightStart = slow.Next
	} else {
		rightStart = slow.Next.Next
	}

	// reverse rightStart
	rightStart = _reverseList(rightStart)
	for rightStart != nil {
		if head.Val != rightStart.Val {
			return false
		}

		head = head.Next
		rightStart = rightStart.Next
	}

	return true
}

func _reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	next := head.Next
	for cur != nil && next != nil {
		temp := next.Next
		next.Next = cur

		cur = next
		next = temp
	}
	head.Next = nil

	return cur
}
