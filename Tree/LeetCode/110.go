/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
	平衡二叉树
 */

package LeetCode

func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Right) - height(root.Left)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}



func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	return max1(left, right) + 1
}


func max1(i, j int) int {
	if i<j{
		return j
	} else {
		return i
	}
}

func abs(x int) int {
	if x < 0{
		return x * (-1)
	} else {
		return x
	}
}