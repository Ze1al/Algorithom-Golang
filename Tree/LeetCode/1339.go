/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	分裂二叉树的最大乘积
*/

package LeetCode

func maxProduct(root *TreeNode) int {
	total := getSum(root)
	return dfs3(root, total)
}


func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	res += root.Val
	left := getSum(root.Left)
	right := getSum(root.Right)
	return res + left + right
}

func dfs3(node *TreeNode, total int) int {
	if node == nil {
		return 0
	}
	var res int

	left := dfs3(node.Left, total)
	right := dfs3(node.Right, total)
	subSum := left + right + node.Val
	res = max(res, subSum * (total-subSum))
	return res
}