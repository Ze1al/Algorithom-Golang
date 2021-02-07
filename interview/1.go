/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/2
	单链表快排 时间复杂度为Onlogn
 */

package interview


// 递归
func sortList(head *ListNode) *ListNode {
	// 递归出口
	if  head == nil || head.Next == nil {
		return head
	}

	// 拆分
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil

	// 归并
	left, right := sortList(head), sortList(mid)
	res := &ListNode{Val:0}
	p := res
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		 p = p.Next
	}

	if left != nil {
		p.Next = left
	} else {
		p.Next = right
	}
	return res.Next
}