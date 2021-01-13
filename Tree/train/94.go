/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/12
	中序 左 中 右
 */

package train

//func inorderTraversal(root *TreeNode) []int {
//	res := []int{}
//	if root == nil {
//		return res
//	}
//	stack := []*TreeNode{}
//	cur := root
//	for cur != nil || len(stack) > 0{
//		if cur != nil {
//			stack = append(stack, cur)
//			cur = cur.Left
//		} else {
//			cur = stack[len(stack)-1]
//			stack = stack[:len(stack)-1]
//			res = append(res, cur.Val)
//			cur = cur.Right
//		}
//	}
//	return res
//}
