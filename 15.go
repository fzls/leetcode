package leetcode

import "sort"

type threeSumKey struct {
	n1, n2, n3 int
}

// 2019/09/17 1:18 by fzls
func threeSum(nums []int) [][]int {
	// 预处理一遍数组，保证所有数字最多出现三次，且数组元素递增
	sortedNums := preProcess(nums)
	if len(sortedNums) < 3 {
		return nil
	}
	if sortedNums[len(sortedNums)-1] < 0 {
		return nil
	}

	resSet := make(map[threeSumKey]struct{})

	for i := 0; i < len(sortedNums) && sortedNums[i] <= 0; i++ {
		j := i + 1
		k := len(sortedNums) - 1
		for j < k {
			sum := sortedNums[i] + sortedNums[j] + sortedNums[k]
			if sum == 0 {
				// 使用集合进行排重
				resSet[threeSumKey{sortedNums[i], sortedNums[j], sortedNums[k]}] = struct{}{}
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	// 取出数据
	var resList [][]int
	for res := range resSet {
		resList = append(resList, []int{res.n1, res.n2, res.n3})
	}
	sort.Slice(resList, func(i, j int) bool {
		if resList[i][0] != resList[j][0] {
			return resList[i][0] < resList[j][0]
		}
		if resList[i][1] != resList[j][1] {
			return resList[i][1] < resList[j][1]
		}
		return resList[i][2] < resList[j][2]
	})

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

	sort.Ints(res)
	return res
}
