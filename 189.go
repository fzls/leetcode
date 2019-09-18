package leetcode

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	// 使用之前从陈越姥姥那边学到的环法
	k = k % len(nums)
	// 统计遍历了多少个数了
	cnt := 0
	for i := 0; i < len(nums) && cnt < len(nums); i++ {
		buf := nums[i]
		current := i
		for j := (i + len(nums) - k) % len(nums); j != i; j = (j + len(nums) - k) % len(nums) {
			nums[current] = nums[j]
			current = j
			cnt++
		}
		nums[current] = buf
		cnt++
	}
}
