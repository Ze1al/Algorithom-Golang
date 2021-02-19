/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
 */

package tencent


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{Val : 0}
	p := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return newHead.Next
}
