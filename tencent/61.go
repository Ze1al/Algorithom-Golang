/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	旋转链表
*/

package tencent

func rotateRight(head *ListNode, k int) *ListNode {
	// close link to a ring
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	p := head
	n := 1
	for p.Next != nil {
		p = p.Next
		n += 1
	}
	p.Next = head

	// 断裂点
	// tail: (n-k % n-1)
	// head: (n-k % n)

	cur := head
	for i := 0; i < n-k % n-1; i++ {
		cur = cur.Next
	}
	newHead := cur.Next
	cur.Next = nil
	return newHead

}
