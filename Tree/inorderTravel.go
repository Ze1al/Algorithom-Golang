/*
*  @author: didichuxing.com
	中序遍历
*/

package Tree

import (
	"container/list"
	"fmt"
)

func (tree *Tree) inorderTravel(node *TreeNode) []int {
	if node == nil {
		return []int{}
	} else {
		cur := node
		stack := []*TreeNode{}
		res := []int{}
		for cur != nil || len(stack) != 0 {
			if cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			} else {
				cur := stack[len(stack)-1]
				res = append(res, cur.Val)
				cur = cur.Right
			}
		}
		return res
	}
}

func (tree *Tree) _inorderTravel(node *TreeNode) {
	stack := list.New()
	res := make([]int, 0)
	curr := node
	for curr != nil || stack.Len() > 0 {
		for curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		}
		if stack.Len() != 0 {
			node := stack.Back()
			curr = node.Value.(*TreeNode)
			res = append(res, curr.Val)
			curr = curr.Right
			stack.Remove(node)
		}

	}
}

// recursion版本
func (tree *Tree) inorderTravelRecursion(node *TreeNode) {
	if node == nil {
		return
	} else {
		tree.inorderTravelRecursion(node.Left)
		fmt.Sprintln(node.Val)
		tree.inorderTravelRecursion(node.Right)
	}
}
