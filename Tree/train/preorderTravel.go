/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/12
	前序遍历 leetcode := 144
	中 左 右
 */

package train


func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0{
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur.Right)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return res
}
