package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome__(head *ListNode) bool {
	// 先考虑直接的实现方法，将链表转换为数组
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] != nums[j] {
			return false
		}

		i++
		j--
	}

	return true
}
