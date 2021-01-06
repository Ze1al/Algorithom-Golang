/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/5
	最大的二叉树
 */

package LeetCode

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val{
		temp := &TreeNode{Val:val}
		temp.Left = root
	}
	right := insertIntoMaxTree(root.Right, val)
	root.Right = right
	return root
}

