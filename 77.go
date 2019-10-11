package leetcode

// 2019/10/11 23:39 by fzls
func combine(n int, k int) [][]int {
	var res [][]int

	//combineCore(n, k, []int{}, &res)
	//combineDictOrder(n, k, &res)
	combineOrder(n, k, &res)

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

	//fmt.Println(*pRes)
	return
}

// 我个人认为的字典序（参照其他博客，感觉这个更像字典序：两个字符串比较时，从左向右，找到第一个不相等的来进行排序，所以应该是从右到左第一个可以增加的+1
func combineOrder(n int, k int, pRes *[][]int) {
	nums := make([]int, k)
	// 初始化为 1,2,3...,k-1,k
	for i := 0; i < k; i++ {
		nums[i] = i + 1
	}

	hasNew := true
	for hasNew {
		// 保存一个结果
		res := make([]int, k)
		copy(res, nums)
		*pRes = append(*pRes, res)

		// 从1,2,3,4升序生成3,4,5,6
		// 字典序：所有组合的元素都是升序的，且后一个组合是当前组合的字典序后序
		// 在上面的约束下：k个元素的组合的第i个元素的最大可能值为n-(k-i)(i=1,2,...k)
		// 找到下一个字典序集合：从右到左，找到第一个第一个小于该位置最大值的元素，将其加1得到下一个排序，若找不到这样的元素，说明已找到所有组合
		// 如n=6， k=4， 2,3,4,6=> 2,3,5,6
		hasNew = false
		for i := k - 1; i >= 0; i-- {
			if nums[i] <= n-(k-i) {
				nums[i]++
				for j := i + 1; j < k; j++ {
					nums[j] = nums[i] + j - i
				}

				hasNew = true
				break
			}
		}
	}

	//fmt.Println(*pRes)
	return
}
