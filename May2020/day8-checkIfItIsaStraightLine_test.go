// You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.
// Example 1:
// https://assets.leetcode.com/uploads/2019/10/15/untitled-diagram-2.jpg
// Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
// Output: true
// Example 2:
// https://assets.leetcode.com/uploads/2019/10/09/untitled-diagram-1.jpg
// Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
// Output: false
// Constraints:
// 2 <= coordinates.length <= 1000
// coordinates[i].length == 2
// -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
// coordinates contains no duplicate point.
//    Hide Hint #1
// If there're only 2 points, return true.
//    Hide Hint #2
// Check if all other points lie on the line defined by the first 2 points.
//    Hide Hint #3
// Use cross product to check collinearity.

package main

import "testing"

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				coordinates: [][]int{{1, 2}, {2, 3}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("%v checkStraightLine() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
