/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
求树的深度
 */

package LeetCode


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}


func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}