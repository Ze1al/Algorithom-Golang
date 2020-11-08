package TreeNode

import (
	"container/list"
	"fmt"
)

// 后序遍历 recursion
func (tree *Tree) postorderTraversalRecursion(node *TreeNode) {
	if node == nil {
		return
	} else {
		tree.postorderTraversal(node.Left)
		tree.postorderTraversal(node.Right)
		fmt.Println(node)
	}
}

// 后序遍历 左右中
func (tree *Tree) postorderTraversal (node *TreeNode) {
	stack := list.New()
	res := make([]int, 0)
	curr := node
	for stack.Len()>0 || curr != nil {
		for curr != nil {
			res = append(res, curr.Data)
			stack.PushBack(curr.Left)
			curr = curr.Right
		}
		if stack.Len() != 0 {
			node := stack.Back()
			curr = node.Value.(*TreeNode)
			res = append(res, curr.Data)
		}
	}
}