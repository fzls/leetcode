package leetcode

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	//i := 0
	//for j := 0; j < len(nums); j++ {
	//	if nums[j] != val {
	//		// 将不等于val的值依次放到前方
	//		nums[i] = nums[j]
	//		i++
	//	}
	//}

	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return i
}
