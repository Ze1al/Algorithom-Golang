/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/22
 */

package Top


func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{
		return nil
	}
	res := lists[0]
	for i:=1; i<len(lists);i++ {
		res = mergeTwoList(res, lists[i])
	}
	return res
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	newHead := &ListNode{Val:0}
	p := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	}
	if l2 == nil {
		p.Next = l1
	}
	return newHead.Next
}