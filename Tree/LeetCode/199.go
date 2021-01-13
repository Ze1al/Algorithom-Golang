/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	二叉树的右视图
	直接层序遍历 然后获取最后一个位置
 */

package LeetCode


func rightSideView(root *TreeNode) []int {
	var res []int
	list := level(root)
	for i := 0; i<len(list); i++ {
		res = append(res, list[i][len(list[i])-1])
	}
	return res
}


func level(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		ret = append(ret, []int{})
		temp := []*TreeNode{}
		for j:=0; j<len(queue); j++ {
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