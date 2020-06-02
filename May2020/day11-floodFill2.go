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

func fillTo2(image [][]int, y, x, srcColor, newColor int) {
	w, d := len(image[0]), len(image)

	image[y][x] = newColor

	if x+1 < w && image[y][x+1] == srcColor {
		fillTo2(image, y, x+1, srcColor, newColor)
	}
	if y+1 < d && image[y+1][x] == srcColor {
		fillTo2(image, y+1, x, srcColor, newColor)
	}
	if x > 0 && image[y][x-1] == srcColor {
		fillTo2(image, y, x-1, srcColor, newColor)
	}
	if y > 0 && image[y-1][x] == srcColor {
		fillTo2(image, y-1, x, srcColor, newColor)
	}

	return
}

func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	srcColor := image[sr][sc]
	if srcColor != newColor {
		fillTo2(image, sr, sc, srcColor, newColor)
	}

	return image
}
