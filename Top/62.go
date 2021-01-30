/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/26
	不同路径
	机器人从左上角走到右下角的方案有几种
	动态规划 i, j代表坐标
	f(i, j) = f(i, j-1) + f(i-1, j)

	1. 边缘状态
	f (0, j) = 1
	f (i, 0) = 1

	2. 状态转义方程
	f(i, j) = f(i, j-1) + f(i-1, j)
*/

package Top

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
