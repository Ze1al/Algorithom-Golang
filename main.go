/*
*  @author: didichuxing.com
 */


package main

import (
	"math/rand"
	"time"
	"git.xiaojukeji.com/Leetcode/offer"
)

func main() {
	matrix := [][]int{
		{1,4,7,11,15},
		{2,5,8,12,19},
		{3,6,9,16,22},
		{10,13,14,17,24},
		{18,21,23,26,30},
	}

	target := 5
	offer.FindNumberIn2DArray(matrix, target)
}

func QSortRecursion(arr []int){
	if len(arr) <2{
		return
	}

	head, tail := 0, len(arr) -1
	randomNum := getRandNum(len(arr)-1)
	arr[0], arr[randomNum] = arr[randomNum], arr[0]
	poivt := arr[0]

	for i := 1; i<=tail; {
		if arr[i] <= poivt {
			arr[i], arr[head] = arr[head], arr[i]
			head ++
			i ++
		} else {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail --
		}
	}
	QSortRecursion(arr[:head])
	QSortRecursion(arr[head+1:])
}

func getRandNum(num int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}