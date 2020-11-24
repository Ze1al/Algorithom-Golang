/*
*  @author: didichuxing.com
 */

package ListNode

func kthToLast(head *ListNode, k int) int {
	fast := head
	slow := head
	for fast != nil && k > 0 {
		fast = fast.Next
		k -= 1
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val

}
