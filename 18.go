package leetcode

import (
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

	// 预先剪枝
	n := len(nums)
	if nums[0]+nums[1]+nums[2]+nums[3] > target || nums[n-1]+nums[n-2]+nums[n-3]+nums[n-4] < target {
		return nil
	}

	var resList [][]int
	// 固定一位，其余的按照two sum的处理
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 避免重复
		}

		// 剪枝
		if i+3 < len(nums) && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		for p := i + 1; p < len(nums); p++ {
			if p > i+1 && nums[p] == nums[p-1] {
				continue
			}

			// 剪枝
			if p+2 < len(nums) && nums[i]+nums[p]+nums[p+1]+nums[p+2] > target {
				break
			}

			j := p + 1
			k := len(nums) - 1
			for j < k {
				// 剪枝
				if nums[i]+nums[p]+nums[j]+nums[j+1] > target {
					break
				}
				if nums[i]+nums[p]+nums[k-1]+nums[k] < target {
					break
				}

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
