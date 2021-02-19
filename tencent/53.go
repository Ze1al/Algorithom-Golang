/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	最大的子序和
	动态规划
	f(i) = max(f(i) + sum, f(i) )

 */

package tencent


func maxSubArray(nums []int) int {
	maxTarget := nums[0]
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		if f[i] > maxTarget {
			maxTarget = f[i]
		}
	}
	return maxTarget
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}



