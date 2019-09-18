package leetcode

func twoSum_(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			// 因为升序，所以肯定是因为i太小
			i++
		} else {
			// 否则是j太大了
			j--
		}
	}

	return nil
}
