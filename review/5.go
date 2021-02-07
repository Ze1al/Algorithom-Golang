/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
 */

package review

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}


