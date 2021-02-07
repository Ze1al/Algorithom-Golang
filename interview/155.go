/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/2
	最小栈
 */

package interview

import "math"

type MinStack struct {
	stack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MinInt64},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}


