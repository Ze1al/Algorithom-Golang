/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/25
	有效的数独
	保证1~9在每一行 每一列 每一个九宫格中都只存在一个
	row col j/3 + (i/3)*3 以这三个为维度
*/

package Top

func isValidSudoku(board [][]byte) bool {
	var row, cols, sbox [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			sboxIndex := j/3 + (i/3)*3
			curentNum := board[i][j] - '1'
			if row[i][curentNum] == 1 {
				return false
			} else {
				row[i][curentNum] = 1
			}

			if cols[j][curentNum] == 1 {
				return false
			} else {
				cols[j][curentNum] = 1
			}

			if sbox[sboxIndex][curentNum] == 1 {
				return false
			} else {
				sbox[sboxIndex][curentNum] = 1
			}
		}
	}
	return true
}
