package leetcode

// 2019/11/07 22:18 by fzls
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	isFirstColumnHasZero := false
	for row := 0; row < m; row++ {
		if matrix[row][0] == 0 {
			isFirstColumnHasZero = true
		}

		for col := 1; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for col := 1; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if isFirstColumnHasZero {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}

	return
}
