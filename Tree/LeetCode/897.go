/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/5
	递增的顺序查找树
	中序遍历 左中右
 */

package LeetCode

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := inorderTraversal(root)

	temp := &TreeNode{Val:res[0]}
	p := temp
	for i := 1; i<len(res);i++{
		p.Right = &TreeNode{Val:res[i]}
		p = p.Right
	}
	return temp
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	ans = append(ans, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}