/*
*  @author: didichuxing.com
 */

package src

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	totalLen := m + n
	if totalLen % 2 == 1 {
		k := (totalLen+1)/2
		return float64(getKthelement(nums1, nums2, k))
	} else {
		return float64(getKthelement(nums1, nums2,(totalLen)/2)) + float64(getKthelement(nums1,nums2,(totalLen)/2+1)) / 2
	}
}


func getKthelement(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
 		if index1 == len(nums1) {
 			return nums2[index2 + k -1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k -1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k/2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}