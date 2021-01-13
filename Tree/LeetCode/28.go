/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/7
	对称二叉树
 */

package LeetCode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return judge(root.Left, root.Right)
}


func judge(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val{
		return false
	}
	return judge(left.Left, right.Right) && judge(left.Right, right.Left)
}