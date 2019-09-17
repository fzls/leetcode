package leetcode

import "math"

// 2019/09/17 23:58 by fzls
func getAbsDiff(sum, target int) int {
	if sum > target {
		return sum - target
	} else {
		return target - sum
	}
}

func threeSumClosest(nums []int, target int) int {
	// 暴力解法
	minDiff := math.MaxInt32
	res := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				diff := getAbsDiff(sum, target)
				if diff < minDiff {
					minDiff = diff
					res = sum
				}
			}
		}
	}

	return res
}
