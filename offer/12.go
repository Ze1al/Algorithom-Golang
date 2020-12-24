/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/23
从二维数组中找到一条路径和word一样
*/

package offer

func exist(board [][]byte, word string) bool {
	if word == "" {
		return false
	}
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		if dfs(board, i+1, j, k+1, word) || dfs(board, i-1, j, k+1, word) || dfs(board, i, j+1, k+1, word) || dfs(board, i, j-1, k+1, word) {
			return true
		} else {
			board[i][j] = temp
		}
	}
	return false

}
