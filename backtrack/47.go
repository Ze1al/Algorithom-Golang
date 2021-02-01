/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/1
 */

package backtrack

import "sort"

var res5 [][]int

func permuteUnique(nums []int) [][]int {
	res5 = [][]int{}
	sort.Ints(nums)
	visit := make([]int, len(nums))
	backtrack4(nums, []int{}, len(nums), visit)
	return res5
}

func backtrack4(nums []int, path []int, length int, visit []int) {
	if len(path) == length {
		temp := make([]int, len(path))
		copy(temp, path)
		res5 = append(res5, temp)
		return
	}


	for i:=0; i<len(nums);i++{
		if visit[i] == 1 {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && visit[i-1] == 1{
			continue
		}
		visit[i] = 1
		path = append(path, nums[i])
		backtrack4(nums, path, length, visit)
		path = path[:len(path)-1]
		visit[i] = 0
	}
}