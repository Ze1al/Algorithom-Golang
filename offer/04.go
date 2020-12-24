/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/15
	二维数组中寻找该整数
*/

package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix[0])
	m := len(matrix)
	for i := 0; i < m; i++ {
		if len(matrix[i]) == 1 && matrix[i][0] == target {
			return true
		} else if matrix[i][0] <= target && matrix[i][n-1] >= target {
			res := binarySearch(matrix[i], target, 0, n-1)
			if res {
				return true
			} else {
				continue
			}
		}

	}
	return false
}

func binarySearch(nums []int, target int, pre, last int) bool {
	if pre > last {
		return false
	}
	mid := (pre + last) / 2
	if nums[mid] == target {
		return true
	} else if nums[mid] < target {
		return binarySearch(nums, target, mid+1, last)
	} else {
		return binarySearch(nums, target, pre, mid-1)
	}
}
