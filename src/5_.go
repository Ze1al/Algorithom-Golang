/*
*  @author: didichuxing.com
 */

package src

func _longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	// 创建二维数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// break
			if j > i {
				break
			}
			// length = 0
			if i == j {
				dp[i][j] = 1
			}
			// length = 1
			if i-j == 1 && s[i] == s[i] {
				dp[i][j] = 1
			}
			// length = 2
			//if i-j >= 2 && s[i-1][j+1] == 1 {
			//	dp[i][j] = 1
			//}
			if dp[i][j] == 1 && i-j+1 > len(ans) {
				ans = s[j : i+1]
			}
		}
	}
	return ans

}
