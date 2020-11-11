/*
*  @author: didichuxing.com
 */

package src

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Test", args{a},6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
