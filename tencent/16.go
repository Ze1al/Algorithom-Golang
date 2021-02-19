/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	最接近的三数之和
*/

package tencent

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := 0
	if len(nums) <= 2 {
		return res
	}
	sort.Ints(nums)
	gap := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if abs(sum - target) < gap {
				gap = abs(sum - target)
				res = sum

			} else if sum > target {
				R -= 1
			} else if sum < target {
				L += 1
			} else {
				return res
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}