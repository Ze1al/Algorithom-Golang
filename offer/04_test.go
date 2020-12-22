/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/17
 */

package offer

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	matrix := [][]int{
		{1,4,7,11,15},
		{2,5,8,12,19},
		{3,6,9,16,22},
		{10,13,14,17,24},
		{18,21,23,26,30},
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test", args{matrix, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("FindNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
