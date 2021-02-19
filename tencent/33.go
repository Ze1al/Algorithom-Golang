/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	旋转数组中搜索target
*/

package tencent

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	pre := 0
	last := len(nums)-1
	for pre <= last {
		mid := (pre + last) / 2
		if nums[mid] == target {
			return mid
		}

		// 需要判断那一边有序

		// 左边有序
		if nums[0] <= nums[mid] {
			if nums[0] <= target && nums[mid] >= target {
				last = mid - 1
			} else {
				pre = mid + 1
			}
			// 右边有序
		} else {
			if nums[mid] <= target && nums[last] >= target {
				pre = mid + 1
			} else {
				last = mid - 1
			}
		}
	}
	return -1
}

func binarySearch(nums []int, target int) int {
	pre := 0
	last := len(nums)
	for pre < last {
		mid := (pre + last) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			last = mid
		} else {
			pre = mid + 1
		}
	}
	return -1
}
