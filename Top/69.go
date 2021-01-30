/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/30
	实现X的平方根
	使用二分法， 下界为0 上界为x
	y^2 < x
 */

package Top

func mySqrt(x int) int {
	left, right := 0, x
	res := -1

	for left < right {
		mid := left + (right - left) / 2
		if mid * mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res

}

