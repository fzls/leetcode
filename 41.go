package leetcode

import "sort"

// 2019/10/08 0:17 by fzls
func firstMissingPositive(nums []int) int {
	// 先用不完美的方法做出来
	sort.Ints(nums)
	minNotExist := 1

	for idx, num := range nums {
		if num <= 0 {
			continue
		}
		if idx != 0 && nums[idx-1] == num {
			continue
		}
		if num == minNotExist {
			minNotExist++
		} else {
			return minNotExist
		}
	}

	return minNotExist
}
