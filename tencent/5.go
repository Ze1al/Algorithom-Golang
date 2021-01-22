/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/14
	回文字符串，使用dp
	P(i, j) 代表从i到j是不是回文字符
	1. 状态转移方程：P(i,j) = P(i+1, j-1) ^ S(i) == S(j)
	2. 初始值： if i == j {P[i, j] = true}

*/

package tencent

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	// 定义dp，dp[i, j]代表从i到j是不是回文字符
	dp := make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n)
	}
	// dp 对角线为true
	for i := 0; i< len(dp); i++{
		dp[i][i] = true
	}
	res := string(s[0])
	for j := 1; j<n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j - i <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] == true && j-i+1 > len(res){
				res = s[i:j+1]
			}
		}
	}
	return res
}
