package leetcode

import "fmt"

// 2019/10/03 1:43 by fzls
func makeOccuranceSetList() []map[byte]struct{} {
	var setList []map[byte]struct{}
	for i := 0; i < 9; i++ {
		setList = append(setList, make(map[byte]struct{}))
	}
	return setList
}

func getBoxIndex(row, column int) int {
	return 3*(row/3) + column/3
}

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", board[i][j]-'0')
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func solveSudoku(board [][]byte) {
	// 输入保证是个9*9的，这里暂时不检查了

	rows := makeOccuranceSetList()
	columns := makeOccuranceSetList()
	boxes := makeOccuranceSetList()

	// 初始化已填数值的格子的信息到各个set中
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				digit := board[i][j]
				rows[i][digit] = struct{}{}
				columns[j][digit] = struct{}{}
				boxes[getBoxIndex(i, j)][digit] = struct{}{}
			}
		}
	}

	//printBoard(board)

	// 开始搜索可行解
	solveSudokuCore(board, 0, 0, rows, columns, boxes)

	//printBoard(board)
}

func solveSudokuCore(board [][]byte, row, column int, rows, columns, boxes []map[byte]struct{}) bool {
	if row == 9 && column == 0 {
		return true
	}

	var nextRow, nextColumn int
	if column == 8 {
		nextRow = row + 1
		nextColumn = 0
	} else {
		nextRow = row
		nextColumn = column + 1
	}

	if board[row][column] == '.' {
		// 需要搜索不同组合
		for digit := byte('1'); digit <= '9'; digit++ {
			if _, used := rows[row][digit]; used {
				continue
			}
			if _, used := columns[column][digit]; used {
				continue
			}
			if _, used := boxes[getBoxIndex(row, column)][digit]; used {
				continue
			}

			// mark
			board[row][column] = digit
			rows[row][digit] = struct{}{}
			columns[column][digit] = struct{}{}
			boxes[getBoxIndex(row, column)][digit] = struct{}{}

			// search
			if solved := solveSudokuCore(board, nextRow, nextColumn, rows, columns, boxes); solved {
				return true
			}

			// unmark
			board[row][column] = '.'
			delete(rows[row], digit)
			delete(columns[column], digit)
			delete(boxes[getBoxIndex(row, column)], digit)
		}
	} else {
		// 这个格子已经填过了
		if solved := solveSudokuCore(board, nextRow, nextColumn, rows, columns, boxes); solved {
			return true
		}
	}

	return false
}
