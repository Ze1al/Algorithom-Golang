/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/1
	组合总和
 */

package backtrack

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	backtrack(candidates, target, []int{})
	return res
}

func backtrack(candidates []int, target int, path []int) {
	if target < 0{
		return
	}
	if target == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}

	for i:= 0; i<len(candidates); i++ {
		path = append(path, candidates[i])
		backtrack(candidates[i:], target-candidates[i], path)
		path = path[:len(path)-1]
	}
}