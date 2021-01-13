/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package train


func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == right {
		return root
	} else if left > right {
		return subtreeWithAllDeepest(root.Left)
	} else {
		return subtreeWithAllDeepest(root.Right)
	}
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max1(left, right) + 1
}


func max1(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}