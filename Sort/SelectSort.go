/*
*  @author: didichuxing.com
 */

package Sort

// 选择排序，每次选择出最小的数值，然后放到第一位

func SelectSort(arr []int) []int {
	nLen := len(arr)
	for i := 0; i < nLen; i++ {
		min := i // 记录每次的最小值
		for j := i + 1; j < nLen; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
