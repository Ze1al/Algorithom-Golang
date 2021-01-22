/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/15
 */

package tencent

func kthSmallest(root *TreeNode, k int) int {
	res :=  inorderTravel(root)
	return res[len(res)-k+1]
}


// 中序遍历 左 中 右
func inorderTravel(root *TreeNode) []int{
	res := []int{}
	if root == nil {
		return res
	}
	cur := root
	stack := []*TreeNode{}
	for cur !=nil || len(stack) > 0{
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
