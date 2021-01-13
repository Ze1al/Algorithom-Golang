/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/7
  从根到叶的二进制数之和
 */

package LeetCode


func sumRootToLeaf(root *TreeNode) int {
	return dfs2(root, 0)
}

func dfs2(root *TreeNode, n int) int {
	if root == nil {
		return 0
	}
	n = n * 2 + root.Val
	if root.Left == nil && root.Right == nil {
		return n
	}
	return (dfs2(root.Left, n) + dfs2(root.Right, n)) % (10 ^ 9 + 7)
}