/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/25
 */

package Top

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	res := [][]int{
		{1,2,3},
	{1,3,2},
	{2,1,3},
	{2,3,1},
	{3,1,2},
	{3,2,1},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test", args{nums:[]int{1,2,3}},res },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
