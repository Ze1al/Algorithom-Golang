/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/30
	爬楼梯
	fibonacci(n) = fibonacci(n-1) + fibonacci(n+2)
  */

package Top

func climbStairs1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return climbStairs1(n-1) + climbStairs1(n-2)
	}
}


func climbStairs2(n int) int {
	res := []int{0, 1, 2}
	if n <= 2 {
		return res[n]
	}
	for i := 2; i < n; i++ {
		res = append(res, res[len(res)-1] + res[len(res)-2])
	}
	return res[n]
}


func climbStairs(n int) int {
	res := [2]int{1, 2}
	if n == 0 {
		return 0
	} else if n <= 2 {
		return res[n-1]
	} else {
		for i := 3; i <= n; i++ {
			temp := res[0] + res[1]
			res[0] = res[1]
			res[1] = temp
		}
		return res[len(res)-1]
	}
}