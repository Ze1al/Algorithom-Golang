/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/15
	二维数组中寻找该整数
*/

package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix[0])
	m := len(matrix)
	for i := 0; i<m;i++{
		if matrix[i][0] < target && matrix[i][n] > target {
			res := binarySearch(matrix[i], target, 0, n)
			if res != -1 {
				return true
			} else {
				continue
			}
		}
	}
	return false
}


func binarySearch(nums []int, target int, pre, last int) int {
	if pre < last {
		return -1
	}
	mid := (pre+last)/2
	if nums[mid] == target {
		return nums[mid]
	} else if nums[mid] < target {
		return binarySearch(nums, target, mid+1, last)
	} else {
		return binarySearch(nums, target, pre, mid-1)
	}
}



