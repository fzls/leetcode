package leetcode

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	lastRow := make([]int, rowIndex+1)
	// row 1
	lastRow[0] = 1
	lastRow[1] = 1
	currentRow := make([]int, rowIndex+1)

	for i := 2; i <= rowIndex; i++ {
		currentRow[0] = 1
		currentRow[i] = 1

		for j := 1; j < i; j++ {
			// j>=1, i>=2
			currentRow[j] = lastRow[j-1] + lastRow[j]
		}

		for j := 0; j <= i; j++ {
			lastRow[j] = currentRow[j]
		}
	}

	return currentRow
}
