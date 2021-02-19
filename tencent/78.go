/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	子集

result = []

func backtrack(path， nums) {
	if 满足条件 {
		result.add(path)
	}
	for i in nums {
		做选择
		backtrack(path, nums)
		撤销选择
	}
}
*/

package tencent

var res4 [][]int

// 给定一个元素
func subsets(nums []int) [][]int {
	res4 = [][]int{}
	backTrack1([]int{}, nums, 0)
	return res4
}

func backTrack1(temp []int, nums []int, start int) {

	track := make([]int, len(temp))
	copy(track, temp)
	res4 = append(res4, track)

	for i := start; i < len(nums); i++{
		temp = append(temp, nums[i])
		backTrack1(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}
