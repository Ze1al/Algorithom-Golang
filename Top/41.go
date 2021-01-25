/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/25
	数组中没有出现的最小正整数
	1. 长度为N的数组
	2. 如果数组中的值都大于N，则一个也不会被打上标记，则返回N+1
	3. 如果数组中的值在1~N中，则将对应的值打上标记

	如果都打上标记了，则return N+1
	如果都没被打上标记

*/

package Top

import "fmt"

func firstMissingPositive(nums []int) int {
	N := len(nums)

	for i := 0; i < N; i++ {
		if nums[i] <= 0 {
			nums[i] = N + 1
		}
	}

	for i := 0; i < N; i++ {
		index := abs(nums[i])
		if index <= N {
			fmt.Println(index-1)
			nums[index-1] = -abs(nums[index-1])
		}
	}

	for i := 0; i < N; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return N + 1
}
