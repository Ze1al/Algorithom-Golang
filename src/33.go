/*
*  @author: didichuxing.com
 */

package src

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	pre := 0
	last := len(nums) - 1

	for pre <= last {
		mid := (pre + last) / 2
		if nums[mid] == target {
			return mid
		}

		// 判断那一边是有序
		// 左半部有序
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				last = mid - 1
			} else {
				pre = mid + 1
			}

		} else {
			if nums[mid] < target && target <= nums[last] {
				pre = mid + 1
			} else {
				last = mid - 1
			}
		}
	}
	return -1
}
