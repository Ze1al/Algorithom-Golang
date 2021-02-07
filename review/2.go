/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/4
 */

package review


func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}