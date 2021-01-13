/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package train

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(hight(root.Left) - hight(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func hight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(hight(root.Left), hight(root.Right)) + 1
}

func maxHight(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}



