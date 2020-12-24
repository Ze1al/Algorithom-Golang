/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2020/12/23
 */

package offer

import "testing"

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test", args{"We are happy."},"We%20are%20happy."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}