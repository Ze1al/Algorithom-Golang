/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/15
	找出数组中重复的数字
*/

package offer

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range(nums) {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v] = 1
		}
	}
	return -1
}
