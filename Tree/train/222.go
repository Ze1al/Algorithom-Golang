/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	计算完全二叉树
 */

package train

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left + right + 1
}




