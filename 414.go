package leetcode

func thirdMax(nums []int) int {
	var maxNums []int

	for _, num := range nums {
		if len(maxNums) == 0 {
			maxNums = append(maxNums, num)
			continue
		}

		isBreaked := false
		for i := 0; i < len(maxNums); i++ {
			if num < maxNums[i] {
				continue
			}

			if num > maxNums[i] {
				if len(maxNums) < 3 {
					maxNums = append(maxNums, 0)
				}
				for j := len(maxNums) - 1; j > i; j-- {
					maxNums[j] = maxNums[j-1]
				}
				maxNums[i] = num
			}

			isBreaked = true
			break
		}

		if !isBreaked && len(maxNums) < 3 {
			maxNums = append(maxNums, num)
		}
	}

	if len(maxNums) < 3 {
		return maxNums[0]
	}
	return maxNums[2]
}
