/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/25
	找出数组中能够满足相加等于target的所有组合
	所有组合 回溯
 */

package Top

import "sort"

var res3 [][]int

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	// clear res
	res3 = [][]int{}
	backTrack([]int{}, candidates, target)
	return res3
}


func backTrack(temp []int, candidates []int, target int) {
	if target == 0{
		track := make([]int, len(temp))
		copy(track, temp)
		res3 = append(res3, track)
	}

	for i, v := range candidates {
		if target - v < 0 {
			return
		}
		backTrack(append(temp, candidates[i]), candidates[i:], target-v)
	}
}

