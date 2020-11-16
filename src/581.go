/*
*  @author: didichuxing.com
    给定一个数组，将其中的一个子数组排序，排序后整个数组是有序的
    比较排序排好的数组
 */

package src

import (
	"reflect"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	sortNum := make([]int, len(nums))
	copy(sortNum, nums)
	sort.Ints(sortNum)

	pre := 0
	last := len(nums) - 1

	if reflect.DeepEqual(nums, sortNum) {
		return 0
	}
	for pre < last && nums[pre] == sortNum[pre] {
		pre += 1
	}
	for pre<last && nums[last] == sortNum[last] {
		last -= 1
	}
	return last - pre + 1
}





