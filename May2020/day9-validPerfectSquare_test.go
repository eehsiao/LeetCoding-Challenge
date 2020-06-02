// Given a positive integer num, write a function which returns True if num is a perfect square else False.
// Note: Do not use any built-in library function such as sqrt.
// Example 1:
// Input: 16
// Output: true
// Example 2:
// Input: 14
// Output: false

package main

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				num: 16,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				num: 12,
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				num: 0,
			},
			want: false,
		},
		{
			name: "case 9",
			args: args{
				num: 9,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("%v isPerfectSquare() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
