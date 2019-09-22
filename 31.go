package leetcode

import "sort"

// 2019/09/22 23:10 by fzls
// 最小的排列是完全顺序的，最大的排列是完全逆序的，也就是说排列递增的过程其实是让逆序对增加的过程，所以可以考虑转化为如果使当前排列的逆序对增加最小值
func nextPermutation(nums []int) {
	last := len(nums) - 1
	for i := last - 1; i >= 0; i-- {
		// 找到第一个连续顺序队 i, i+1
		if nums[i] < nums[i+1] {
			// 从后往前找到第一个与该顺序对左元素构成顺序对的元素
			j := last
			for j > i && !(nums[j] > nums[i]) {
				j--
			}
			// 交换这组逆序对
			nums[j], nums[i] = nums[i], nums[j]
			// 将剩余部分排序，使该部分逆序数最小
			sort.Ints(nums[i+1:])
			return
		}
	}

	// 完全逆序
	i := 0
	j := last
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
