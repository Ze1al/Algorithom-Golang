/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/30
	子集
	所有可能的子集
	回溯

result = []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}
*/

package backtrack

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













