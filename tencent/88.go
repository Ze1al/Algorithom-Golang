/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/18
	合并两个有序数组
 */

package tencent

func merge(nums1 []int, m int, nums2 []int, n int)  {
	res := make([]int, m + n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i != m {
		res = append(res, nums1[i:]...)
	}
	if j != n {
		res = append(res, nums2[j:]...)
	}
	nums1 = res
}
