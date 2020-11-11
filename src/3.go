/*
*  @author: didichuxing.com
 */

package src


// "abcabcbb"
// 维护一个window

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	for i, _ := range s {
		temp := []string{}
		for j := i; j < len(s); j++ {
			if inArray(string(s[j]), temp) == false {
				temp = append(temp, string(s[j]))
			} else {
				break
			}

		}
		if len(temp) > maxLength {
			maxLength = len(temp)
		}
	}
	return maxLength
}

func inArray(s string, arr []string) (res bool){
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}