/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/6
	判断是不是平衡二叉树
 */

package LeetCode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if abs(left - right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}


