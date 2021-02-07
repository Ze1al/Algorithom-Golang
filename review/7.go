/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
 */

package review

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		temp := []*TreeNode{}
		resTemp := []int{}
		for i := 0; i<len(queue); i++ {
			cur := queue[i]
			resTemp = append(resTemp, cur.Val)
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		queue = temp
		res = append(res, resTemp)
	}
	return res
}