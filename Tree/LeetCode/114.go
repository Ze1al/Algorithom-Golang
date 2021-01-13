/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	二叉树展开为链表
 */

package LeetCode




func flatten(root *TreeNode)  {
	list := preorderTravel(root)
	for i := 1; i<len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}


func preorderTravel(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	if root == nil {
		return res
	}
	res = append(res, root)
	res = append(res, preorderTravel(root.Left)...)
	res = append(res, preorderTravel(root.Right)...)
	return res
}