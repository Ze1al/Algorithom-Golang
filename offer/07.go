/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/23
	重建二叉树
	前序： preorder 中左右
	中序： inorder 左中右
 */

package offer

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:preorder[0],
	}

	var index int
	for i := range inorder {
		if inorder[i] == preorder[0]{
			index = i
			break
		}
	}
	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}






