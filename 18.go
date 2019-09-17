package leetcode

import (
	"fmt"
	"sort"
)

// 2019/09/18 0:44 by fzls
func fourSum(nums []int, target int) [][]int {
	// 直接按照三和的思路写一遍
	// 若数字少于四个，则必然无解
	if len(nums) < 4 {
		return nil
	}
	// 预处理一遍数组，保证数组元素递增
	sort.Ints(nums)
	fmt.Println(nums)

	var resList [][]int

	// 固定一位，其余的按照two sum的处理
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 避免重复
		}

		for p := i + 1; p < len(nums); p++ {
			if p > i+1 && nums[p] == nums[p-1] {
				continue
			}

			j := p + 1
			k := len(nums) - 1
			for j < k {
				sum := nums[i] + nums[p] + nums[j] + nums[k]
				if sum == target {
					// 使用集合进行排重
					resList = append(resList, []int{nums[i], nums[p], nums[j], nums[k]})
					j++
					k--

					// 去除重复解：这里参考了他人的做法
					for j < k && nums[j] == nums[j-1] {
						j++
					}
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else if sum > target {
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
	}

	return resList
}
