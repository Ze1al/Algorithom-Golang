/*
*  @author: didichuxing.com
 */

package Sort

// 归并排序，归并排序的时间复杂度O(nlogn) 空间复杂度是O(n)
// 稳定

func Merge(arr []int) []int {
	nLen := len(arr)
	if nLen < 2 {
		return arr
	}
	mid := nLen / 2
	leftArr := Merge(arr[:mid])
	rightArr := Merge(arr[mid:])
	res := MergeSort(leftArr, rightArr)
	return res
}

func MergeSort(leftArr []int, rightArr []int) []int {
	arrNew := []int{}
	var i, j int
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			arrNew = append(arrNew, leftArr[i])
			i += 1
		} else {
			arrNew = append(arrNew, rightArr[j])
			j += 1
		}
	}
	arrNew = append(arrNew, leftArr[i:]...)
	arrNew = append(arrNew, rightArr[j:]...)
	return arrNew
}
