/*
*  @author: didichuxing.com
 */

package ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow := head
	fast := head

	for fast != nil && k > 0 {
		fast = fast.Next
		k -= 1
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
