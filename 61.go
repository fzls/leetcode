package leetcode

// 2019/11/02 23:35 by fzls
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	// get linklist length
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	// check special case
	if n <= 1 {
		return head
	}
	k = k % n
	if k == 0 {
		return head
	}

	// Start
	// 1	 	...    	n-k 		n-k+1 		... 	n
	// head 	... 	newTail 	newHead		...		tail
	// Want
	// n-k+1 	... 	n 			1 			...		n-k
	// newHead	...		tail		head		...		newTail
	tail := head
	for i := 0; i < k; i++ {
		tail = tail.Next
	}
	newTail := head
	for tail.Next != nil {
		newTail = newTail.Next
		tail = tail.Next
	}
	newHead := newTail.Next

	tail.Next = head
	newTail.Next = nil

	return newHead
}
