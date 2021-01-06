/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/5
	二叉搜索树的第k大节点
	中序遍历
 */

package LeetCode


func kthLargest(root *TreeNode, k int) int {
	res := []int{}
	cur := root
	stack := []*TreeNode{}
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
	return res[len(res)-k]
}