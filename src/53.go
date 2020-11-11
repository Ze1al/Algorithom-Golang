/*
*  @author: didichuxing.com
 */

// 给定数组，求连续和最大

package src

// [-2,1,-3,4,-1,2,1,-5,4]
// 6

func maxSubArray(nums []int) int {
	maxNum := nums[0]
	f := make([]int, len(nums), len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		if f[i] > maxNum {
			maxNum = f[i]
		}
	}
	return maxNum
}

func max (i, j int) int {
	if i < j {
		return j
	}
	return i
}
