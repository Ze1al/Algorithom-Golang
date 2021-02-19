/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/8
 */

package tencent

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if reverseNum(x) == x {
		return true
	} else {
		return false
	}
}

func reverseNum(x int) int {
	stack := []int{}
	for x != 0 {
		temp := x % 10
		stack = append(stack, temp)
		x = x / 10
	}
	res := 0
	for _, v := range stack {
		res = res * 10 + v
	}
	return res
}