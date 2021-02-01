/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/1
 */

package backtrack


var res6 [][]int

// 给定一个元素
func subsetsWithDup(nums []int) [][]int {
	res6 = [][]int{}
	backTrack6([]int{}, nums, 0)
	return res6
}

func backTrack6(temp []int, nums []int, start int) {

	track := make([]int, len(temp))
	copy(track, temp)
	res6 = append(res6, track)

	for i := start; i < len(nums); i++{
		if i > start && nums[i-1] == nums[i] {
			continue
		}
		temp = append(temp, nums[i])
		backTrack6(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}







