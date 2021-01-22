/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/14
 */

package tencent

func isPalindrome(x int) bool {
	return string(rune(x)) == reverseString(string(rune(x)))
}

func reverseString(s string) string {
	s1 := []rune(s)
	for i, j := 0, len(s1)-1; i<j; {
		s1[i], s1[j] = s1[j], s1[i]
		i++
		j--
	}
	return string(s1)
}
