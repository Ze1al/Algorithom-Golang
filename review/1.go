/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
	链表中第几个元素
*/

package review

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	for k > 0 && fast != nil{

		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
