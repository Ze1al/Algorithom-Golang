/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/4
	N叉树的前序遍历
	中 左 右
 */

package LeetCode


type Node struct {
	Val int
	Children []*Node
}


func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	ans = append(ans, root.Val)
	for _, node := range root.Children {
		ans = append(ans, preorder(node)...)
	}
	return ans
}