/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/20
	盛最多水的容器
	输入：[1,8,6,2,5,4,8,3,7]
 */

package Top

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j{
		if height[i] < height[j] {
			temp := (j - i) * height[i]
			if temp > res {
				res = temp
			}
			i++
		} else {
			temp := (j - i) * height[j]
			if temp > res {
				res = temp
			}
			j--
		}
	}
	return res
}