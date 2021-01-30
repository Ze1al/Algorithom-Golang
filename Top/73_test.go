/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/1/30
 */

package Top

import "testing"

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test", args{[][]int{{1,1,1}, {1,0,1}, {1,1,1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}