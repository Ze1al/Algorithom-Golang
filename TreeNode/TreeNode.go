package TreeNode


type TreeNode struct {
	Data int
	Left *TreeNode
	Right *TreeNode
}

type Tree struct{
	root *TreeNode
}

func (tree *Tree) Add(data int) {
	var queue []*TreeNode
	newNode := &TreeNode{Data:data}
	if tree.root == nil {
		tree.root = newNode
	} else {
		queue = append(queue, tree.root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[0+1:]...)

			// insert to right tree
			if data > cur.Data {
				if cur.Right == nil {
					cur.Right = newNode
				} else {
					queue = append(queue, cur.Right)
				}
			} else {
				// insert to left tree
				if cur.Left == nil {
					cur.Left = newNode
				} else {
					queue = append(queue, cur.Left)
				}
			}
		}
	}
}

