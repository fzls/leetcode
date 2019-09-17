package leetcode

import (
	"math"
	"sort"
)

// 2019/09/17 23:58 by fzls
func getAbsDiff(sum, target int) int {
	if sum > target {
		return sum - target
	} else {
		return target - sum
	}
}

func threeSumClosest(nums []int, target int) int {
	// 若数字少于三个，则必然无解
	if len(nums) < 3 {
		return 0
	}

	// 优化版本
	minDiff := math.MaxInt32
	res := math.MaxInt32

	// 预处理一遍数组，保证数组元素递增
	sort.Ints(nums)

	// 固定一位，其余的按照two sum的处理
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 避免重复计算
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}

			diff := getAbsDiff(sum, target)
			if diff < minDiff {
				minDiff = diff
				res = sum
			}
			if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				// 去除重复解：这里参考了他人的做法
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}

	return res
}
