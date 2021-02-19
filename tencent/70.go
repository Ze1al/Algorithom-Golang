/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	爬楼梯
*/

package tencent

func climbStairs(n int) int {
	res := []int{1, 2}
	if n <= 2 {
		return res[n-1]
	}

	for i := 0; i < n-2; i++ {
		temp := res[0] + res[1]
		res[0] = res[1]
		res[1] = temp
	}
	return res[len(res)-1]
}
