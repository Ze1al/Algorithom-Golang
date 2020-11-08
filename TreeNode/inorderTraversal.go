package TreeNode

import (
	"container/list"
	"fmt"
)

// recursion
func (tree *Tree) inorderTraversalRecursion (node *TreeNode){
	if node == nil {
		return
	} else {
		tree.inorderTraversalRecursion(node.Left)
		fmt.Println(node)
		tree.inorderTraversalRecursion(node.Right)
	}
}

func (tree *Tree) inorderTraversal (node *TreeNode) {
	stack := list.New()
	res := make([]int, 0)
	curr := node
	for stack.Len()>0 || curr != nil {
		for curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		}
		if stack.Len() != 0 {
			node := stack.Back()
			curr = node.Value.(*TreeNode)
			res = append(res, curr.Data)
			curr = curr.Right
			stack.Remove(node)
		}
	}
}















