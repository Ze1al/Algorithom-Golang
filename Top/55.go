/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/26
	跳跃游戏
	给一个数组，每个数字代表可以跳跃的最大距离，判断是否可以跳到最后一个位置
 */

package Top

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	maxLength := 0
	for i, v := range nums {
		if maxLength >= i && (i+v)>maxLength {
			maxLength = i + v
		}
	}
	return maxLength >= len(nums)-1
}