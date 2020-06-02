// In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
// Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
// We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
// Return true if and only if the nodes corresponding to the values x and y are cousins.
// Example 1:
//     1
//    / \
//   2   3
//  /
// 4
// Input: root = [1,2,3,4], x = 4, y = 3
// Output: false
// Example 2:
//     1
//    / \
//   2   3
//    \   \
//     4   5
// Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
// Output: true
// Example 3:
//     1
//    / \
//   2   3
//    \
//     4
// Input: root = [1,2,3,null,4], x = 2, y = 3
// Output: false
// Note:
// The number of nodes in the tree will be between 2 and 100.
// Each node has a unique integer value from 1 to 100.

package main

import (
	"fmt"
	"testing"
)

func Test_isCousins(t *testing.T) {
	type args struct {
		nums []interface{}
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				nums: []interface{}{1, 2, 3, 4},
				x:    4,
				y:    3,
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				nums: []interface{}{1, 2, 3, nil, 4, nil, 5},
				x:    5,
				y:    4,
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				nums: []interface{}{1, 2, 3, nil, 4},
				x:    2,
				y:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := insertLevelOrder(tt.args.nums, 0)
			postOrder(root)
			fmt.Println()
			if got := isCousins(root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("%v isCousins() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
