/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/5
	从上到下打印二叉树
 */

package LeetCode


func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil{
		return ret
	}
	queue := []*TreeNode{root}

	for i:=0; len(queue)>0; i++{
		ret = append(ret, []int{})
		temp := []*TreeNode{}

		for j := 0; j<len(queue); j++{
			cur := queue[j]
			ret[i] = append(ret[i], cur.Val)
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		queue = temp
	}
	return ret
}