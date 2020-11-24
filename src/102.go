/*
*  @author: didichuxing.com
 */

package src


func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}


func bfs(root *TreeNode) []int{
	queue := []*TreeNode{root}
	ret := []int{}
	for len(queue) != 0 {
		node := queue[len(queue)-1]
		ret = append(ret, node.Val)
		if node.Left != nil{
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return ret
}




