package leetcode

func plusOne(digits []int) []int {
	addon := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		sum := digit + addon

		digits[i] = sum % 10
		addon = sum / 10
	}

	if addon == 0 {
		return digits
	} else {
		return append([]int{1}, digits...)
	}
}
