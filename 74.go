package leetcode

// 2019/11/10 23:37 by fzls
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	// now we have m row, and n col
	// 1. if k in [0, m-1] and i < j, then a[k][i] < a[k][j]
	// 2. for k in [1, m-1], then a[k-1][*] < a[k][*]
	// so we can conduct that, if we iter number row by row, and from left to right in each row, the number will in acceding order, so we can use binary search
	low := 0
	high := m*n - 1

	for low <= high {
		mid := (low + high) / 2
		row, col := extractRowAndColFromIndex(mid, n)
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func extractRowAndColFromIndex(index int, itemPerRow int) (row, col int) {
	return index / itemPerRow, index % itemPerRow
}
