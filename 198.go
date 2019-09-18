package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 假设f(n)为前n个房子中能获取的最大值，f(0)=0, f(1)=nums[0]
	// 则有f(k) = max(f(k-1), nums[k-1] + f(k-2))
	maxArr := make([]int, len(nums)+1)
	maxArr[0] = 0
	maxArr[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		maxArr[i] = maxArr[i-1]
		if nums[i-1]+maxArr[i-2] > maxArr[i] {
			maxArr[i] = nums[i-1] + maxArr[i-2]
		}
	}
	return maxArr[len(nums)]
}
