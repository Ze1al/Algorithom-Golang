/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/19
 */

package tencent

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for  fast != nil && fast.Next != nil  {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
