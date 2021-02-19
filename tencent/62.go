/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	不同路径
	动态规划
	dp[i][j] = dp[i-1][j] + dp[i][j-1]
*/

package tencent

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	// 边界条件
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 递推公式
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
