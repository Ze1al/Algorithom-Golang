/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/6
	求和路径
 */

package LeetCode


func pathSum(root *TreeNode, sum int)  int {
	var res int
	if root == nil {
		return res
	}
	res = dfs1(root, sum)
	if root.Left != nil {
		res += pathSum(root.Left, sum)
	}
	if root.Right != nil {
		res += pathSum(root.Right, sum)
	}
	return res
}

func dfs1(node *TreeNode, sum int)  int {
	var res int
	if sum - node.Val == 0{
		res ++
	}
	if node.Left != nil {
		res += dfs1(node.Left, sum - node.Val)
	}
	if node.Right != nil {
		res += dfs1(node.Right, sum - node.Val)
	}
	return res
}

