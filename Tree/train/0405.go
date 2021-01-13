/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	合法二叉树
 */

package train

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return helper(root, math.MinInt16, math.MaxInt16)
}


func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val < lower || root.Val > upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Left, root.Val, upper)
}