package leetcode

import (
	"sort"
)

// 2019/09/17 1:18 by fzls
func threeSum(nums []int) [][]int {
	// 若数字少于三个，则必然无解
	if len(nums) < 3 {
		return nil
	}
	// 预处理一遍数组，保证数组元素递增
	sort.Ints(nums)
	// 若预处理后，最小的值大于0或者最大的值小于0，则显然无解
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return nil
	}

	var resList [][]int

	// 固定一位，其余的按照two sum的处理
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 避免重复
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				// 使用集合进行排重
				resList = append(resList, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				// 去除重复解：这里参考了他人的做法
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
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

	return resList
}
