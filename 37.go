package leetcode

import "fmt"

// 2019/10/03 1:43 by fzls
func makeOccuranceSetList() [][]bool {
	var setList [][]bool
	for i := 0; i < 9; i++ {
		setList = append(setList, make([]bool, 10))
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
				digit := board[i][j] - '0'
				rows[i][digit] = true
				columns[j][digit] = true
				boxes[getBoxIndex(i, j)][digit] = true
			}
		}
	}

	//printBoard(board)

	// 开始搜索可行解
	solveSudokuCore(board, 0, 0, rows, columns, boxes)

	//printBoard(board)
}

func solveSudokuCore(board [][]byte, row, column int, rows, columns, boxes [][]bool) bool {
	if row == 9 && column == 0 {
		// found solution
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
		boxIdx := getBoxIndex(row, column)
		for digit := 1; digit <= 9; digit++ {
			if rows[row][digit] || columns[column][digit] || boxes[boxIdx][digit] {
				continue
			}

			// mark
			board[row][column] = byte(digit) + '0'
			rows[row][digit] = true
			columns[column][digit] = true
			boxes[boxIdx][digit] = true

			// search
			if solved := solveSudokuCore(board, nextRow, nextColumn, rows, columns, boxes); solved {
				return true
			}

			// unmark
			board[row][column] = '.'
			rows[row][digit] = false
			columns[column][digit] = false
			boxes[boxIdx][digit] = false
		}
	} else {
		// 这个格子已经填过了
		if solved := solveSudokuCore(board, nextRow, nextColumn, rows, columns, boxes); solved {
			return true
		}
	}

	return false
}
