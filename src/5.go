/*
*  @author: didichuxing.com
 */

package src

// 计算回文子串，最长的长度

/*
输入："babad"
输出："bab"
*/

func longestPalindrome(s string) string {
	maxLen := 0
	var maxString string
	if len(s) <= 1 {
		return s
	}
	for point1 := 0; point1 < len(s); point1++ {
		for point2 := point1+1; point2 <= len(s); point2++ {
			if isReverse(s[point1:point2]) && (point2-point1)>maxLen {
				maxLen = point2-point1
				maxString = s[point1:point2]
			}
		}
	}
	return maxString
}

func isReverse(s string) bool {
	if reverserString(s) == s {
		return true
	}
	return false
}


func reverserString(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
