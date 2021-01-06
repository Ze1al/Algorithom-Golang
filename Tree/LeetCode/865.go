/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/6
	具有所有最深节点的最小子树
 */

package LeetCode


func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == right {
		return root
	} else if left < right {
		return subtreeWithAllDeepest(root.Right)
	} else {
		return subtreeWithAllDeepest(root.Left)
	}
}