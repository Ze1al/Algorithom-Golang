/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
 */

package tencent


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}