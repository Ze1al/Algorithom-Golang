/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/8
	盛水最多的容器
*/

package tencent

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	area := 0
	for i := 0; i < len(height); i++ {
		for j := i+1; j < len(height); j++ {
			high := min(height[i], height[j])
			if high * (j - i) > area {
				area = high * (j - i)
			}
		}
	}
	return area
}


//func min(x, y int) int {
//	if x < y {
//		return x
//	} else {
//		return y
//	}
//}
