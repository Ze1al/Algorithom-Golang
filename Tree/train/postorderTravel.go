/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/12
	后序遍历 leetcode: 145
	左 右 中 -- > 中 右 左
 */

package train

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur.Left)
			cur = cur.Right
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return reverseList(res)
}


func reverseList(res []int) []int {
	for i, j := 0, len(res)-1; i<j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
