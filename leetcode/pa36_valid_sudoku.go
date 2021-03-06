package main

func isValidSudoku(board [][]byte) bool {
	if board == nil || len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	row := make([]map[byte]bool, 9)
	col := make([]map[byte]bool, 9)
	square := make([]map[byte]bool, 9)
	for i, b := range board {
		for j := range b {
			x := board[i][j]
			if x == '.' {
				continue
			} else {
				r := row[i]
				if r == nil {
					r = make(map[byte]bool)
					row[i] = r
				}

				c := col[j]
				if c == nil {
					c = make(map[byte]bool)
					col[j] = c
				}

				k := (i / 3)  + 3 * (j/3)
				s := square[k]
				if s == nil {
					s = make(map[byte]bool)
					square[k]=s
				}
				if _, ok := row[i][x]; ok {
					return false
				} else {
					row[i][x] = true
				}
				if _, ok := col[j][x]; ok {
					return false
				} else {
					col[j][x] = true
				}
				if _, ok := square[k][x]; ok {
					return false
				} else {
					square[k][x] = true
				}
			}
		}
	}
	return true
}
