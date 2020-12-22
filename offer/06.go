/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/17
	从尾部开始打印链表，从尾部开始
 */

package offer

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := []int{}
	p := head
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	for i, j := 0, len(res); i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}




