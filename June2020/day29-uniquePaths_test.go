// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?
// Above is a 7 x 3 grid. How many possible unique paths are there?
// Example 1:
// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Right -> Down
// 2. Right -> Down -> Right
// 3. Down -> Right -> Right
// Example 2:
// Input: m = 7, n = 3
// Output: 28
// Constraints:
// 1 <= m, n <= 100
// It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.

package main

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
