/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/19
	只出现一次的值
	使用异或
	a ^ 0 = a
	a ^ a = 0
 */

package tencent

func singleNumber(nums []int) int {
	x := 0
	for _, v := range nums {
		x = x ^ v
	}
	return x
}