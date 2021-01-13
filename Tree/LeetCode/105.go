/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	从前序和中序遍历的序列构造二叉树

	前序遍历 preorder = [3,9,20,15,7]  中 左 右
	中序遍历 inorder = [9,3,15,20,7]	  左 中 右
 */

package LeetCode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}
	root := &TreeNode{Val:preorder[0]}
	index := findIndex(inorder, root.Val)

	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}


func findIndex(inorder []int, target int) int {
	for i:= 0; i<len(inorder); i++ {
		if inorder[i] == target {
			return i
		}
	}
	return 0
}