/*
*  @author: didichuxing.com
 */

package Tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


type Tree struct {
	root *TreeNode
}

func (tree *Tree) insert (data int) {
	var queue []*TreeNode
	newTreeNode := &TreeNode{Val:data}
	if tree.root == nil {
		tree.root = newTreeNode
	} else {
		queue = append(queue, tree.root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[1:]...)
			// insert right tree
			if data > cur.Val {
				if cur.Right == nil {
					cur.Right = newTreeNode
				} else {
					queue = append(queue, cur.Right)
				}
			} else {
				if cur.Left == nil {
					cur.Left = newTreeNode
				} else {
					queue = append(queue, cur.Left)
				}
			}
		}

	}
}



