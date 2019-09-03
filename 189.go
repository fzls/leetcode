package leetcode

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	// 一次一次移动
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j >= 1; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}
