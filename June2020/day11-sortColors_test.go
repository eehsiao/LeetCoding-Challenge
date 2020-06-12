// Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
// Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
// Note: You are not suppose to use the library's sort function for this problem.
// Example:
// Input: [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
// Follow up:
// A rather straight forward solution is a two-pass algorithm using counting sort.
// First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
// Could you come up with a one-pass algorithm using only constant space?

package main

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "case 2",
			args: args{
				nums: []int{2, 0, 2, 1, 1},
			},
			want: []int{0, 1, 1, 2, 2},
		},
		{
			name: "case 3",
			args: args{
				nums: []int{0, 0},
			},
			want: []int{0, 0},
		},
		{
			name: "case 4",
			args: args{
				nums: []int{1, 0},
			},
			want: []int{0, 1},
		},
		{
			name: "case 5",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: []int{0, 0, 1},
		},
		{
			name: "case 6",
			args: args{
				nums: []int{2, 1, 2},
			},
			want: []int{1, 2, 2},
		},
		{
			name: "case 7",
			args: args{
				nums: []int{1, 0, 1},
			},
			want: []int{0, 1, 1},
		},
		{
			name: "case 8",
			args: args{
				nums: []int{0, 1, 2},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "case 9",
			args: args{
				nums: []int{0, 1, 1, 1, 1, 1, 2},
			},
			want: []int{0, 1, 1, 1, 1, 1, 2},
		},
		{
			name: "case 10",
			args: args{
				nums: []int{0, 2, 2, 2, 0, 2, 1, 1},
			},
			want: []int{0, 0, 1, 1, 2, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sortColors(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("%v searchInsert() = %v, want %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}
