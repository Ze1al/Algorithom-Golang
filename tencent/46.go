/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	全排列

resule = []
func backtrack(路径， 选择列表) {
	if 满足条件 {
		resule.add(path)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack（路径， 选择列表）
		撤销选择
	}
}


*/

package tencent

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	visit := make([]int, len(nums))
	backtrack([]int{}, nums, len(nums), visit)
	return res
}

func backtrack(path []int, nums []int, length int, visit []int) {
	if len(path) == length {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visit[i] == 1 {
			continue
		}

		visit[i] = 1
		path = append(path, nums[i])
		backtrack(path, nums, length, visit)
		path = path[:len(path)-1]
		visit[i] = 0
	}
}
















