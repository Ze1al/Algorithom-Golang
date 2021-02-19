/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/19
 */

package tencent

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numMap[v]; ok {
			numMap[v]++
			if numMap[v] > len(nums)/2 {
				return v
			}
		} else {
			numMap[v] = 1
		}
	}
	return 0
}
