/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	买卖股票的最佳时机2
	可以允许多次买卖
*/

package tencent

func maxProfit(prices []int) int {
	// 定义状态
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	// 定义边界条件
	dp[0][0] = 0
	dp[0][1] = -1 * prices[0]
	// 状态转移方程
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}
