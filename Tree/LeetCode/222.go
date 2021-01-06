/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/6
 */

package LeetCode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left + right + 1
}



