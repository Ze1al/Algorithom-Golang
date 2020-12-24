/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/23
	从尾到头打印链表
 */

package offer

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	cur := head
	res := []int{}
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	for i, j := 0, len(res)-1;i<j;{
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

