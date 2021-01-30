/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/30
	矩阵置零
	如果某一行某一列为0则该行该列全为0
*/

package Top

func setZeroes(matrix [][]int) {
	flagRaw := false
	flagCol := false

	raw := len(matrix)
	col := len(matrix[0])

	for i := 0; i < raw; i++ {
		if matrix[i][0] == 0 {
			flagRaw = true
		}
	}

	for j := 0; j < col; j++ {
		if matrix[0][j] == 0 {
			flagCol = true
		}
	}

	// 存储几行几列是为0
	for i := 1; i < raw; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0{
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 将原来的置为0
	for i := 1; i < raw; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0{
				matrix[i][j] = 0
			}
		}
	}

	// 第一行第一列情况
	if flagRaw {
		for i := 0; i < raw; i++ {
			matrix[i][0] = 0
		}
	}

	if flagCol {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}
}

