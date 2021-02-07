/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
	中序 左 中 右
 */

package review


func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0{
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
