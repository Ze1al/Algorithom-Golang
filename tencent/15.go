/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	三数字之和等于0
	定义三个指针吧
	L = i + 1
	R = n - 1
 */

package tencent

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		L := i+1
		R := len(nums)-1
		if i>= 1 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0{
			return res
		}
		for L < R {
			if nums[i] + nums[L] + nums[R] == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				for (L<R) && nums[L] == nums[L+1] {
					L += 1
				}
				for (L<R) && nums[R] == nums[R-1] {
					R -= 1
				}
				L += 1
				R -= 1

			} else if nums[i] + nums[L] + nums[R] > 0 {
				R -= 1
			} else {
				L += 1
			}
		}
	}
	return res
}


