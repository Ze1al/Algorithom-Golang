/*
*  @author: didichuxing.com
 */

/*
	爬楼梯
*/

package src

// 使用滚动数组的思想

func climbStairs(n int) int {
	var pre, mid, last int
	last = 1
	for i := 0; i < n; i++ {
		pre = mid
		mid = last
		last = pre + mid
	}
	return last
}

func _climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return _climbStairs(n-1) + _climbStairs(n-2)
}
