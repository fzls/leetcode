package leetcode

func getInsertPosition(nums1 []int, currentLen int, val int) int {
	low := 0
	high := currentLen - 1
	for low <= high {
		j := low + (high-low)>>1
		if nums1[j] <= val {
			low = j + 1
		} else {
			high = j - 1
		}
	}

	return low
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	right := m
	for i := 0; i < n; i++ {
		// 找到插入的位置
		ip := getInsertPosition(nums1, right, nums2[i])

		// 插入
		for idx := right; idx > ip; idx-- {
			nums1[idx] = nums1[idx-1]
		}
		nums1[ip] = nums2[i]

		right++
	}
}
