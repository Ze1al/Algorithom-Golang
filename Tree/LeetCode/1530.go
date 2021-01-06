/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/5
	好叶子节点对的数量
 */

package LeetCode

func countPairs(root *TreeNode, distance int) int {
	_, res := dfs(root, distance)
	return res
}


func dfs(root *TreeNode, distance int) ([]int, int) {
	depths := make([]int, distance+1)
	var isLeaf bool
	if root.Right == nil && root.Left == nil {
		isLeaf = true
	}
	if isLeaf{
		depths[0] = 1
		return depths, 0
	}
	leftDepths, rightDepths := make([]int, distance+1), make([]int, distance+1)
	leftCount, rightCount := 0, 0

	if root.Left != nil {
		leftDepths, leftCount = dfs(root.Left, distance)
	}
	if root.Right != nil {
		rightDepths, rightCount = dfs(root.Right, distance)
	}

	for i := 0; i < distance; i++ {
		depths[i + 1] += leftDepths[i]
		depths[i + 1] += rightDepths[i]
	}

	cnt := 0
	for i := 0 ; i<distance; i++{
		for j := 0; j < distance -i-1;j++{
			cnt += leftDepths[i] + rightDepths[j]
		}
	}
	return depths, cnt  + leftCount + rightCount
}
