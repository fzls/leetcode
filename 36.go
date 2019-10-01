package leetcode

const (
	_ROWS    = 9
	_COLUMNS = 9
	_GONG_GE = 9
)

// 2019/10/02 1:52 by fzls
func isValidSudoku(board [][]byte) bool {
	if len(board) != _ROWS {
		return false
	}
	for i := 0; i < _ROWS; i++ {
		if len(board[i]) != _COLUMNS {
			return false
		}
	}

	used := make([]bool, 10)
	checkUsed := func(row, column int) bool {
		if board[row][column] == '.' {
			return false
		}
		num := board[row][column] - '0'
		if used[num] {
			return true
		}
		used[num] = true
		return false
	}
	// 判断每一行
	for row := 0; row < _ROWS; row++ {
		reset(used)
		for column := 0; column < _COLUMNS; column++ {
			if checkUsed(row, column) {
				return false
			}
		}
	}

	// 判断每一列
	for column := 0; column < _COLUMNS; column++ {
		reset(used)
		for row := 0; row < _ROWS; row++ {
			if checkUsed(row, column) {
				return false
			}
		}
	}

	// 判断每一个宫格
	for i := 0; i < _GONG_GE; i++ {
		reset(used)
		startRow := 3 * (i / 3)
		startColumn := 3 * (i % 3)
		for row := startRow; row < startRow+3; row++ {
			for column := startColumn; column < startColumn+3; column++ {
				if checkUsed(row, column) {
					return false
				}
			}
		}
	}

	return true
}

func reset(used []bool) {
	for i := 0; i < len(used); i++ {
		used[i] = false
	}
}
