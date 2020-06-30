// Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.
// Example 1:
// Input: [1,3,4,2,2]
// Output: 2
// Example 2:
// Input: [3,1,3,4,2]
// Output: 3
// Note:
// 1.You must not modify the array (assume the array is read only).
// 2.You must use only constant, O(1) extra space.
// 3.Your runtime complexity should be less than O(n2).
// 4.There is only one duplicate number in the array, but it could be repeated more than once.

package main

import "testing"

func Test_findDuplicate(t *testing.T) {
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
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
