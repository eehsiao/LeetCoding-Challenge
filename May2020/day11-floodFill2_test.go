// An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).
// Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.
// To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.
// At the end, return the modified image.
// Example 1:
// Input:
// image = [[1,1,1],[1,1,0],[1,0,1]]
// sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation:
// From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
// by a path of the same color as the starting pixel are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected
// to the starting pixel.
// Note:
// The length of image and image[0] will be in the range [1, 50].
// The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
// The value of each color in image[i][j] and newColor will be an integer in [0, 65535].
//    Hide Hint #1
// Write a recursive function that paints the pixel if it's the correct color, then recurses on neighboring pixels.

package main

import (
	"reflect"
	"testing"
)

func Test_floodFill2(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case 1",
			args: args{
				image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			name: "case 2",
			args: args{
				image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
				sr:       2,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 2, 1}},
		},
		{
			name: "case 3",
			args: args{
				image:    [][]int{{0, 0, 0}, {0, 1, 0}},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{{0, 0, 0}, {0, 2, 0}},
		},
		{
			name: "case 4",
			args: args{
				image:    [][]int{{0, 0, 0}, {0, 1, 1}},
				sr:       1,
				sc:       1,
				newColor: 1,
			},
			want: [][]int{{0, 0, 0}, {0, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill2(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
