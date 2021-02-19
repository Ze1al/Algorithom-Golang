/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/19
	二叉树中，最大路径和
	可以从子节点到父节点

	1. 以该节点作为根节点计算其最大的值
	2. 回到原来的树，这个节点的最大值是该节点+左右节点的最大值
 */

package tencent

import "math"

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt64
	maxGain(root)
	return maxSum
}


func maxGain(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(maxGain(root.Left), 0)
	right := max(maxGain(root.Right), 0)

	price := root.Val + left + right

	maxSum = max(maxSum, price)
	return root.Val + max(left, right)
}