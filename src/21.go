/*
*  @author: didichuxing.com
 */
package src

// 合并两个链表

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	res := newListNode

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val{
			newListNode.Next = l1
			l1 = l1.Next
		} else {
			newListNode.Next = l2
			l2 = l2.Next
		}
		newListNode = newListNode.Next
	}
	if l1 != nil {
		newListNode.Next = l1
	}
	if l2 != nil{
		newListNode.Next = l2
	}
	return res.Next
}
