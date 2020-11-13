/*
*  @author: didichuxing.com
 */

package src

// 给定数组，标记其中有哪些存在，哪些不存在
// 利用索引标记一部分信息
// 将值对应的索引上的值标记为负数，之后位置上大于0的就代表没有出现过

import "math"

func findDisappearedNumbers(nums []int) []int {
	// 将值对应的索引上的值标记为负数
	for i := 0; i < len(nums); i++ {
		index := math.Abs(float64(nums[i])) - 1
		if nums[int(index)] > 0 {
			nums[int(index)] *= -1
		}
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]>0{
			res = append(res, i+1)
		}
	}
	return res
}
