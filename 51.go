package leetcode

// 2019/10/24 22:42 by fzls
func solveNQueens(n int) [][]string {
	qb := NewQueenBoard(n)
	qb.Solve()

	return qb.Results
}

type QueenBoard struct {
	N     int
	Board *BitMap

	Rows        *BitMap
	Colomns     *BitMap
	LeftAngles  *BitMap // /
	RightAngles *BitMap // \

	Results [][]string
}

func NewQueenBoard(n int) *QueenBoard {
	board := make([][]bool, n)
	data := make([]bool, n*n)
	for i := 0; i < n; i++ {
		board[i], data = data[:n], data[n:]
	}
	return &QueenBoard{
		N:           n,
		Board:       newBitMap(n * n),
		Rows:        newBitMap(n),
		Colomns:     newBitMap(n),
		LeftAngles:  newBitMap(2*n - 1),
		RightAngles: newBitMap(2*n - 1),
	}
}

func (qb *QueenBoard) CanPut(row, col int) bool {
	return !(qb.Rows.IsSet(row) || qb.Colomns.IsSet(col) || qb.LeftAngles.IsSet(row+col) || qb.RightAngles.IsSet(row-col+(qb.N-1)))
}

func (qb *QueenBoard) Put(row, col int) {
	qb.Board.Set(row*qb.N + col)
	qb.Rows.Set(row)
	qb.Colomns.Set(col)
	qb.LeftAngles.Set(row + col)
	qb.RightAngles.Set(row - col + (qb.N - 1))
}

func (qb *QueenBoard) Remove(row, col int) {
	qb.Board.UnSet(row*qb.N + col)
	qb.Rows.UnSet(row)
	qb.Colomns.UnSet(col)
	qb.LeftAngles.UnSet(row + col)
	qb.RightAngles.UnSet(row - col + (qb.N - 1))
}

func (qb *QueenBoard) Solve() {
	qb.solveCore(0)
}

func (qb *QueenBoard) solveCore(row int) {
	if row == qb.N {
		// 得到一个合法的结果，保存
		var res []string
		for i := 0; i < qb.N; i++ {
			line := make([]byte, qb.N)
			for j := 0; j < qb.N; j++ {
				if qb.Board.IsSet(i*qb.N + j) {
					line[j] = 'Q'
				} else {
					line[j] = '.'
				}
			}
			res = append(res, string(line))
		}
		qb.Results = append(qb.Results, res)
		return
	}

	// TODO：看看能否剪枝

	for col := 0; col < qb.N; col++ {
		if !qb.CanPut(row, col) {
			continue
		}

		qb.Put(row, col)
		qb.solveCore(row + 1)
		qb.Remove(row, col)
	}
}

// 优化内存使用
type BitMap struct {
	N     int
	Bytes []byte
}

func newBitMap(n int) *BitMap {
	return &BitMap{
		N:     n,
		Bytes: make([]byte, (n+7)/8),
	}
}

func (bm *BitMap) getByteAndBitIndex(n int) (byteIndex int, bitIndex uint) {
	return n / 8, uint(n % 8)
}

func (bm *BitMap) IsSet(n int) bool {
	bt, idx := bm.getByteAndBitIndex(n)
	return bm.Bytes[bt]&(1<<(7-idx)) != 0
}

func (bm *BitMap) Set(n int) {
	bt, idx := bm.getByteAndBitIndex(n)
	bm.Bytes[bt] |= 1 << (7 - idx)
}

func (bm *BitMap) UnSet(n int) {
	bt, idx := bm.getByteAndBitIndex(n)
	bm.Bytes[bt] &= ^(1 << (7 - idx))
}
