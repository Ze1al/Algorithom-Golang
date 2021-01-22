/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/19
	最长回文子串
	使用dp

 */

package Top

import "strings"

func _longestPalindrome(s string) string {
	if len(s) <= 1{
		return s
	}
	maxLen := 0
	res := ""
	for i := 0; i < len(s); i++ {
		for j := i+1; j <= len(s); j++ {
			if isEqual(s[i:j]) && j-i > maxLen {
				maxLen = j-i
				res = s[i:j]
			}
		}
	}
	return res
}

func reverseString(s string) string {
	res := []string{}
	for i:= len(s)-1; i>=0; i-- {
		res = append(res, string(s[i]))
	}
	return strings.Join(res, "")
}

func isEqual(s string) bool{
	reverseString := reverseString(s)
	if reverseString == s {
		return true
	}
	return false
}


func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][i] = 1
	}
	res := string(s[0])
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = 0
			} else {
				if j -i <= 2{
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && len(res) < j-i+1 {
				res = s[i:j+1]
			}
		}
	}
	return res
}