package leetcode

// A[leftA|RightA]
// B[LeftB|RightB]

// i is the length of leftA, find length of leftB, that is J
func getJ(totalLen int, i int) int {
	if totalLen&1 == 0 {
		return totalLen/2 - i
	} else {
		return (totalLen-1)/2 - i
	}
}

func median(numsA []int, numsB []int, leftA, leftB int) float64 {
	if (len(numsA)+len(numsB))&1 == 0 {
		// 如果是偶数，则是下标有效时，max(A[i-1], B[j-1]) 与 min(A[i], B[j])的平均值
		var left, right int
		left = -1 << 31
		if leftA > 0 {
			left = numsA[leftA-1]
		}
		if leftB > 0 && numsB[leftB-1] > left {
			left = numsB[leftB-1]
		}

		right = 1<<31 - 1
		if leftA < len(numsA) {
			right = numsA[leftA]
		}
		if leftB < len(numsB) && numsB[leftB] < right {
			right = numsB[leftB]
		}

		return (float64(left) + float64(right)) / 2
	} else {
		// 如果是偶数，则是A[i]和B[j]中一个，取决于B[j-1]<=A[i]<=B[j]和A[i-1]<=B[j]<=A[i]哪个成立了
		var media int
		// 假设中位数在a数组中
		if leftA < len(numsA) {
			media = numsA[leftA]
			valid := true
			if leftB > 0 && media < numsB[leftB-1] {
				valid = false
			}
			if leftB < len(numsB) && media > numsB[leftB] {
				valid = false
			}
			if valid {
				return float64(media)
			}
		}

		// 假设中位数在b数组中
		if leftB < len(numsB) {
			media = numsB[leftB]
			valid := true
			if leftA > 0 && media < numsA[leftA-1] {
				valid = false
			}
			if leftA < len(numsA) && media > numsA[leftA] {
				valid = false
			}
			if valid {
				return float64(media)
			}
		}

		return 0
	}
}

// len(numsA) >= len(numsB)
func _findMedianSortedArrays(numsA []int, numsB []int) float64 {
	// 如果第二个数组为空，则可直接确定分片结果
	if len(numsB) == 0 {
		return median(numsA, numsB, len(numsA)/2, 0)
	}

	// 二分查找，找到合适的两个数组的分片，是的左侧两块数目与右侧两块一致，且左侧均小于右侧
	low := 0
	high := len(numsA)
	// i,j分别表示将AB数组切割后，左侧一半的数字数目
	var i, j int
	for low <= high {
		i = low + (high-low)/2
		// j随着i的增大而减小
		j = getJ(len(numsA)+len(numsB), i)
		// 需要确保j在合理的范围内，若不再，则此种组合不可能，需要调整i的上下限
		// 如果j太大了，则需要使i变大，也就是调整下限
		if j > len(numsB) {
			low = i + 1
			continue
		}
		// 如果j太小了，则需要使i变小，也就是调整上限
		if j < 0 {
			high = i - 1
			continue
		}

		// 如果j的范围合适，则判断A[i-1]和B[j] 以及A[i]和B[j-1]的关系
		// 当下标存在时，只有在A[i-1] <= B[j] 且 B[j-1] <= A[i]的时候有唯一解，否则需要调整上下限
		if i > 0 && j < len(numsB) && numsA[i-1] > numsB[j] {
			// 当不满足第一条时，说明A[i-1]太大，也就是i太大，所以需要调整上限
			high = i - 1
			continue
		}
		if j > 0 && j <= len(numsB) && i < len(numsA) && numsB[j-1] > numsA[i] {
			// 当不满足第二条时，说明A[i]太小，也就是i太小，所以需要调整下限
			low = i + 1
			continue
		}

		// 此处表示当A左侧取i个，B左侧取J个的时候，我们可以在他们附近找到中位数了
		break
	}

	// 按照一定规则找到中位数
	return median(numsA, numsB, i, j)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 保证第一个数组总是长于第二个数组，方便处理
	if len(nums1) >= len(nums2) {
		return _findMedianSortedArrays(nums1, nums2)
	} else {
		return _findMedianSortedArrays(nums2, nums1)
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}
