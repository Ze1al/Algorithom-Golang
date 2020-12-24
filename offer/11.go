/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/23
	旋转数组
	找到旋转数组中最小的那个值
	二分的方向解题
*/

package offer

func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		piovt := low + (high-low)/2
		if numbers[piovt] > numbers[high] {
			low = piovt + 1
		} else if numbers[piovt] < numbers[high] {
			high = piovt
		} else {
			high--
		}
	}
	return numbers[low]
}
