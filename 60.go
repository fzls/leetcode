package leetcode

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
	checkParams(n, k)

	nums := getInitalPerm(n)
	batchPerm(k, nums)
	return strPerm(nums)
}

func checkParams(n int, k int) {
	if !(n >= 1 && n <= 9 && k >= 1 && k <= fact[n]) {
		panic("invalid params")
	}
}

func getInitalPerm(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return nums
}

// batch perm by step size of fact[i](1<=i<=n-1])
func batchPerm(k int, nums []int) {
	n := len(nums)

	// find each number from left to right until k is 1
	idx := 0
	for k != 1 {
		batchSize := fact[n-(idx+1)] // number of perms for idx+1,idx+2,...n-2,n-1

		// find the first number by repeated reducing k by multiple of fact[n-1]
		targetIdx := idx + k/batchSize
		k = k % batchSize
		if k == 0 {
			k = batchSize
			targetIdx--
		}
		// now move target number to front and those before it to right
		target := nums[targetIdx]
		// copy idx, idx+1 ... targetIdx-1 to idx+1, idx+2 ... targetIdx
		for i := targetIdx; i >= idx+1; i-- {
			nums[i] = nums[i-1]
		}
		nums[idx] = target

		idx++
	}
}

func strPerm(nums []int) string {
	n := len(nums)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = '0' + byte(nums[i])
	}
	return string(res)
}
