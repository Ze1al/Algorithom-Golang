/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/25
 */

package Top

func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft, maxRight := 0, 0
		for j:=0; j<i;j++ {
			maxLeft = max(maxLeft, height[j])
		}
		for j:=len(height)-1; j>i;j-- {
			maxRight = max(maxRight, height[j])
		}
		if maxLeft <= height[i] || maxRight<=height[i] {
			continue
		}

		res += min(maxLeft, maxRight) - height[i]
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}