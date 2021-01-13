/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/12
 */

package train

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	res := inorder(root)
	resRoot := &TreeNode{Val:res[0]}
	p := resRoot
	for i := 1; i<len(res); i++ {
		temp := &TreeNode{Val:res[i]}
		p.Right = temp
		p = temp
	}
	return resRoot
}


func inorder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0|| cur != nil {
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

