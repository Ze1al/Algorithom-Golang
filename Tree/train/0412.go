/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	字节跳动 校招面试题
 */

package train

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := dfs2(root, sum)
	if root.Left != nil {
		res += pathSum(root.Left, sum)
	}
	if root.Right != nil {
		res += pathSum(root.Right, sum)
	}
	return res
}


func dfs2(node *TreeNode, sum int) int {
	var res int
	if node.Val == sum {
		res ++
	}
	if node.Left != nil {
		res += dfs2(node.Left, sum-node.Val)
	}
	if node.Right != nil {
		res += dfs2(node.Right, sum-node.Val)
	}
	return res
}