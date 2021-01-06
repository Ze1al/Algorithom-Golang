/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/6
	二叉树剪枝
 */

package LeetCode

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if containsOne(root) {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	return root
}


func containsOne(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Val == 1{
		return false
	}
	return containsOne(node.Left) && containsOne(node.Right)
}


