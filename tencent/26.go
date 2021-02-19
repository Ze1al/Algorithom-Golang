/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	删除排序数组中的重复项
	原地修改
 */

package tencent


func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j<len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i
}


