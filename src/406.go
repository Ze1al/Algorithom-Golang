/*
*  @author: didichuxing.com
	排队，使每个整数满足（h,k）的规则，h代表身高，k代表前面有几个成员
	解题思路：高个子先排好队，然后让矮个按照K往里面插入
*/

package src

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return (people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i+1])
		people[p[1]] = p
	}
	return people
}
