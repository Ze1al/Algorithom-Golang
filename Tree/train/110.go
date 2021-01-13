/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/11
	判断是不是平衡二叉树
 */

package train

func isBalance(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(Height(root.Left)-Height(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}


func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(Height(root.Left), Height(root.Right)) + 1
}


func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func abs(x int) int {
	if x < 0 {
		return (-1) * x
	} else {
		return x
	}
}