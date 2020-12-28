/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/24
 */

package ListNode


func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}