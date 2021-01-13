/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	二叉树锯齿形遍历
*/

package train

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		res = append(res, []int{})
		temp := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
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
	return zigzagList(res)
}

func zigzagList(res [][]int) [][]int {
	for i := 0; i < len(res); i++ {
		if i %2 == 1 {
			res[i] = reverseList(res[i])
		}
	}
	return res
}




















