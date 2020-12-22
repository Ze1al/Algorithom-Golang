/*
*  @author: didichuxing.com
 */

package Sort

// 冒泡排序，冒泡排序不断的对比两个相邻直接的数值大小，稳定的，时间复杂度为O(n2)

func BubbleSort(arr []int) []int {
	nLen := len(arr)
	for i := 0; i < nLen; i++ {
		for j := i + 1; j < nLen; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}
