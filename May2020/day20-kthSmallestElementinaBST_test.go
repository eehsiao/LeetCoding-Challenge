// Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.
// Note:
// You may assume k is always valid, 1 ≤ k ≤ BST's total elements.
// Example 1:
// Input: root = [3,1,4,null,2], k = 1
//    3
//   / \
//  1   4
//   \
//    2
// Output: 1
// Example 2:
// Input: root = [5,3,6,2,4,null,null,1], k = 3
//        5
//       / \
//      3   6
//     / \
//    2   4
//   /
//  1
// Output: 3
// Follow up:
// What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?
//    Hide Hint #1
// Try to utilize the property of a BST.
//    Hide Hint #2
// Try in-order traversal. (Credits to @chan13)
//    Hide Hint #3
// What if you could modify the BST node's structure?
//    Hide Hint #4
// The optimal runtime complexity is O(height of BST).

package main

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		nums []interface{}
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []interface{}{3, 1, 4, nil, 2},
				k:    1,
			},
			want: 1,
		},
		{
			name: "case 2",
			args: args{
				nums: []interface{}{5, 3, 6, 2, 4, nil, nil, 1},
				k:    3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := insertLevelOrder(tt.args.nums, 0)
			if got := kthSmallest(root, tt.args.k); got != tt.want {
				t.Errorf("%v kthSmallest() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
