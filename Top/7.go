/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/19
	整数反转
	以栈的思想处理
 */

package Top

func reverse(x int) int {
	var res int
	for x != 0 {
		res = res * 10 + x % 10
		x = x/10
	}
	if !(-(1<<31)-1 <= res && res <= (1<<31)-1) {
		return 0
	}
	return res
}

