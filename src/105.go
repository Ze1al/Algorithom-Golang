/*
*  @author: didichuxing.com
 */

package src

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}else if len(preorder) == 1 {
		return &TreeNode{Val:preorder[0]}
	} else {
		return &TreeNode{
			Val:preorder[0],
			Left:buildTree(preorder[1:findIndex(inorder, preorder[0])+1], inorder[:findIndex(inorder, preorder[0])]),
			Right:buildTree(preorder[findIndex(inorder, preorder[0])+1:], inorder[findIndex(inorder, preorder[0])+1:]),
		}
	}
}

func findIndex(nums []int, target int) int {
	for i, v := range nums{
		if v == target {
			return i
		}
	}
	return 0
}