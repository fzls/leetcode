package leetcode

// 2019/09/26 1:48 by fzls
func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	counter := make(map[int]int)
	mostNums := []int{nums[0]}
	for _, num := range nums {
		maxCnt := counter[mostNums[0]]

		counter[num]++
		cnt := counter[num]
		if cnt > maxCnt {
			mostNums = []int{num}
		} else if cnt == maxCnt {
			mostNums = append(mostNums, num)
		}
	}

	minLen := len(nums) + 1

	for _, num := range mostNums {
		i := 0
		j := len(nums) - 1
		for nums[i] != num {
			i++
		}
		for nums[j] != num {
			j--
		}
		if j-i+1 < minLen {
			minLen = j - i + 1
		}
	}

	return minLen
}
