package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func genLinkedList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummyHead := &ListNode{}
	p := dummyHead

	for _, num := range nums {
		p.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		p = p.Next
	}

	p = dummyHead.Next
	dummyHead.Next = nil
	return p
}

func linkedList2List(head *ListNode) (nums []int) {
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return
}
