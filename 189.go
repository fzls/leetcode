package leetcode

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	// 使用编程之美中的逆转法
	// 1~n-k n-k+1~n
	// n-k~1 n~n-k+1
	// n-k+1~n 1~n-k
	k = k % len(nums)
	_reverse(nums, 0, len(nums)-k-1)
	_reverse(nums, len(nums)-k, len(nums)-1)
	_reverse(nums, 0, len(nums)-1)
}

func _reverse(nums []int, b int, e int) {
	for b < e {
		temp := nums[b]
		nums[b] = nums[e]
		nums[e] = temp

		b++
		e--
	}
}
