/*
*  @author: didichuxing.com
 */

package src

type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNodeHead := &ListNode{0, nil}
	cur := newListNodeHead

	c := 0  // 进位

	for l1 != nil || l2 != nil{
		var x int
		var y int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + c
		c = sum / 10
		value := sum % 10

		cur.Next = &ListNode{value, nil}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if c != 0 {
		cur.Next = &ListNode{Val: c}
	}
	return newListNodeHead.Next
}




