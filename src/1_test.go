/*
*  @author: didichuxing.com
 */

package src

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	nums := []int{2, 7, 11, 15}
	wants := []int{0,1}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"TestTwoSum", args{nums,9},wants},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
