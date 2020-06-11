// Given an integer, write a function to determine if it is a power of two.
// Example 1:
// Input: 1
// Output: true
// Explanation: 20 = 1
// Example 2:
// Input: 16
// Output: true
// Explanation: 24 = 16
// Example 3:
// Input: 218
// Output: false

package main

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{n: 1},
			want: true,
		},
		{
			name: "case 2",
			args: args{n: 16},
			want: true,
		},
		{
			name: "case 3",
			args: args{n: 218},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
