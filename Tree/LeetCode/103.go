/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/6
 */

package LeetCode

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ans = append(ans, []int{})
		temp := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			cur := queue[j]
			ans[i] = append(ans[i], cur.Val)
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		queue = temp
	}
	for i, v := range ans{
		if i % 2 == 1 {
			ans[i] = reverseList(v)
		}
	}
	return ans
}

func reverseList(temp []int) []int {
	if len(temp) == 0 {
		return temp
	}

	for i, j := 0, len(temp)-1; i < j; {
		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}
	return temp
}
