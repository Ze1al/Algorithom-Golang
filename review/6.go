/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
	平衡二叉树
 */

package review

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(height(root.Left) - height(root.Right)) > 1{
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(x ,y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func abs(x int)int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}