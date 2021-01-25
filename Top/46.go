/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/25
	全排列

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

package Top

var res2 [][]int

func permute(nums []int) [][]int {
	n := len(nums)
	visit := make([]int, n)
	res2 = [][]int{}
	backtrack1(nums, []int{}, visit)
	return res2
}


func backtrack1(nums []int, temp []int, visit []int) {
	if len(nums) == len(temp) {
		trace := make([]int, len(temp))
		copy(trace, temp)
		res2 = append(res2, trace)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visit[i] == 1 {
			continue
		}
		visit[i] = 1
		temp = append(temp, nums[i])
		backtrack1(nums, temp, visit)
		visit[i] = 0
		temp = temp[:len(temp)-1]
	}
}
