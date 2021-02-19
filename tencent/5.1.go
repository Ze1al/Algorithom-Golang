/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/8
 */

package tencent

func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	res := string(s[0])
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[j] != s[i] {
				dp[i][j] = 0
			} else {
				if j - i <= 2 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && len(res) < j-i+1{
				res = s[i:j+1]
			}
 		}
	}
	return res
}
