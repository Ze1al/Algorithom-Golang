/*
*  @author: didichuxing.com
倒数第N个节点
*/

package src

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head != nil {
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
