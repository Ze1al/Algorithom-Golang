/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/13
	两数相加
 */

package tencent


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	pHead := &ListNode{Val:0}
	p3 := pHead
	var c int
	for p1 != nil && p2 != nil {
		res := p1.Val + p2.Val
		cur := (res+c)%10
		tempNode := &ListNode{Val:cur}
		p3.Next = tempNode
		p3 = p3.Next
		c = (res+c)/10
		p1 = p1.Next
		p2 = p2.Next
	}
	for p1 != nil {
		res := p1.Val + c
		temp := &ListNode{Val:res%10}
		p3.Next = temp
		p3 = p3.Next
		c = res/10
		p1 = p1.Next
	}
	for p2 != nil {
		res := p2.Val + c
		temp := &ListNode{Val:res%10}
		p3.Next = temp
		p3 = p3.Next
		c = res/10
		p2 = p2.Next
	}
	if c != 0 {
		temp := &ListNode{Val:c}
		p3.Next = temp
		p3 = p3.Next
	}
	return pHead.Next
}

