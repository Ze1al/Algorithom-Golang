/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/7
	两数相加
 */

package tencent

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{Val:0}
	p:= newHead
	c := 0
	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + c
		temp := sum % 10
		p.Next = &ListNode{Val: temp}
		p = p.Next
		c = sum / 10
	}
	if c != 0 {
		p.Next = &ListNode{Val: c}
	}
	return newHead.Next
}