/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/21
	三数之和
	给定一个n的数组，使其中存在三个元素a+b+c = 0
*/

package Top

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < n-1; i++ {
		if nums[i] > 0 {
			break
		}
		L := i + 1
		R := n-1
		for L < R {
			if nums[i] + nums[L] + nums[R] == 0 {
				res = append(res, []int{i, L, R})
			} else if nums[i] + nums[L] + nums[R] > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}
