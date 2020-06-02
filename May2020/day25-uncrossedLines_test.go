// We write the integers of A and B (in the order they are given) on two separate horizontal lines.
// Now, we may draw connecting lines: a straight line connecting two numbers A[i] and B[j] such that:
// A[i] == B[j];
// The line we draw does not intersect any other connecting (non-horizontal) line.
// Note that a connecting lines cannot intersect even at the endpoints: each number can only belong to one connecting line.
// Return the maximum number of connecting lines we can draw in this way.
// Example 1:
// 1 4 2
// |  \
// 1 2 4
// Input: A = [1,4,2], B = [1,2,4]
// Output: 2
// Explanation: We can draw 2 uncrossed lines as in the diagram.
// We cannot draw 3 uncrossed lines, because the line from A[1]=4 to B[2]=4 will intersect the line from A[2]=2 to B[1]=2.
// Example 2:
// Input: A = [2,5,1,2,5], B = [10,5,2,1,5,2]
// Output: 3
// Example 3:
// Input: A = [1,3,7,1,7,5], B = [1,9,2,5,1]
// Output: 2
// Note:
// 1 <= A.length <= 500
// 1 <= B.length <= 500
// 1 <= A[i], B[i] <= 2000
//    Hide Hint #1
// Think dynamic programming. Given an oracle dp(i,j) that tells us how many lines A[i:], B[j:] [the sequence A[i], A[i+1], ... and B[j], B[j+1], ...] are uncrossed, can we write this as a recursion?

package main

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				A: []int{1, 4, 2},
				B: []int{1, 2, 4},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				A: []int{2, 5, 1, 2, 5},
				B: []int{10, 5, 2, 1, 5, 2},
			},
			want: 3,
		},
		{
			name: "case 3",
			args: args{
				A: []int{1, 3, 7, 1, 7, 5},
				B: []int{1, 9, 2, 5, 1},
			},
			want: 2,
		},
		{
			name: "case 4",
			args: args{
				A: []int{2, 1, 2, 3, 1},
				B: []int{1, 3, 3, 3, 1, 1},
			},
			want: 3,
		},
		{
			name: "case 5",
			args: args{
				A: []int{3, 1, 4, 1, 1, 3, 5, 1, 2, 2},
				B: []int{4, 1, 5, 2, 1, 1, 1, 5, 3, 1, 1, 1, 2, 3, 1, 4, 3, 5, 5, 3, 1, 2, 3, 2, 4, 1, 1, 1, 5, 3},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("%v maxUncrossedLines() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
