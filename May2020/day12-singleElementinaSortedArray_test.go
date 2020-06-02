// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.
// Example 1:
// Input: [1,1,2,3,3,4,4,8,8]
// Output: 2
// Example 2:
// Input: [3,3,7,7,10,11,11]
// Output: 10
// Note: Your solution should run in O(log n) time and O(1) space.

package main

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 3, 7, 7, 10, 11, 11},
			},
			want: 10,
		},
		{
			name: "case 3",
			args: args{
				nums: []int{3, 3, 7, 7, 10, 11, 11, 11},
			},
			want: 10,
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
