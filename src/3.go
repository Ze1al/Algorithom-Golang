/*
*  @author: didichuxing.com
 */

package src


// "abcabcbb"
// 维护一个window

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1{
		return len(s)
	}
	maxLen := 0
	for i, v := range s {
		window := []string{}
		for k := i+1; i < len(s); i++ {
			if InArray(string(s[k]), window) {
				window = append(window, string(v))
			} else {
				break
			}
		}
		if len(window) > maxLen {
			maxLen = len(window)
		}
	}
	return maxLen
}

func InArray(n string, nums []string) (result bool) {
	for _, v := range nums {
		if n == v {
			return true
		}
	}
	return false
}