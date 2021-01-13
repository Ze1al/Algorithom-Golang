/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/8
	剪绳子
 */

package offer


func cuttingRopeRecursion(n int) int {
	if n == 2 {
		return 1
	}
	var res int
	for i := 0; i < n; i++ {
		res = max(i*(n-i), i*cuttingRopeRecursion(n-i), res)
	}
	return res
}

func max(x, y, z int) int {
	if x < y && y <z {
		return z
	} else if x < z && z < y{
		return y
	} else {
		return x
	}
}









