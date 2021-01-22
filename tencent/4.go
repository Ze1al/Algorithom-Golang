/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/14
	寻找两个正序数组的中位数
	使用O(log(m+n))
	二分法
 */

package tencent

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total % 2 == 1 {
		midIndex := total/2
		return float64(getKthElement(nums1, nums2, midIndex + 1))
	} else {
		midIndex1, midIndex2 := total/2 - 1, total/2
		return (float64(getKthElement(nums1, nums2, midIndex1 + 1)) + float64(getKthElement(nums1, nums2, midIndex2 + 1)))/2
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	if nums1 == nil {
		return nums2[index2+k+1]
	}
	if nums2 == nil {
		return nums2[index1+k+1]
	}


	if k == 1{
		return min(nums1[k], nums2[k])
	}
	half := k/2
	newIndex1 := min(index1 + half, len(nums1)) - 1
	newIndex2 := min(index2 + half, len(nums2)) - 1
	pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
	if pivot1 < pivot2 {
		k -= (newIndex1 - index1 + 1)
		index1 = newIndex1 + 1
	} else {
		k -= (newIndex2 - index2 + 1)
		index2 = newIndex2 + 1
	}
	return 0
}


func min(x, y int) int {
	if x<y{
		return x
	} else {
		return y
	}
}