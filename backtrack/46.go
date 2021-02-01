/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/1
	全排列
 */

package backtrack

var res3 [][]int

func permute(nums []int) [][]int {
	res3 = [][]int{}
	visit := make([]int, len(nums))
	backtrack3(nums, []int{}, len(nums), visit)
	return res3
}

func backtrack3(nums []int, path []int, length int, visit []int) {
	if len(path) == length {
		temp := make([]int, len(path))
		copy(temp, path)
		res3 = append(res3, temp)
		return
	}

	for i:=0; i<len(nums);i++{
		if visit[i] == 1 {
			continue
		}
		visit[i] = 1
		path = append(path, nums[i])
		backtrack3(nums, path, length, visit)
		path = path[:len(path)-1]
		visit[i] = 0
	}
}