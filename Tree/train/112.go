/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package train

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}