/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package LeetCode


func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := height(root.Left)
	right := height(root.Right)
	if left == right {
		return root
	} else if left < right {
		return lcaDeepestLeaves(root.Right)
	} else {
		return lcaDeepestLeaves(root.Left)
	}
}