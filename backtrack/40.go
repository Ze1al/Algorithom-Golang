/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/1
	全排列
*/

package backtrack

import "sort"

var res2 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res2 = [][]int{}
	sort.Ints(candidates)
	backtrack2(candidates, target, []int{})
	return res2
}

func backtrack2(candidates []int, target int, path []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res2 = append(res2, temp)
		return
	}

	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		backtrack2(candidates[i+1:], target-candidates[i], path)
		path = path[:len(path)-1]
	}
}
