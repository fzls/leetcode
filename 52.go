package leetcode

// 2019/10/28 23:52 by fzls
func totalNQueens(n int) int {
	total := 0

	var (
		columns     = make([]bool, n)
		leftAngles  = make([]bool, 2*n-1)
		rightAngles = make([]bool, 2*n-1)
	)
	totalNQueensCore(n, 0, columns, leftAngles, rightAngles, &total)

	return total
}

func totalNQueensCore(n int, row int, columns, leftAngles, rightAngles []bool, pTotal *int) {
	if row == n {
		*pTotal++
		return
	}

	for col := 0; col < n; col++ {
		if canPutQueen(n, row, col, columns, leftAngles, rightAngles) {
			columns[col] = true
			leftAngles[row-col+n-1] = true
			rightAngles[row+col] = true

			totalNQueensCore(n, row+1, columns, leftAngles, rightAngles, pTotal)

			columns[col] = false
			leftAngles[row-col+n-1] = false
			rightAngles[row+col] = false
		}
	}
}

func canPutQueen(n, row, col int, columns, leftAngles, rightAngles []bool) bool {
	return !(columns[col] || leftAngles[row-col+n-1] || rightAngles[row+col])
}
