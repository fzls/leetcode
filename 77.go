package leetcode

import "fmt"

// 2019/10/11 23:39 by fzls
func combine(n int, k int) [][]int {
	var res [][]int

	//combineCore(n, k, []int{}, &res)
	combineDictOrder(n, k, &res)

	return res
}

func combineCore(n int, k int, comb []int, pRes *[][]int) {
	if k == 0 {
		res := make([]int, len(comb))
		copy(res, comb)
		*pRes = append(*pRes, res)
		return
	}

	if n < k {
		return
	}

	combineCore(n-1, k, comb, pRes)
	combineCore(n-1, k-1, append(comb, n), pRes)
}

// 参考官方题解的字典序组合实现
func combineDictOrder(n int, k int, pRes *[][]int) {
	nums := make([]int, k+1)
	// 初始化为 1,2,3...,k-1,k, n+1(哨兵)
	for i := 0; i < k; i++ {
		nums[i] = i + 1
	}
	nums[k] = n + 1

	j := 0
	for j < k {
		// 保存一个结果
		res := make([]int, k)
		copy(res, nums)
		*pRes = append(*pRes, res)

		// 找到下一个字典序集合：将第一个不与下一个元素连续的位置加1，并将其左侧依次设为1,2,3...
		// 如n=6， k=4， 2,3,4,6,[7]=> 1,2,{5},6,[7]
		j = 0
		for j < k && nums[j]+1 == nums[j+1] {
			nums[j] = j + 1
			j++
		}
		nums[j]++
	}

	fmt.Println(*pRes)
	return
}
