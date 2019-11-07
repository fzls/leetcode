package leetcode

// 2019/11/07 22:18 by fzls
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	copyed := make([][]int, m)
	for i := 0; i < m; i++ {
		copyed[i] = make([]int, n)
		copy(copyed[i], matrix[i])
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if copyed[row][col] == 0 {
				for i := 0; i < m; i++ {
					matrix[i][col] = 0
				}
				for j := 0; j < n; j++ {
					matrix[row][j] = 0
				}
			}
		}
	}

	return
}
