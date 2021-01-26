/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/26
 */

package Top

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	intervals := [][]int {
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	want := [][]int {
		{1,6},
		{8,10},
		{15,18},
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"test", args{intervals: intervals}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
