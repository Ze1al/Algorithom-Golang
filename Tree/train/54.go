/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/12
	二叉树的第k大节点
 */

package train

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	res := inorderTraversal(root)
	return res[len(res)-k]
}
