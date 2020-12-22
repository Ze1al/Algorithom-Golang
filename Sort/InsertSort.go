/*
*  @author: didichuxing.com
 */

package Sort

// 插入排序，将没有排序好的值插到前面以后排好顺序的值里头
// 插入排序由于是从后往前插入的，所以不会改变相对位置 所以是稳定的

func InsertSort(arr []int) []int {
	nLen := len(arr)
	for i := 1; i < nLen; i++ {
		for j := i; j > 0; j-- {
			// 如果小于旁边的值，则交换相对位置
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}
