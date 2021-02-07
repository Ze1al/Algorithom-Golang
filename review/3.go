/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
	环路检测链表
 */

package review


func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}