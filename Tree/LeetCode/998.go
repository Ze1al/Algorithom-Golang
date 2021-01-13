/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/5
	最大的二叉树
 */

package LeetCode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val{
		temp := &TreeNode{Val:val}
		temp.Left = root
		return temp
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

