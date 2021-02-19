/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/8
 */

package tencent


func longestPalindrome(s string) string {
	maxLen := 0
	res := ""
	if len(s) == 0 || len(s) == 1{
		return s
	}
	for i := 0; i < len(s); i++ {
		for j := i+1; j <= len(s); j++ {
			temp := s[i:j]
			if reverseString(temp) == temp  && (j - i) > maxLen{
				res = temp
				maxLen = j - i
			}
		}
	}
	return res
}

func reverseString (s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}