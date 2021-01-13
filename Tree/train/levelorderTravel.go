/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/12
	层序遍历 使用队列

 */

package train


func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for i:=0; len(queue)>0; i++{
		res = append(res, []int{})
		temp := []*TreeNode{}
		for j:=0;j<len(queue);j++{
			cur := queue[j]
			res[i] = append(res[i], cur.Val)
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		queue = temp
	}
	return res
}