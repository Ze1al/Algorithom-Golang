/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/9
	合并K个链表
 */

package tencent


func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}



