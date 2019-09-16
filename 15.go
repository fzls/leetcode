package leetcode

import "sort"

type threeSumKey struct {
	n1, n2, n3 int
}

// 2019/09/17 1:18 by fzls
func threeSum(nums []int) [][]int {
	// 预处理一遍数组，保证所有数字最多出现三次
	nums = preProcess(nums)

	resSet := make(map[threeSumKey]struct{})

	for i := 0; i < len(nums)-2; i++ {
		// 转化为two sum问题
		if twoNumbers := _twoSum(nums[i+1:], -nums[i]); len(twoNumbers) != 0 {
			for _, twoNumber := range twoNumbers {
				// 排序方便排重
				ts := append([]int{nums[i]}, twoNumber...)
				sort.Ints(ts)
				// 使用集合进行排重
				resSet[threeSumKey{ts[0], ts[1], ts[2]}] = struct{}{}
			}
		}
	}

	// 取出数据
	var resList [][]int
	for res := range resSet {
		resList = append(resList, []int{res.n1, res.n2, res.n3})
	}

	return resList
}

func preProcess(nums []int) []int {
	cntMap := make(map[int]int)
	var res []int
	for _, num := range nums {
		cntMap[num]++
		if cntMap[num] <= 3 {
			res = append(res, num)
		}
	}

	return res
}

// 返回所有值加起来等于target的一对数
func _twoSum(nums []int, target int) [][]int {
	var res [][]int

	cache := make(map[int]int) // 值=>此值出现的下标

	for idx, num := range nums {
		if targetIdx, exists := cache[target-num]; exists {
			res = append(res, []int{nums[targetIdx], num})
		}

		cache[num] = idx
	}

	return res
}
