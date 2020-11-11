/*
*  @author: didichuxing.com
 */

package src

import "testing"

func Test__longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test", args{"ddcdcd"}, "dcdcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("_longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
