/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/8
 */

package tencent


func reverse(x int) int {
	ans := 0
	for x != 0 {
		temp := x % 10
		ans = ans * 10 + temp
		x = x / 10
	}
	if !(-(1<<31) - 1 <= ans && ans <= (1<<31)-1) {
		return 0
	} else {
		return ans
	}
}