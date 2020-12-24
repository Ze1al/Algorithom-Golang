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
		{-1, 3},

	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test", args{matrix, -1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("FindNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
