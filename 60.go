package leetcode

import (
	"sort"
)

var fact = [...]int{
	1,      // 0
	1,      // 1
	2,      // 2
	6,      // 3
	24,     // 4
	120,    // 5
	720,    // 6
	5040,   // 7
	40320,  // 8
	362880, // 9

}

// 2019/11/01 23:28 by fzls
func getPermutation(n int, k int) string {
	// check
	if !(n >= 1 && n <= 9 && k >= 1 && k <= fact[n]) {
		panic("invalid params")
	}

	// init
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// prune
	if n > 1 {
		// find the first number by repeated reducing k by multiple of fact[n-1]
		first := 1
		for k > fact[n-1] {
			first++
			k -= fact[n-1]
		}
		// now move first number to front and those before it to right
		for i := first - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = first
	}

	//iter（this is first)
	//fmt.Println(n, k)
	//fmt.Println(nums)
	for i := 1; i < k; i++ {
		nextPerm(nums)
		//fmt.Println(nums)
	}
	//fmt.Println()

	// transform
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = '0' + byte(nums[i])
	}
	return string(res)
}

func nextPerm(nums []int) {
	n := len(nums)
	// from right to left, find first i such that nums[i] < nums[i+1]
	left := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			left = i
			break
		}
	}
	if left == -1 {
		panic("k is invalid")
	}

	// find j such that nums[j] is the minim elem that large than nums[i]
	toSwap := n - 1 - sort.Search(n-1-left, func(i int) bool {
		return nums[n-1-i] > nums[left]
	})
	//fmt.Println("left=", left, "swap=", toSwap)

	// swap nums[i] and nums[j]
	nums[left], nums[toSwap] = nums[toSwap], nums[left]

	// because nums[i] < nums[i+1] and nums[k] > nums[k+1] for all k > i, we know that nums[j-1] > nums[i], and nums[i] > nums[j+1]
	// so after swap, nums[i+1:] is sorted desc, we only need to reverse it to make the minim next perm
	i := left + 1
	j := n - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]

		i++
		j--
	}
}
