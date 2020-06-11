// Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You may assume no duplicates in the array.
// Example 1:
// Input: [1,3,5,6], 5
// Output: 2
// Example 2:
// Input: [1,3,5,6], 2
// Output: 1
// Example 3:
// Input: [1,3,5,6], 7
// Output: 4
// Example 4:
// Input: [1,3,5,6], 0
// Output: 0

package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			name: "case 4",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			want: 0,
		},
		{
			name: "case 5",
			args: args{
				nums:   []int{1, 3},
				target: 2,
			},
			want: 1,
		},
		{
			name: "case 6",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "case 7",
			args: args{
				nums:   []int{1, 2, 4, 6, 7},
				target: 3,
			},
			want: 2,
		},
		{
			name: "case 8",
			args: args{
				nums:   []int{3, 5, 7, 9, 10},
				target: 8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("%v searchInsert() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
