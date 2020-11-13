/*
*  @author: didichuxing.com
 */

package src

import "math"

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		index := math.Abs(float64(nums[i]))
		if nums[int(index)] > 0 {
			nums[int(index)] *= -1
		} else{
			return int(index)
		}
	}
	return -1
}
