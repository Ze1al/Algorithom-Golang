/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
 */

package LeetCode


func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrder(root)
	return reverseListList(res)
}

func reverseListList(res [][]int) [][]int {
	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}
	return res
}

