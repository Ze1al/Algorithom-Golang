/*
*  @author: didichuxing.com
	单词拆分
*/

package src

// 动态规划

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if isEqual(s[i:j], wordDict) && dp[i]==true{
				dp[j] = true
			}
		}
	}
	return dp[n]
}


func isEqual(s string, wordDict []string) bool {
	for _, k := range wordDict {
		if k == s{
			return true
		}
	}
	return false
}