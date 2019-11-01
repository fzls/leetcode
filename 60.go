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
	// check
	if !(n >= 1 && n <= 9 && k >= 1 && k <= fact[n]) {
		panic("invalid params")
	}

	// init
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	// find each number from left to right until k is 1
	idx := 0
	//fmt.Println(idx, k, nums)
	for k != 1 {
		rightPerm := fact[n-idx-1]

		// find the first number by repeated reducing k by multiple of fact[n-1]
		toPutLeftIdx := idx
		for k > rightPerm {
			toPutLeftIdx++
			k -= rightPerm
		}
		// now move first number to front and those before it to right
		toPutLeft := nums[toPutLeftIdx]
		for i := toPutLeftIdx; i > idx; i-- {
			nums[i] = nums[i-1]
		}
		nums[idx] = toPutLeft

		idx++
		//fmt.Println(idx, k, nums)
	}

	// transform
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = '0' + byte(nums[i])
	}
	return string(res)
}
