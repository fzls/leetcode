package leetcode

// 2019/11/07 22:18 by fzls
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	zeroRows := make([]bool, m)
	zeroCols := make([]bool, n)

	// find row and col that should zero out
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				zeroRows[row] = true
				zeroCols[col] = true
			}
		}
	}

	// zero rows
	for row := 0; row < m; row++ {
		if !zeroRows[row] {
			continue
		}

		for col := 0; col < n; col++ {
			matrix[row][col] = 0
		}
	}

	// zero cols
	for col := 0; col < n; col++ {
		if !zeroCols[col] {
			continue
		}

		for row := 0; row < m; row++ {
			matrix[row][col] = 0
		}
	}

	return
}
