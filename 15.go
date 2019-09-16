package leetcode

import "sort"

type threeSumKey struct {
	n1, n2, n3 int
}

// 2019/09/17 1:18 by fzls
func threeSum(nums []int) [][]int {
	// 若数字少于三个，则必然无解
	if len(nums) < 3 {
		return nil
	}
	// 预处理一遍数组，保证所有数字最多出现三次，且数组元素递增
	sortedNums := preProcess(nums)
	// 若预处理后，最小的值大于0或者最大的值小于0，则显然无解
	if sortedNums[0] > 0 || sortedNums[len(sortedNums)-1] < 0 {
		return nil
	}

	resSet := make(map[threeSumKey]struct{})

	// 固定一位，其余的按照two sum的处理
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
