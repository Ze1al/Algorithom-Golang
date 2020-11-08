/*
*  @author: didichuxing.com
 */

package src

// two element of list sum

// [2, 7, 9, 8]
// target = 9

func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for k := index + 1; k<len(nums); k++ {
			if nums[k] + value == target {
				return []int{index, k}
			}
		}
	}
	return []int{}
}


