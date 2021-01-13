/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package LeetCode


func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

