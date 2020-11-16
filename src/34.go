/*
*  @author: didichuxing.com

找到数组中这个数的范围，要求O(logn)
*/

package src

func searchRange(nums []int, target int) []int {
	nLen := len(nums)
	if nLen <= 2 {
		if nLen == 0 {
			return []int{-1, -1}
		}
		if nLen == 1 && nums[0] == target {
			return []int{0,0}
		} else if nLen == 2{
			if nums[0] == target && nums[1] == target{
				return []int{0, 1}
			}
			if nums[0] == target{
				return []int{0, 0}
			}
			if nums[1] == target{
				return []int{1, 1}
			}
		}
		return []int{-1, -1}
	}

	pre := 0
	last := nLen
	res := []int{}
	for pre < last  && last - pre > 1{
		mid := (pre + last) / 2
		if nums[mid] == target {
			res = findNum(nums, mid)
			return res
		}
		if nums[mid] > target {
			last = mid
		} else {
			pre = mid
		}
	}
	if nums[pre] == target {
		return []int{pre, pre}
	}
	return res
}

func findNum(nums []int, index int) []int {
	pre := index
	last := index
	for pre > 0 && last < len(nums) && nums[pre] == nums[index] {
		pre -= 1
	}
	for pre > 0 && last < len(nums) && nums[last] == nums[index] {
		last += 1
	}
	res := []int{}
	for i := pre + 1; i < last; i++ {
		res = append(res, i)
	}
	return res
}
