package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func recurReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	_head := head
	if head.Next != nil {
		// 循环将后续列表翻转
		_head = recurReverseList(head.Next)
		// 将下一个节点的Next指向本节点
		head.Next.Next = head
	}
	return _head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 循环将后续列表翻转
	res := recurReverseList(head)
	// 将原先的头结点的Next指向nil
	head.Next = nil
	return res
}
