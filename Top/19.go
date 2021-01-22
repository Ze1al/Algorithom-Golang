/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/22
	删除倒数第几个链表
 */

package Top

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: 0, Next: head}

	fast := head
	slow := newHead

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return newHead.Next
}











