/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/26
	合并区间
*/

package Top

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool{
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[len(prev)-1] < intervals[i][0] {
			res = append(res, prev)
			prev = intervals[i]
		} else {
			prev[len(prev)-1] = max(intervals[i][len(intervals[i])-1], prev[len(prev)-1])
		}
	}
	if len(prev) != 0 {
		res = append(res, prev)
	}
	return res
}
